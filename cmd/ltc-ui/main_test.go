package main

import (
	"testing"

	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/ui"
)

func TestLocaleFromEnvironmentPrecedence(t *testing.T) {
	t.Setenv("LANG", "en_US.UTF-8")
	t.Setenv("LC_MESSAGES", "pt_PT.UTF-8")
	t.Setenv("LC_ALL", "pt_PT.UTF-8")
	if got := localeFromEnvironment(); got != "pt_PT.UTF-8" {
		t.Fatalf("localeFromEnvironment() = %q", got)
	}
}

func TestLocaleFromEnvironmentFallsBackToEnglish(t *testing.T) {
	t.Setenv("LANG", "")
	t.Setenv("LC_MESSAGES", "")
	t.Setenv("LC_ALL", "")
	if got := localeFromEnvironment(); got != ui.LocaleEnglish {
		t.Fatalf("localeFromEnvironment() = %q, want %q", got, ui.LocaleEnglish)
	}
}
