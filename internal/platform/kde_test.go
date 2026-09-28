package platform

import (
	"errors"
	"testing"
)

func TestProbeKDEPlasma6Wayland(t *testing.T) {
	t.Parallel()
	env := map[string]string{
		"XDG_CURRENT_DESKTOP": "KDE",
		"KDE_SESSION_VERSION": "6",
		"XDG_SESSION_TYPE":    "wayland",
	}
	paths := map[string]string{
		"kreadconfig6":    "/usr/bin/kreadconfig6",
		"kwriteconfig6":   "/usr/bin/kwriteconfig6",
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
	if !got.Capabilities[0].Supported || !got.Capabilities[1].Supported || !got.Capabilities[2].Supported {
		t.Fatalf("expected config/theme capabilities: %+v", got.Capabilities)
	}
	if got.Capabilities[3].Supported {
		t.Fatalf("qdbus6 should be unavailable: %+v", got.Capabilities[3])
	}
}

func TestProbeKDENonKDESessionIsUnsupported(t *testing.T) {
	t.Parallel()
	got := ProbeKDE(
		func(k string) string {
			if k == "XDG_CURRENT_DESKTOP" {
				return "GNOME"
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
