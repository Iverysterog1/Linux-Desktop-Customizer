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
:root{font-family:system-ui,sans-serif;color-scheme:light dark}*{box-sizing:border-box}body{margin:0;display:grid;grid-template-columns:minmax(14rem,22rem) 1fr;min-height:100vh}.skip{position:absolute;inset-inline-start:.75rem;top:.5rem;transform:translateY(-180%);padding:.55rem .75rem;background:Canvas;color:CanvasText;border:2px solid ButtonText;border-radius:.4rem;z-index:10}.skip:focus{transform:none}nav{padding:1.5rem;border-inline-end:1px solid ButtonBorder}.brand{font-weight:750;font-size:1.2rem}.tagline{margin:.35rem 0 0;max-width:24rem;opacity:.82}.version{margin-top:.35rem;opacity:.7}.nav-list{display:grid;gap:.5rem;margin-top:1.25rem}.screen{display:block;padding:.8rem;border-radius:.6rem;border:1px solid ButtonBorder;text-decoration:none;color:inherit}.screen:focus-visible{outline:3px solid Highlight;outline-offset:2px}.screen[aria-disabled="true"]{opacity:.65}.state{display:inline-block;margin-top:.25rem;font-size:.82rem;font-weight:650}.state.gated{border:1px solid ButtonBorder;border-radius:999px;padding:.1rem .45rem}main{padding:2rem;max-width:70rem}.status{margin-bottom:1.5rem;padding:1rem;border:1px solid ButtonBorder;border-radius:.7rem}.warning{padding:.8rem;border:1px solid ButtonBorder;border-radius:.6rem}section{scroll-margin-top:1rem}ul{padding-inline-start:1.25rem}@media(prefers-reduced-motion:reduce){*{scroll-behavior:auto!important;transition:none!important;animation:none!important}}@media(forced-colors:active){.screen,.status,.warning{forced-color-adjust:auto}.screen[aria-disabled="true"]{opacity:1;border-style:dashed}}@media(max-width:700px){body{grid-template-columns:1fr}nav{border-inline-end:0;border-block-end:1px solid ButtonBorder}}
</style>
</head>
<body>
<a class="skip" href="#content">{{.Labels.Status}}</a>
<nav aria-label="{{.Labels.Status}}">
<div class="brand">{{.Model.Product}}</div>
{{if .Model.BrandTagline}}<p class="tagline">{{.Model.BrandTagline}}</p>{{end}}
<div class="version">{{.Model.Version}}</div>
<div class="nav-list">{{range .Model.Screens}}<a class="screen" href="#{{.ID}}" aria-disabled="{{if .Enabled}}false{{else}}true{{end}}"><strong>{{.Title}}</strong><br><span class="state {{if .Enabled}}available{{else}}gated{{end}}">{{if .Enabled}}{{$.Labels.Available}}{{else}}{{$.Labels.Gated}}{{end}}</span></a>{{end}}</div>
</nav>
<main id="content" tabindex="-1">
<div class="status"><h1>{{.Labels.Status}}</h1><p>{{.Model.Product}} {{.Model.Version}}</p></div>
{{range .Model.Screens}}<section id="{{.ID}}" aria-labelledby="{{.ID}}-title"><h2 id="{{.ID}}-title">{{.Title}}</h2><p>{{.Description}}</p>{{if .Reason}}<p><strong>{{$.Labels.Gated}}:</strong> {{.Reason}}</p>{{end}}</section>{{end}}
{{range .Model.Warnings}}<p class="warning"><strong>{{$.Labels.Warning}}:</strong> {{.}}</p>{{end}}
</main>
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
