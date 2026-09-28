package transaction

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/adapter"
)

var ErrUnsupported = errors.New("transaction: plan contains unsupported changes")

type Engine struct {
	registry *adapter.Registry
	store    *Store
}

func NewEngine(registry *adapter.Registry, store *Store) (*Engine, error) {
	if registry == nil || store == nil {
		return nil, fmt.Errorf("transaction: registry and store are required")
	}
	return &Engine{registry: registry, store: store}, nil
}

func (e *Engine) Preview(ctx context.Context, plan Plan) (Preview, error) {
	if err := ValidatePlan(plan); err != nil {
		return Preview{}, err
	}

	out := Preview{Supported: true, Changes: make([]PreviewChange, 0, len(plan.Changes))}
	for i, change := range plan.Changes {
		pc := PreviewChange{Index: i, Change: change, Supported: true}

		a, ok := e.registry.Get(change.Adapter)
		if !ok {
			pc.Supported = false
			pc.Reason = fmt.Sprintf("adapter %q is not available", change.Adapter)
			out.Supported = false
			out.Changes = append(out.Changes, pc)
			continue
		}

		value, exists, err := a.Read(ctx, change.Key)
		if err != nil {
			return Preview{}, fmt.Errorf("transaction: read %s/%s: %w", change.Adapter, change.Key, err)
		}
		pc.Before = Snapshot{Exists: exists, Value: value}
		switch change.Action {
		case ActionSet:
			pc.After = Snapshot{Exists: true, Value: change.Value}
			pc.Noop = exists && value == change.Value
		case ActionUnset:
			pc.After = Snapshot{Exists: false}
			pc.Noop = !exists
		}
		out.Changes = append(out.Changes, pc)
	}
	return out, nil
}

func (e *Engine) Apply(ctx context.Context, plan Plan) (Transaction, error) {
	preview, err := e.Preview(ctx, plan)
	if err != nil {
		return Transaction{}, err
	}
	if !preview.Supported {
		return Transaction{}, ErrUnsupported
	}

	id, err := newID()
	if err != nil {
		return Transaction{}, err
	}
	tx := Transaction{
		ID:        id,
		CreatedAt: time.Now().UTC(),
		Status:    StatusApplying,
		Changes:   make([]RecordedChange, len(preview.Changes)),
	}
	for i, pc := range preview.Changes {
		tx.Changes[i] = RecordedChange{Change: pc.Change, Before: pc.Before, Noop: pc.Noop}
	}

	if err := e.store.Save(tx); err != nil {
		return tx, err
	}

	applied := make([]int, 0, len(tx.Changes))
	for i, rc := range tx.Changes {
		if rc.Noop {
			continue
		}
		a, _ := e.registry.Get(rc.Change.Adapter)
		switch rc.Change.Action {
		case ActionSet:
			err = a.Set(ctx, rc.Change.Key, rc.Change.Value)
		case ActionUnset:
			err = a.Unset(ctx, rc.Change.Key)
		}
		if err != nil {
			tx.Error = fmt.Sprintf("apply change %d: %v", i, err)
			if rollbackErr := e.restore(ctx, tx, applied); rollbackErr != nil {
				tx.Status = StatusRollbackFailed
				tx.Error += "; rollback: " + rollbackErr.Error()
			} else {
				tx.Status = StatusRolledBack
			}
			_ = e.store.Save(tx)
			return tx, fmt.Errorf("transaction: %s", tx.Error)
		}
		applied = append(applied, i)
	}

	tx.Status = StatusApplied
	if err := e.store.Save(tx); err != nil {
		tx.Error = "persist applied status: " + err.Error()
		if rollbackErr := e.restore(ctx, tx, applied); rollbackErr != nil {
			tx.Status = StatusRollbackFailed
			tx.Error += "; rollback: " + rollbackErr.Error()
		} else {
			tx.Status = StatusRolledBack
		}
		_ = e.store.Save(tx)
		return tx, fmt.Errorf("transaction: %s", tx.Error)
	}
	return tx, nil
}

func (e *Engine) Rollback(ctx context.Context, id string) (Transaction, error) {
	tx, err := e.store.Load(id)
	if err != nil {
		return Transaction{}, err
	}
	if tx.Status == StatusRolledBack {
		return tx, nil
	}
	switch tx.Status {
	case StatusApplied, StatusApplying, StatusRollbackFailed:
	default:
		return tx, fmt.Errorf("transaction: cannot roll back status %q", tx.Status)
	}

	indices := make([]int, 0, len(tx.Changes))
	for i, rc := range tx.Changes {
		if !rc.Noop {
			indices = append(indices, i)
		}
	}
	if err := e.restore(ctx, tx, indices); err != nil {
		tx.Status = StatusRollbackFailed
		tx.Error = err.Error()
		_ = e.store.Save(tx)
		return tx, err
	}
	tx.Status = StatusRolledBack
	tx.Error = ""
	if err := e.store.Save(tx); err != nil {
		return tx, err
	}
	return tx, nil
}

func (e *Engine) restore(ctx context.Context, tx Transaction, indices []int) error {
	var errs []error
	for i := len(indices) - 1; i >= 0; i-- {
		rc := tx.Changes[indices[i]]
		a, ok := e.registry.Get(rc.Change.Adapter)
		if !ok {
			errs = append(errs, fmt.Errorf("adapter %q unavailable during rollback", rc.Change.Adapter))
			continue
		}
		var err error
		if rc.Before.Exists {
			err = a.Set(ctx, rc.Change.Key, rc.Before.Value)
		} else {
			err = a.Unset(ctx, rc.Change.Key)
		}
		if err != nil {
			errs = append(errs, fmt.Errorf("restore %s/%s: %w", rc.Change.Adapter, rc.Change.Key, err))
		}
	}
	return errors.Join(errs...)
}

func validateChange(change Change) error {
	if change.Adapter == "" {
		return fmt.Errorf("adapter is required")
	}
	if change.Key == "" {
		return fmt.Errorf("key is required")
	}
	switch change.Action {
	case ActionSet:
		return nil
	case ActionUnset:
		if change.Value != "" {
			return fmt.Errorf("unset change must not contain a value")
		}
		return nil
	default:
		return fmt.Errorf("unsupported action %q", change.Action)
	}
}

func newID() (string, error) {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "", fmt.Errorf("transaction: generate id: %w", err)
	}
	return hex.EncodeToString(b[:]), nil
}
