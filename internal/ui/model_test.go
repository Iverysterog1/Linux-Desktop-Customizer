package ui

import (
	"context"
	"strings"
	"testing"

	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/adapter"
	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/capability"
)

type modelTestAdapter struct{ name string }

func (a modelTestAdapter) Name() string { return a.name }
func (a modelTestAdapter) Capabilities(context.Context) []capability.Capability { return nil }
func (a modelTestAdapter) Read(context.Context, string) (string, bool, error) { return "", false, nil }
func (a modelTestAdapter) Set(context.Context, string, string) error { return nil }
func (a modelTestAdapter) Unset(context.Context, string) error { return nil }

func TestFoundationModelDoesNotOverclaimDesktopSupport(t *testing.T) {
	t.Parallel()
	a, err := adapter.NewFileAdapter(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	r, err := adapter.NewRegistry(a)
	if err != nil {
		t.Fatal(err)
	}
	m := FoundationModel("test", r)
	if m.Mode != "rebuild-foundation" {
		t.Fatalf("mode = %q", m.Mode)
	}
	if m.Locale != LocaleEnglish {
		t.Fatalf("default locale = %q", m.Locale)
	}
	if m.BrandTagline != "Created by one. Improved by many. Available to all." {
		t.Fatalf("brand tagline = %q", m.BrandTagline)
	}
	for _, screen := range m.Screens {
		if screen.ID == "desktop" && screen.Enabled {
			t.Fatal("desktop screen must remain gated when no native desktop adapter is present")
		}
	}
	if len(m.Warnings) == 0 {
		t.Fatal("foundation model should explain current adapter limits")
	}
	if !m.Accessibility.ReducedMotionSupported ||
		!m.Accessibility.ReadableTextSupported ||
		!m.Accessibility.HighContrastSupported ||
		!m.Accessibility.KeyboardNavigationSupported {
		t.Fatal("foundation model must expose accessibility capabilities")
	}
}

func TestFoundationModelEnablesScopedKDEDesktopCapability(t *testing.T) {
	t.Parallel()
	r, err := adapter.NewRegistry(modelTestAdapter{name: "kde-config"})
	if err != nil {
		t.Fatal(err)
	}
	m := FoundationModel("test", r)
	var desktop Screen
	for _, screen := range m.Screens {
		if screen.ID == "desktop" {
			desktop = screen
			break
		}
	}
	if !desktop.Enabled {
		t.Fatal("desktop screen should be enabled when reviewed kde-config adapter is available")
	}
	if !strings.Contains(desktop.Description, "color-scheme") {
		t.Fatalf("desktop description must remain scoped, got %q", desktop.Description)
	}
	if desktop.Reason != "" {
		t.Fatalf("enabled desktop screen should not carry a gated reason: %q", desktop.Reason)
	}
}

func TestFoundationModelPortuguese(t *testing.T) {
	t.Parallel()
	m := FoundationModelForLocale("test", nil, "pt_PT.UTF-8")
	if m.Locale != LocalePortuguese {
		t.Fatalf("locale = %q", m.Locale)
	}
	if m.BrandTagline != "Criado por um. Melhorado por muitos. Disponível para todos." {
		t.Fatalf("Portuguese tagline = %q", m.BrandTagline)
	}
	if m.Screens[0].Title != "Início" {
		t.Fatalf("home title = %q", m.Screens[0].Title)
	}
	if len(m.Warnings) != 1 || !strings.Contains(m.Warnings[0], "Não existem") {
		t.Fatalf("Portuguese warning = %#v", m.Warnings)
	}
}

func TestFoundationModelUnknownLocaleFallsBackToEnglish(t *testing.T) {
	t.Parallel()
	m := FoundationModelForLocale("test", nil, "fr-FR")
	if m.Locale != LocaleEnglish {
		t.Fatalf("locale = %q, want %q", m.Locale, LocaleEnglish)
	}
	if m.BrandTagline != "Created by one. Improved by many. Available to all." {
		t.Fatalf("fallback tagline = %q", m.BrandTagline)
	}
	if m.Screens[0].Title != "Home" {
		t.Fatalf("fallback title = %q", m.Screens[0].Title)
	}
}
