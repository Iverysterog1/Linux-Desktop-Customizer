package adapter

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/capability"
)

// FileAdapter owns a deliberately narrow generated layer. Keys are relative
// paths below root; absolute paths, traversal and symlink traversal are denied.
type FileAdapter struct {
	root string
}

func NewFileAdapter(root string) (*FileAdapter, error) {
	if root == "" {
		return nil, fmt.Errorf("file adapter: empty root")
	}
	abs, err := filepath.Abs(root)
	if err != nil {
		return nil, fmt.Errorf("file adapter: resolve root: %w", err)
	}
	if err := os.MkdirAll(abs, 0o700); err != nil {
		return nil, fmt.Errorf("file adapter: create root: %w", err)
	}
	info, err := os.Lstat(abs)
	if err != nil {
		return nil, fmt.Errorf("file adapter: inspect root: %w", err)
	}
	if info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
		return nil, fmt.Errorf("file adapter: root must be a real directory")
	}
	resolved, err := filepath.EvalSymlinks(abs)
	if err != nil {
		return nil, fmt.Errorf("file adapter: resolve root symlinks: %w", err)
	}
	if filepath.Clean(resolved) != filepath.Clean(abs) {
		return nil, fmt.Errorf("file adapter: symlinked root is not allowed")
	}
	return &FileAdapter{root: abs}, nil
}

func (a *FileAdapter) Name() string { return "local-file" }

func (a *FileAdapter) Capabilities(context.Context) []capability.Capability {
	return []capability.Capability{{
		ID:        "managed-file-layer",
		Supported: true,
		Reason:    "writes are confined to the Customizer-owned generated layer",
	}}
}

func (a *FileAdapter) Read(_ context.Context, key string) (string, bool, error) {
	target, err := a.safePath(key)
	if err != nil {
		return "", false, err
	}
	info, err := os.Lstat(target)
	if errors.Is(err, os.ErrNotExist) {
		return "", false, nil
	}
	if err != nil {
		return "", false, fmt.Errorf("file adapter: inspect %q: %w", key, err)
	}
	if info.Mode()&os.ModeSymlink != 0 || !info.Mode().IsRegular() {
		return "", false, fmt.Errorf("file adapter: %q is not a regular managed file", key)
	}
	b, err := os.ReadFile(target)
	if err != nil {
		return "", false, fmt.Errorf("file adapter: read %q: %w", key, err)
	}
	return string(b), true, nil
}

func (a *FileAdapter) Set(_ context.Context, key, value string) error {
	target, err := a.safePath(key)
	if err != nil {
		return err
	}
	if err := a.ensureSafeParent(filepath.Dir(target)); err != nil {
		return err
	}
	if info, err := os.Lstat(target); err == nil {
		if info.Mode()&os.ModeSymlink != 0 || !info.Mode().IsRegular() {
			return fmt.Errorf("file adapter: refusing non-regular target %q", key)
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("file adapter: inspect target %q: %w", key, err)
	}

	tmp, err := os.CreateTemp(filepath.Dir(target), ".ltc-write-*")
	if err != nil {
		return fmt.Errorf("file adapter: create temporary file: %w", err)
	}
	tmpName := tmp.Name()
	defer func() { _ = os.Remove(tmpName) }()

	if err := tmp.Chmod(0o600); err != nil {
		_ = tmp.Close()
		return fmt.Errorf("file adapter: chmod temporary file: %w", err)
	}
	if _, err := tmp.WriteString(value); err != nil {
		_ = tmp.Close()
		return fmt.Errorf("file adapter: write temporary file: %w", err)
	}
	if err := tmp.Sync(); err != nil {
		_ = tmp.Close()
		return fmt.Errorf("file adapter: sync temporary file: %w", err)
	}
	if err := tmp.Close(); err != nil {
		return fmt.Errorf("file adapter: close temporary file: %w", err)
	}
	if err := os.Rename(tmpName, target); err != nil {
		return fmt.Errorf("file adapter: atomically replace %q: %w", key, err)
	}
	return nil
}

func (a *FileAdapter) Unset(_ context.Context, key string) error {
	target, err := a.safePath(key)
	if err != nil {
		return err
	}
	info, err := os.Lstat(target)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("file adapter: inspect target %q: %w", key, err)
	}
	if info.Mode()&os.ModeSymlink != 0 || !info.Mode().IsRegular() {
		return fmt.Errorf("file adapter: refusing non-regular target %q", key)
	}
	if err := os.Remove(target); err != nil {
		return fmt.Errorf("file adapter: remove %q: %w", key, err)
	}
	return nil
}

func (a *FileAdapter) safePath(key string) (string, error) {
	if key == "" || filepath.IsAbs(key) {
		return "", fmt.Errorf("file adapter: key must be a non-empty relative path")
	}
	clean := filepath.Clean(filepath.FromSlash(key))
	if clean == "." || clean == ".." || strings.HasPrefix(clean, ".."+string(filepath.Separator)) {
		return "", fmt.Errorf("file adapter: unsafe key %q", key)
	}
	target := filepath.Join(a.root, clean)
	rel, err := filepath.Rel(a.root, target)
	if err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return "", fmt.Errorf("file adapter: key escapes managed root")
	}

	current := a.root
	parts := strings.Split(rel, string(filepath.Separator))
	for i, part := range parts {
		if part == "" {
			continue
		}
		current = filepath.Join(current, part)
		info, err := os.Lstat(current)
		if errors.Is(err, os.ErrNotExist) {
			break
		}
		if err != nil {
			return "", fmt.Errorf("file adapter: inspect path component: %w", err)
		}
		if info.Mode()&os.ModeSymlink != 0 {
			return "", fmt.Errorf("file adapter: symlink path component rejected")
		}
		if i < len(parts)-1 && !info.IsDir() {
			return "", fmt.Errorf("file adapter: non-directory path component rejected")
		}
	}
	return target, nil
}

func (a *FileAdapter) ensureSafeParent(parent string) error {
	rel, err := filepath.Rel(a.root, parent)
	if err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return fmt.Errorf("file adapter: parent escapes managed root")
	}
	current := a.root
	if rel == "." {
		return nil
	}
	for _, part := range strings.Split(rel, string(filepath.Separator)) {
		if part == "" {
			continue
		}
		current = filepath.Join(current, part)
		info, err := os.Lstat(current)
		if errors.Is(err, os.ErrNotExist) {
			if err := os.Mkdir(current, 0o700); err != nil && !errors.Is(err, os.ErrExist) {
				return fmt.Errorf("file adapter: create managed directory: %w", err)
			}
			continue
		}
		if err != nil {
			return fmt.Errorf("file adapter: inspect managed directory: %w", err)
		}
		if info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
			return fmt.Errorf("file adapter: unsafe managed directory component")
		}
	}
	return nil
}
