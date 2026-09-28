package platform

import (
	"fmt"
	"strings"

	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/capability"
)

type CommandStatus struct {
	Name      string `json:"name"`
	Path      string `json:"path,omitempty"`
	Available bool   `json:"available"`
}

type KDEStatus struct {
	Detected       bool                    `json:"detected"`
	Desktop        string                  `json:"desktop,omitempty"`
	SessionVersion string                  `json:"session_version,omitempty"`
	SessionType    string                  `json:"session_type,omitempty"`
	Commands       []CommandStatus         `json:"commands"`
	Capabilities   []capability.Capability `json:"capabilities"`
}

// ProbeKDE is read-only. It does not execute KDE commands or mutate user state.
func ProbeKDE(getenv func(string) string, lookPath func(string) (string, error)) KDEStatus {
	desktop := getenv("XDG_CURRENT_DESKTOP")
	lower := strings.ToLower(desktop)
	detected := strings.Contains(lower, "kde") || strings.Contains(lower, "plasma") || getenv("KDE_FULL_SESSION") != ""

	status := KDEStatus{
		Detected:       detected,
		Desktop:        desktop,
		SessionVersion: getenv("KDE_SESSION_VERSION"),
		SessionType:    getenv("XDG_SESSION_TYPE"),
	}

	specs := []struct {
		name string
		id   string
	}{
		{"kreadconfig6", "kde.read-config"},
		{"kwriteconfig6", "kde.write-config"},
		{"lookandfeeltool6", "kde.global-theme"},
		{"qdbus6", "kde.dbus-reconfigure"},
	}

	for _, spec := range specs {
		path, err := lookPath(spec.name)
		available := err == nil && path != ""
		status.Commands = append(status.Commands, CommandStatus{
			Name: spec.name, Path: path, Available: available,
		})

		supported := detected && available
		reason := ""
		switch {
		case !detected:
			reason = "KDE Plasma session not detected"
		case !available:
			reason = fmt.Sprintf("%s not found in PATH", spec.name)
		default:
			reason = fmt.Sprintf("%s available; mutation remains disabled until a reviewed adapter is implemented", spec.name)
		}
		status.Capabilities = append(status.Capabilities, capability.Capability{
			ID: spec.id, Supported: supported, Reason: reason,
		})
	}

	return status
}
