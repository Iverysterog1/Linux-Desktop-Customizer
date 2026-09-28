package ui

import (
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
	for _, screen := range m.Screens {
		if screen.ID == "desktop" && screen.Enabled {
			t.Fatal("desktop screen must remain gated until a native adapter is integrated")
		}
	}
	if len(m.Warnings) == 0 {
		t.Fatal("foundation model should explain current adapter limits")
	}
}
