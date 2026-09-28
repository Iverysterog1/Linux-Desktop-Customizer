package main

import (
	"fmt"
	"html/template"
	"net"
	"net/http"

	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/ui"
)

const graphicalPage = `<!doctype html>
<html lang="{{.Model.Locale}}">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width,initial-scale=1">
<title>{{.Model.Product}}</title>
<style>
:root{font-family:system-ui,sans-serif;color-scheme:light dark}body{margin:0;display:grid;grid-template-columns:minmax(13rem,22rem) 1fr;min-height:100vh}nav{padding:1.5rem;border-inline-end:1px solid ButtonBorder}main{padding:2rem;max-width:70rem}.brand{font-weight:700;font-size:1.2rem}.version{opacity:.7}.screen{display:block;padding:.8rem;margin:.5rem 0;border-radius:.6rem;border:1px solid ButtonBorder}.screen[aria-disabled="true"]{opacity:.55}.warning{padding:.8rem;border:1px solid ButtonBorder;border-radius:.6rem}ul{padding-inline-start:1.25rem}@media(prefers-reduced-motion:reduce){*{scroll-behavior:auto!important}}@media(max-width:700px){body{grid-template-columns:1fr}nav{border-inline-end:0;border-block-end:1px solid ButtonBorder}}
</style>
</head>
<body>
<nav aria-label="{{.Labels.Status}}"><div class="brand">{{.Model.Product}}</div><div class="version">{{.Model.Version}}</div>{{range .Model.Screens}}<div class="screen" aria-disabled="{{if .Enabled}}false{{else}}true{{end}}"><strong>{{.Title}}</strong><br><small>{{if .Enabled}}{{$.Labels.Available}}{{else}}{{$.Labels.Gated}}{{end}}</small></div>{{end}}</nav>
<main><h1>{{.Labels.Status}}</h1>{{range .Model.Screens}}<section id="{{.ID}}"><h2>{{.Title}}</h2><p>{{.Description}}</p>{{if .Reason}}<p><strong>{{$.Labels.Gated}}:</strong> {{.Reason}}</p>{{end}}</section>{{end}}{{range .Model.Warnings}}<p class="warning"><strong>{{$.Labels.Warning}}:</strong> {{.}}</p>{{end}}</main>
</body></html>`

type graphicalView struct {
	Model  ui.Model
	Labels ui.PresentationLabels
}

func graphicalHandler(model ui.Model) (http.Handler, error) {
	t, err := template.New("graphical").Parse(graphicalPage)
	if err != nil {
		return nil, err
	}
	view := graphicalView{Model: model, Labels: ui.LabelsForLocale(model.Locale)}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Content-Security-Policy", "default-src 'none'; style-src 'unsafe-inline'; base-uri 'none'; form-action 'none'; frame-ancestors 'none'")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		if err := t.Execute(w, view); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}), nil
}

func validateLoopbackAddress(addr string) error {
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return fmt.Errorf("invalid listen address: %w", err)
	}
	ip := net.ParseIP(host)
	if ip == nil || !ip.IsLoopback() {
		return fmt.Errorf("graphical runtime must bind to an explicit loopback IP")
	}
	return nil
}

func serveGraphical(model ui.Model, addr string) error {
	if err := validateLoopbackAddress(addr); err != nil {
		return err
	}
	handler, err := graphicalHandler(model)
	if err != nil {
		return err
	}
	server := &http.Server{Addr: addr, Handler: handler, ReadHeaderTimeout: 5_000_000_000}
	return server.ListenAndServe()
}
