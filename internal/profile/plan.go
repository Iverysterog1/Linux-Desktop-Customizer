package profile

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/transaction"
)

// LoadFile reads a portable profile from a bounded, regular, non-symlink file.
// Parsing remains purely declarative; no adapter or transaction execution occurs.
func LoadFile(path string) (Profile, error) {
	if strings.TrimSpace(path) == "" {
		return Profile{}, fmt.Errorf("profile: path must not be empty")
	}

	info, err := os.Lstat(path)
	if err != nil {
		return Profile{}, fmt.Errorf("profile: inspect file: %w", err)
	}
	if info.Mode()&os.ModeSymlink != 0 {
		return Profile{}, fmt.Errorf("profile: symlink paths are not allowed")
	}
	if !info.Mode().IsRegular() {
		return Profile{}, fmt.Errorf("profile: path is not a regular file")
	}
	if info.Size() > int64(maxProfileBytes) {
		return Profile{}, fmt.Errorf("profile: document exceeds %d bytes", maxProfileBytes)
	}

	f, err := os.Open(path)
	if err != nil {
		return Profile{}, fmt.Errorf("profile: open file: %w", err)
	}
	defer f.Close()

	openedInfo, err := f.Stat()
	if err != nil {
		return Profile{}, fmt.Errorf("profile: stat opened file: %w", err)
	}
	if !openedInfo.Mode().IsRegular() || !os.SameFile(info, openedInfo) {
		return Profile{}, fmt.Errorf("profile: file changed while opening")
	}

	data, err := io.ReadAll(io.LimitReader(f, int64(maxProfileBytes)+1))
	if err != nil {
		return Profile{}, fmt.Errorf("profile: read file: %w", err)
	}
	if len(data) > maxProfileBytes {
		return Profile{}, fmt.Errorf("profile: document exceeds %d bytes", maxProfileBytes)
	}
	return Parse(data)
}

// ToPlan converts only a validated declarative profile into a transaction plan.
// The transaction engine still performs its own validation and preview before apply.
func ToPlan(p Profile) (transaction.Plan, error) {
	if err := Validate(p); err != nil {
		return transaction.Plan{}, err
	}

	plan := transaction.Plan{Changes: make([]transaction.Change, 0, len(p.Operations))}
	for i, op := range p.Operations {
		change := transaction.Change{
			Adapter: op.Adapter,
			Key:     op.Key,
			Value:   op.Value,
		}
		switch op.Action {
		case ActionSet:
			change.Action = transaction.ActionSet
		case ActionUnset:
			change.Action = transaction.ActionUnset
		default:
			return transaction.Plan{}, fmt.Errorf("profile: operation %d uses unsupported action %q", i, op.Action)
		}
		plan.Changes = append(plan.Changes, change)
	}
	return plan, nil
}
