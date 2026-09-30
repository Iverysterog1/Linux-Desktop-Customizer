package main

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/ui"
)

func TestGraphicalHandlerPortugueseAndEscaping(t *testing.T) {
	model := ui.Model{
		Product:      "Linux Desktop Customizer <unsafe>",
		Version:      "test",
		Locale:       ui.LocalePortuguese,
		BrandTagline: "Criado por um. Melhorado por muitos. Disponível para todos.",
		KDEValidation: ui.ValidationStatus{ID: "kde.real-plasma", Title: "Validação Plasma real", State: ui.ValidationNotRun, Detail: "Ainda não validado visualmente."},
		Screens: []ui.Screen{{ID: "home", Title: "Início", Description: "Seguro <script>alert(1)</script>", Enabled: true}},
		Warnings: []string{"Aviso <b>seguro</b>"},
	}
	h, err := graphicalHandler(model)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/", nil))
	body := rec.Body.String()
	for _, want := range []string{
		"Estado da interface:",
		"disponível",
		"Início",
		"Criado por um. Melhorado por muitos. Disponível para todos.",
		`href="#home"`,
		`href="#content"`,
		"&lt;unsafe&gt;",
		"&lt;script&gt;alert(1)&lt;/script&gt;",
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing %q: %s", want, body)
		}
	}
	if strings.Contains(body, "<script>alert(1)</script>") {
		t.Fatal("unescaped model content reached HTML")
	}
	if got := rec.Header().Get("Content-Security-Policy"); got == "" {
		t.Fatal("missing Content-Security-Policy")
	}
}

func TestGraphicalHandlerEnglishFallbackAndGatedState(t *testing.T) {
	model := ui.Model{
		Product:      "LTC",
		Version:      "test",
		Locale:       "unknown",
		BrandTagline: "Created by one. Improved by many. Available to all.",
		Screens:      []ui.Screen{{ID: "themes", Title: "Themes", Enabled: false, Reason: "gated"}},
	}
	h, err := graphicalHandler(model)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/", nil))
	body := rec.Body.String()
	for _, want := range []string{
		"UI status:",
		"gated",
		"Themes",
		`aria-disabled="true"`,
		"Created by one. Improved by many. Available to all.",
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("response missing %q", want)
		}
	}
}

func TestGraphicalPageIncludesAccessibilityCSS(t *testing.T) {
	for _, want := range []string{
		"prefers-reduced-motion:reduce",
		"forced-colors:active",
		":focus-visible",
	} {
		if !strings.Contains(graphicalPage, want) {
			t.Fatalf("graphicalPage missing %q", want)
		}
	}
}

func TestValidateLoopbackAddress(t *testing.T) {
	for _, addr := range []string{"127.0.0.1:7788", "[::1]:7788"} {
		if err := validateLoopbackAddress(addr); err != nil {
			t.Fatalf("validateLoopbackAddress(%q): %v", addr, err)
		}
	}
	for _, addr := range []string{"0.0.0.0:7788", "[::]:7788", "localhost:7788", ":7788", "bad"} {
		if err := validateLoopbackAddress(addr); err == nil {
			t.Fatalf("validateLoopbackAddress(%q) unexpectedly accepted", addr)
		}
	}
}

func TestGraphicalHandlerNotFound(t *testing.T) {
	h, err := graphicalHandler(ui.Model{Product: "LTC", Locale: ui.LocaleEnglish})
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/other", nil))
	if rec.Code != 404 {
		t.Fatalf("status = %d, want 404", rec.Code)
	}
}
