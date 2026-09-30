package ui

import (
	"fmt"

	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/adapter"
)

type Screen struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Reason      string `json:"reason,omitempty"`
}

type Accessibility struct {
	ReducedMotionSupported      bool `json:"reduced_motion_supported"`
	ReadableTextSupported       bool `json:"readable_text_supported"`
	HighContrastSupported       bool `json:"high_contrast_supported"`
	KeyboardNavigationSupported bool `json:"keyboard_navigation_supported"`
}

type Model struct {
	Product       string        `json:"product"`
	Version       string        `json:"version"`
	Mode          string        `json:"mode"`
	Locale        string        `json:"locale"`
	BrandTagline  string        `json:"brand_tagline"`
	Screens       []Screen      `json:"screens"`
	Accessibility Accessibility `json:"accessibility"`
	Warnings      []string      `json:"warnings,omitempty"`
}

// FoundationModel preserves the original English default for existing callers.
func FoundationModel(version string, registry *adapter.Registry) Model {
	return FoundationModelForLocale(version, registry, LocaleEnglish)
}

func FoundationModelForLocale(version string, registry *adapter.Registry, locale string) Model {
	locale = NormalizeLocale(locale)
	kdeAvailable := adapterAvailable(registry, "kde-config")
	desktopDescription := message(locale, "screen.desktop.description")
	desktopReason := message(locale, "screen.desktop.reason")
	if kdeAvailable {
		desktopDescription = message(locale, "screen.desktop.description.kde")
		desktopReason = ""
	}

	m := Model{
		Product:      "Linux Desktop Customizer",
		Version:      version,
		Mode:         "rebuild-foundation",
		Locale:       locale,
		BrandTagline: message(locale, "splash.tagline"),
		Screens: []Screen{
			{ID: "home", Title: message(locale, "screen.home.title"), Description: message(locale, "screen.home.description"), Enabled: true},
			{ID: "preview", Title: message(locale, "screen.preview.title"), Description: message(locale, "screen.preview.description"), Enabled: true},
			{ID: "history", Title: message(locale, "screen.history.title"), Description: message(locale, "screen.history.description"), Enabled: true},
			{ID: "desktop", Title: message(locale, "screen.desktop.title"), Description: desktopDescription, Enabled: kdeAvailable, Reason: desktopReason},
			{ID: "themes", Title: message(locale, "screen.themes.title"), Description: message(locale, "screen.themes.description"), Enabled: false, Reason: message(locale, "screen.themes.reason")},
		},
		Accessibility: Accessibility{
			ReducedMotionSupported:      true,
			ReadableTextSupported:       true,
			HighContrastSupported:       true,
			KeyboardNavigationSupported: true,
		},
	}
	if registry == nil || len(registry.All()) == 0 {
		m.Warnings = append(m.Warnings, message(locale, "warning.no_adapters"))
		return m
	}
	m.Warnings = append(m.Warnings,
		fmt.Sprintf(message(locale, "warning.safe_adapters"), len(registry.All())),
	)
	return m
}

func adapterAvailable(registry *adapter.Registry, name string) bool {
	if registry == nil {
		return false
	}
	_, ok := registry.Get(name)
	return ok
}
