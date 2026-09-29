package platform

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/capability"
)

type CommandStatus struct {
	Name      string `json:"name"`
	Path      string `json:"path,omitempty"`
	Available bool   `json:"available"`
}

type KDEPreflight struct {
	SessionProtocol           string   `json:"session_protocol"`
	PlasmaVersion             string   `json:"plasma_version,omitempty"`
	QtVersionHint             string   `json:"qt_version_hint,omitempty"`
	ScaleFactor               string   `json:"scale_factor,omitempty"`
	ScaleKnown                bool     `json:"scale_known"`
	FractionalScaleDetected   bool     `json:"fractional_scale_detected"`
	Warnings                  []string `json:"warnings,omitempty"`
}

type KDEStatus struct {
	Detected       bool                    `json:"detected"`
	Desktop        string                  `json:"desktop,omitempty"`
	SessionVersion string                  `json:"session_version,omitempty"`
	SessionType    string                  `json:"session_type,omitempty"`
	Preflight      KDEPreflight             `json:"preflight"`
	Commands       []CommandStatus          `json:"commands"`
	Capabilities   []capability.Capability  `json:"capabilities"`
}

// ProbeKDE is read-only. It does not execute KDE commands or mutate user state.
func ProbeKDE(getenv func(string) string, lookPath func(string) (string, error)) KDEStatus {
	desktop := getenv("XDG_CURRENT_DESKTOP")
	lower := strings.ToLower(desktop)
	detected := strings.Contains(lower, "kde") || strings.Contains(lower, "plasma") || getenv("KDE_FULL_SESSION") != ""

	sessionType := strings.ToLower(strings.TrimSpace(getenv("XDG_SESSION_TYPE")))
	protocol := "unknown"
	switch sessionType {
	case "wayland", "x11":
		protocol = sessionType
	}

	scaleRaw := firstNonEmpty(
		strings.TrimSpace(getenv("QT_SCALE_FACTOR")),
		strings.TrimSpace(getenv("GDK_SCALE")),
	)
	scaleKnown, fractional := classifyScale(scaleRaw)

	preflight := KDEPreflight{
		SessionProtocol:         protocol,
		PlasmaVersion:           strings.TrimSpace(getenv("KDE_SESSION_VERSION")),
		QtVersionHint:           strings.TrimSpace(getenv("QT_VERSION")),
		ScaleFactor:             scaleRaw,
		ScaleKnown:              scaleKnown,
		FractionalScaleDetected: fractional,
	}
	if protocol == "unknown" {
		preflight.Warnings = append(preflight.Warnings, "display protocol is not exposed as Wayland or X11")
	}
	if !scaleKnown {
		preflight.Warnings = append(preflight.Warnings, "scale factor is not explicitly exposed; fractional-scaling compatibility remains unknown")
	}
	if preflight.QtVersionHint == "" {
		preflight.Warnings = append(preflight.Warnings, "Qt version is not exposed by the environment; compatibility remains unknown")
	}

	status := KDEStatus{
		Detected:       detected,
		Desktop:        desktop,
		SessionVersion: preflight.PlasmaVersion,
		SessionType:    sessionType,
		Preflight:      preflight,
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

func classifyScale(raw string) (known bool, fractional bool) {
	if raw == "" {
		return false, false
	}
	v, err := strconv.ParseFloat(raw, 64)
	if err != nil || v <= 0 {
		return false, false
	}
	return true, v != float64(int(v))
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}
