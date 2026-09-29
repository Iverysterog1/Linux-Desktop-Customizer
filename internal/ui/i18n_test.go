package ui

import "testing"

func TestNormalizeLocale(t *testing.T) {
	t.Parallel()
	cases := map[string]string{
		"": LocaleEnglish,
		"en": LocaleEnglish,
		"en-US": LocaleEnglish,
		"pt": LocalePortuguese,
		"pt_PT": LocalePortuguese,
		"pt-PT.UTF-8": LocalePortuguese,
		"fr-FR": LocaleEnglish,
	}
	for input, want := range cases {
		if got := NormalizeLocale(input); got != want {
			t.Errorf("NormalizeLocale(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestMessageFallsBackToEnglish(t *testing.T) {
	t.Parallel()
	if got := message("fr-FR", "screen.home.title"); got != "Home" {
		t.Fatalf("fallback = %q, want Home", got)
	}
}

func TestBilingualProductStringsStayInParity(t *testing.T) {
	t.Parallel()
	keys := []string{
		"splash.tagline",
		"a11y.high_contrast",
		"a11y.keyboard",
	}
	for _, key := range keys {
		if catalog[LocaleEnglish][key] == "" {
			t.Errorf("missing English string for %q", key)
		}
		if catalog[LocalePortuguese][key] == "" {
			t.Errorf("missing pt-PT string for %q", key)
		}
	}
}
