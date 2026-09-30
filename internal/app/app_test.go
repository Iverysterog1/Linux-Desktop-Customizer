package app

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/platform"
)

func TestNewRegistersKDEAdapterOnlyWhenSessionAndCommandsAreReady(t *testing.T) {
	cases := []struct {
		name          string
		desktop       string
		installRead   bool
		installWrite  bool
		wantKDE       bool
	}{
		{name: "KDE ready", desktop: "KDE", installRead: true, installWrite: true, wantKDE: true},
		{name: "KDE missing write", desktop: "KDE", installRead: true, installWrite: false, wantKDE: false},
		{name: "non KDE", desktop: "GNOME", installRead: true, installWrite: true, wantKDE: false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			binDir := filepath.Join(root, "bin")
			if err := os.MkdirAll(binDir, 0700); err != nil {
				t.Fatal(err)
			}
			marker := filepath.Join(root, "executed")
			if tc.installRead {
				writeFakeKDECommand(t, binDir, "kreadconfig6", marker)
			}
			if tc.installWrite {
				writeFakeKDECommand(t, binDir, "kwriteconfig6", marker)
			}

			t.Setenv("PATH", binDir)
			t.Setenv("XDG_CURRENT_DESKTOP", tc.desktop)
			t.Setenv("KDE_FULL_SESSION", "")
			t.Setenv("KDE_SESSION_VERSION", "6")
			t.Setenv("XDG_SESSION_TYPE", "wayland")
			t.Setenv("LTC_MANAGED_DIR", filepath.Join(root, "managed"))
			t.Setenv("LTC_STATE_DIR", filepath.Join(root, "state"))
			t.Setenv("XDG_DATA_DIRS", filepath.Join(root, "data"))

			rt, err := New()
			if err != nil {
				t.Fatalf("New() error = %v", err)
			}
			_, gotKDE := rt.Registry.Get("kde-config")
			if gotKDE != tc.wantKDE {
				t.Fatalf("kde-config registered = %v, want %v", gotKDE, tc.wantKDE)
			}
			if _, err := os.Stat(marker); !os.IsNotExist(err) {
				t.Fatalf("KDE command executed during registration; stat err = %v", err)
			}
		})
	}
}

func TestKDERuntimeReadyRequiresDetectionAndReadWriteCommands(t *testing.T) {
	status := platformStatus(true,
		command("kreadconfig6", true),
		command("kwriteconfig6", true),
	)
	if !kdeRuntimeReady(status) {
		t.Fatal("expected KDE runtime to be ready")
	}

	if kdeRuntimeReady(platformStatus(false,
		command("kreadconfig6", true),
		command("kwriteconfig6", true),
	)) {
		t.Fatal("non-KDE session must not be ready")
	}

	if kdeRuntimeReady(platformStatus(true,
		command("kreadconfig6", true),
		command("kwriteconfig6", false),
	)) {
		t.Fatal("KDE session without write command must not be ready")
	}
}

func writeFakeKDECommand(t *testing.T, dir, name, marker string) {
	t.Helper()
	path := filepath.Join(dir, name)
	content := "#!/bin/sh\nprintf '%s\\n' executed > " + shellQuote(marker) + "\n"
	if err := os.WriteFile(path, []byte(content), 0755); err != nil {
		t.Fatal(err)
	}
}

func shellQuote(value string) string {
	out := "'"
	for _, r := range value {
		if r == '\'' {
			out += "'\\''"
		} else {
			out += string(r)
		}
	}
	return out + "'"
}

func platformStatus(detected bool, commands ...platform.CommandStatus) platform.KDEStatus {
	return platform.KDEStatus{Detected: detected, Commands: commands}
}

func command(name string, available bool) platform.CommandStatus {
	return platform.CommandStatus{Name: name, Available: available}
}
