package adapter

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestFileAdapterSetReadUnset(t *testing.T) {
	t.Parallel()
	a, err := NewFileAdapter(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	if err := a.Set(ctx, "theme/name", "dark"); err != nil {
		t.Fatal(err)
	}
	got, exists, err := a.Read(ctx, "theme/name")
	if err != nil {
		t.Fatal(err)
	}
	if !exists || got != "dark" {
		t.Fatalf("read = (%q, %v), want (%q, true)", got, exists, "dark")
	}

	if err := a.Set(ctx, "theme/name", "dark"); err != nil {
		t.Fatalf("idempotent set failed: %v", err)
	}
	if err := a.Unset(ctx, "theme/name"); err != nil {
		t.Fatal(err)
	}
	if err := a.Unset(ctx, "theme/name"); err != nil {
		t.Fatalf("idempotent unset failed: %v", err)
	}
	_, exists, err = a.Read(ctx, "theme/name")
	if err != nil {
		t.Fatal(err)
	}
	if exists {
		t.Fatal("managed file still exists after unset")
	}
}

func TestFileAdapterRejectsTraversal(t *testing.T) {
	t.Parallel()
	a, err := NewFileAdapter(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	if err := a.Set(context.Background(), "../escape", "no"); err == nil {
		t.Fatal("expected traversal to be rejected")
	}
	if err := a.Set(context.Background(), "/absolute", "no"); err == nil {
		t.Fatal("expected absolute path to be rejected")
	}
}

func TestFileAdapterRejectsSymlinkEscape(t *testing.T) {
	t.Parallel()
	root := t.TempDir()
	outside := t.TempDir()

	a, err := NewFileAdapter(root)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(outside, filepath.Join(root, "link")); err != nil {
		t.Fatal(err)
	}
	if err := a.Set(context.Background(), "link/pwned", "no"); err == nil {
		t.Fatal("expected symlink traversal to be rejected")
	}
	if _, err := os.Stat(filepath.Join(outside, "pwned")); !os.IsNotExist(err) {
		t.Fatalf("outside target unexpectedly exists: %v", err)
	}
}
