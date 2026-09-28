package ui

import (
	"strings"
	"testing"

	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/adapter"
)

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
	for _, screen := range m.Screens {
		if screen.ID == "desktop" && screen.Enabled {
			t.Fatal("desktop screen must remain gated until a native adapter is integrated")
		}
	}
	if len(m.Warnings) == 0 {
		t.Fatal("foundation model should explain current adapter limits")
	}
	if !m.Accessibility.ReducedMotionSupported || !m.Accessibility.ReadableTextSupported {
		t.Fatal("foundation model must expose accessibility capabilities")
	}
}

func TestFoundationModelPortuguese(t *testing.T) {
	t.Parallel()
	m := FoundationModelForLocale("test", nil, "pt_PT.UTF-8")
	if m.Locale != LocalePortuguese {
		t.Fatalf("locale = %q", m.Locale)
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
	if m.Screens[0].Title != "Home" {
		t.Fatalf("fallback title = %q", m.Screens[0].Title)
	}
}
