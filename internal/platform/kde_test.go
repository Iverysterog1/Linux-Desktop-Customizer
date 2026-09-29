package platform

import (
	"errors"
	"testing"
)

func TestProbeKDEPlasma6WaylandFractionalScale(t *testing.T) {
	t.Parallel()
	env := map[string]string{
		"XDG_CURRENT_DESKTOP": "KDE",
		"KDE_SESSION_VERSION": "6",
		"XDG_SESSION_TYPE":    "wayland",
		"QT_VERSION":          "6.8",
		"QT_SCALE_FACTOR":     "1.25",
	}
	paths := map[string]string{
		"kreadconfig6":     "/usr/bin/kreadconfig6",
		"kwriteconfig6":    "/usr/bin/kwriteconfig6",
		"lookandfeeltool6": "/usr/bin/lookandfeeltool6",
	}
	got := ProbeKDE(
		func(k string) string { return env[k] },
		func(name string) (string, error) {
			if p := paths[name]; p != "" {
				return p, nil
			}
			return "", errors.New("not found")
		},
	)
	if !got.Detected || got.SessionVersion != "6" || got.SessionType != "wayland" {
		t.Fatalf("unexpected KDE status: %+v", got)
	}
	if got.Preflight.SessionProtocol != "wayland" {
		t.Fatalf("protocol = %q", got.Preflight.SessionProtocol)
	}
	if !got.Preflight.ScaleKnown || !got.Preflight.FractionalScaleDetected || got.Preflight.ScaleFactor != "1.25" {
		t.Fatalf("unexpected scale preflight: %+v", got.Preflight)
	}
	if got.Preflight.QtVersionHint != "6.8" {
		t.Fatalf("Qt version hint = %q", got.Preflight.QtVersionHint)
	}
	if !got.Capabilities[0].Supported || !got.Capabilities[1].Supported || !got.Capabilities[2].Supported {
		t.Fatalf("expected config/theme capabilities: %+v", got.Capabilities)
	}
	if got.Capabilities[3].Supported {
		t.Fatalf("qdbus6 should be unavailable: %+v", got.Capabilities[3])
	}
}

func TestProbeKDEX11Scale150(t *testing.T) {
	t.Parallel()
	env := map[string]string{
		"XDG_CURRENT_DESKTOP": "Plasma",
		"KDE_SESSION_VERSION": "6",
		"XDG_SESSION_TYPE":    "x11",
		"QT_VERSION":          "6.8",
		"QT_SCALE_FACTOR":     "1.5",
	}
	got := ProbeKDE(
		func(k string) string { return env[k] },
		func(name string) (string, error) { return "/usr/bin/" + name, nil },
	)
	if got.Preflight.SessionProtocol != "x11" {
		t.Fatalf("protocol = %q", got.Preflight.SessionProtocol)
	}
	if !got.Preflight.ScaleKnown || !got.Preflight.FractionalScaleDetected {
		t.Fatalf("expected explicit 150%% fractional scale: %+v", got.Preflight)
	}
}

func TestProbeKDEUnknownCompatibilityFactsStayUnknown(t *testing.T) {
	t.Parallel()
	env := map[string]string{
		"XDG_CURRENT_DESKTOP": "KDE",
		"KDE_SESSION_VERSION": "6",
		"XDG_SESSION_TYPE":    "tty",
	}
	got := ProbeKDE(
		func(k string) string { return env[k] },
		func(name string) (string, error) { return "/usr/bin/" + name, nil },
	)
	if got.Preflight.SessionProtocol != "unknown" {
		t.Fatalf("protocol = %q", got.Preflight.SessionProtocol)
	}
	if got.Preflight.ScaleKnown || got.Preflight.FractionalScaleDetected {
		t.Fatalf("scale should remain unknown: %+v", got.Preflight)
	}
	if len(got.Preflight.Warnings) < 3 {
		t.Fatalf("expected explicit unknown-state warnings: %+v", got.Preflight.Warnings)
	}
}

func TestProbeKDENonKDESessionIsUnsupported(t *testing.T) {
	t.Parallel()
	got := ProbeKDE(
		func(k string) string {
			if k == "XDG_CURRENT_DESKTOP" {
				return "GNOME"
			}
			if k == "XDG_SESSION_TYPE" {
				return "wayland"
			}
			return ""
		},
		func(name string) (string, error) { return "/usr/bin/" + name, nil },
	)
	if got.Detected {
		t.Fatalf("unexpected KDE detection: %+v", got)
	}
	for _, c := range got.Capabilities {
		if c.Supported {
			t.Fatalf("capability %s should not be supported outside KDE", c.ID)
		}
	}
}

func TestClassifyScale(t *testing.T) {
	t.Parallel()
	cases := []struct {
		raw        string
		known      bool
		fractional bool
	}{
		{"", false, false},
		{"invalid", false, false},
		{"1", true, false},
		{"2", true, false},
		{"1.25", true, true},
		{"1.5", true, true},
	}
	for _, tc := range cases {
		known, fractional := classifyScale(tc.raw)
		if known != tc.known || fractional != tc.fractional {
			t.Errorf("classifyScale(%q) = (%v,%v), want (%v,%v)", tc.raw, known, fractional, tc.known, tc.fractional)
		}
	}
}
