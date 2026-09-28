package transaction

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var transactionIDPattern = regexp.MustCompile(`^[a-f0-9]{32}$`)

type Store struct {
	root string
}

func NewStore(root string) (*Store, error) {
	if root == "" {
		return nil, fmt.Errorf("transaction store: empty root")
	}
	abs, err := filepath.Abs(root)
	if err != nil {
		return nil, fmt.Errorf("transaction store: resolve root: %w", err)
	}
	if err := os.MkdirAll(abs, 0o700); err != nil {
		return nil, fmt.Errorf("transaction store: create root: %w", err)
	}
	info, err := os.Lstat(abs)
	if err != nil {
		return nil, fmt.Errorf("transaction store: inspect root: %w", err)
	}
	if info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
		return nil, fmt.Errorf("transaction store: root must be a real directory")
	}
	resolved, err := filepath.EvalSymlinks(abs)
	if err != nil {
		return nil, fmt.Errorf("transaction store: resolve root symlinks: %w", err)
	}
	if filepath.Clean(resolved) != filepath.Clean(abs) {
		return nil, fmt.Errorf("transaction store: symlinked root is not allowed")
	}
	return &Store{root: abs}, nil
}

func (s *Store) Save(tx Transaction) error {
	if !transactionIDPattern.MatchString(tx.ID) {
		return fmt.Errorf("transaction store: invalid transaction id")
	}
	data, err := json.MarshalIndent(tx, "", "  ")
	if err != nil {
		return fmt.Errorf("transaction store: encode transaction: %w", err)
	}
	data = append(data, '\n')

	tmp, err := os.CreateTemp(s.root, ".transaction-*")
	if err != nil {
		return fmt.Errorf("transaction store: create temporary file: %w", err)
	}
	tmpName := tmp.Name()
	defer func() { _ = os.Remove(tmpName) }()

	if err := tmp.Chmod(0o600); err != nil {
		_ = tmp.Close()
		return fmt.Errorf("transaction store: chmod temporary file: %w", err)
	}
	if _, err := tmp.Write(data); err != nil {
		_ = tmp.Close()
		return fmt.Errorf("transaction store: write temporary file: %w", err)
	}
	if err := tmp.Sync(); err != nil {
		_ = tmp.Close()
		return fmt.Errorf("transaction store: sync temporary file: %w", err)
	}
	if err := tmp.Close(); err != nil {
		return fmt.Errorf("transaction store: close temporary file: %w", err)
	}
	if err := os.Rename(tmpName, s.path(tx.ID)); err != nil {
		return fmt.Errorf("transaction store: replace transaction: %w", err)
	}
	return nil
}

func (s *Store) Load(id string) (Transaction, error) {
	if !transactionIDPattern.MatchString(id) {
		return Transaction{}, fmt.Errorf("transaction store: invalid transaction id")
	}
	data, err := os.ReadFile(s.path(id))
	if errors.Is(err, os.ErrNotExist) {
		return Transaction{}, fmt.Errorf("transaction store: transaction %s not found", id)
	}
	if err != nil {
		return Transaction{}, fmt.Errorf("transaction store: read transaction: %w", err)
	}
	var tx Transaction
	if err := json.Unmarshal(data, &tx); err != nil {
		return Transaction{}, fmt.Errorf("transaction store: decode transaction: %w", err)
	}
	if tx.ID != id {
		return Transaction{}, fmt.Errorf("transaction store: transaction id mismatch")
	}
	return tx, nil
}

func (s *Store) path(id string) string {
	return filepath.Join(s.root, id+".json")
}
