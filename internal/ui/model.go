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

type Model struct {
	Product  string   `json:"product"`
	Version  string   `json:"version"`
	Mode     string   `json:"mode"`
	Screens  []Screen `json:"screens"`
	Warnings []string `json:"warnings,omitempty"`
}

func FoundationModel(version string, registry *adapter.Registry) Model {
	m := Model{
		Product: "Linux Desktop Customizer",
		Version: version,
		Mode:    "rebuild-foundation",
		Screens: []Screen{
			{ID: "home", Title: "Home", Description: "Repository and runtime status", Enabled: true},
			{ID: "preview", Title: "Preview", Description: "Review declarative changes before apply", Enabled: true},
			{ID: "history", Title: "History & Rollback", Description: "Recover previously applied transactions", Enabled: true},
			{ID: "desktop", Title: "Desktop", Description: "Native desktop customization adapters", Enabled: false, Reason: "KDE/GNOME/XFCE mutation adapters are not integrated yet"},
			{ID: "themes", Title: "Themes", Description: "Theme creation/import", Enabled: false, Reason: "theme pipeline is not integrated yet"},
		},
	}
	if registry == nil || len(registry.All()) == 0 {
		m.Warnings = append(m.Warnings, "No mutation adapters are available.")
		return m
	}
	m.Warnings = append(m.Warnings,
		fmt.Sprintf("%d safe adapter(s) currently available; desktop-specific mutation remains gated.", len(registry.All())),
	)
	return m
}
