package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/transaction"
)

func TestLoadPlanUsesPortableProfileSchema(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	path := filepath.Join(dir, "profile.json")
	data := []byte(`{"version":1,"name":"Portable","operations":[{"adapter":"kde-config","action":"set","key":"color-scheme","value":"BreezeDark"}]}`)
	if err := os.WriteFile(path, data, 0600); err != nil {
		t.Fatal(err)
	}

	plan, err := loadPlan(path)
	if err != nil {
		t.Fatalf("loadPlan() error = %v", err)
	}
	if len(plan.Changes) != 1 {
		t.Fatalf("changes = %d", len(plan.Changes))
	}
	if plan.Changes[0].Adapter != "kde-config" ||
		plan.Changes[0].Action != transaction.ActionSet ||
		plan.Changes[0].Key != "color-scheme" ||
		plan.Changes[0].Value != "BreezeDark" {
		t.Fatalf("unexpected change: %+v", plan.Changes[0])
	}
}

func TestLoadPlanRejectsLegacyRawTransactionPlan(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	path := filepath.Join(dir, "legacy.json")
	data := []byte(`{"changes":[{"adapter":"kde-config","key":"color-scheme","action":"set","value":"BreezeDark"}]}`)
	if err := os.WriteFile(path, data, 0600); err != nil {
		t.Fatal(err)
	}

	_, err := loadPlan(path)
	if err == nil || !strings.Contains(err.Error(), "unknown field") {
		t.Fatalf("legacy plan error = %v", err)
	}
}

func TestLoadPlanRejectsExecutableField(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	path := filepath.Join(dir, "unsafe.json")
	data := []byte(`{"version":1,"name":"Unsafe","script":"echo pwn","operations":[{"adapter":"kde-config","action":"set","key":"color-scheme","value":"BreezeDark"}]}`)
	if err := os.WriteFile(path, data, 0600); err != nil {
		t.Fatal(err)
	}

	_, err := loadPlan(path)
	if err == nil || !strings.Contains(err.Error(), "unknown field") {
		t.Fatalf("unsafe profile error = %v", err)
	}
}
