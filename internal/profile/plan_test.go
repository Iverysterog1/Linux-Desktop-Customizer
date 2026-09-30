package profile

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/transaction"
)

func TestToPlanMapsValidatedOperations(t *testing.T) {
	t.Parallel()
	p := Profile{
		Version: SchemaVersion,
		Name:    "Portable",
		Operations: []Operation{
			{Adapter: "kde-config", Action: ActionSet, Key: "color-scheme", Value: "BreezeDark"},
			{Adapter: "local-file", Action: ActionUnset, Key: "state/old"},
		},
	}

	plan, err := ToPlan(p)
	if err != nil {
		t.Fatalf("ToPlan() error = %v", err)
	}
	if len(plan.Changes) != 2 {
		t.Fatalf("changes = %d", len(plan.Changes))
	}
	if plan.Changes[0].Action != transaction.ActionSet || plan.Changes[0].Value != "BreezeDark" {
		t.Fatalf("unexpected first change: %+v", plan.Changes[0])
	}
	if plan.Changes[1].Action != transaction.ActionUnset || plan.Changes[1].Value != "" {
		t.Fatalf("unexpected second change: %+v", plan.Changes[1])
	}
}

func TestLoadFileUsesBoundedProfileParser(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	path := filepath.Join(dir, "profile.json")
	data := []byte(`{"version":1,"name":"Portable","operations":[{"adapter":"kde-config","action":"set","key":"color-scheme","value":"BreezeDark"}]}`)
	if err := os.WriteFile(path, data, 0600); err != nil {
		t.Fatal(err)
	}

	p, err := LoadFile(path)
	if err != nil {
		t.Fatalf("LoadFile() error = %v", err)
	}
	if p.Name != "Portable" {
		t.Fatalf("name = %q", p.Name)
	}
}

func TestLoadFileRejectsSymlinkAndNonRegularPaths(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	target := filepath.Join(dir, "target.json")
	if err := os.WriteFile(target, []byte(`{"version":1,"name":"Portable","operations":[{"adapter":"kde-config","action":"set","key":"color-scheme","value":"BreezeDark"}]}`), 0600); err != nil {
		t.Fatal(err)
	}
	link := filepath.Join(dir, "profile-link.json")
	if err := os.Symlink(target, link); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadFile(link); err == nil || !strings.Contains(err.Error(), "symlink") {
		t.Fatalf("symlink error = %v", err)
	}
	if _, err := LoadFile(dir); err == nil || !strings.Contains(err.Error(), "regular file") {
		t.Fatalf("directory error = %v", err)
	}
}

func TestLoadFileRejectsOversizedDocumentBeforeParse(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	path := filepath.Join(dir, "large.json")
	if err := os.WriteFile(path, []byte(strings.Repeat("x", maxProfileBytes+1)), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadFile(path); err == nil || !strings.Contains(err.Error(), "document exceeds") {
		t.Fatalf("oversized error = %v", err)
	}
}
