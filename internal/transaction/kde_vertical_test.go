package transaction

import (
	"context"
	"testing"

	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/adapter"
	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/capability"
)

type memoryKDEAdapter struct {
	value  string
	exists bool
}

func (a *memoryKDEAdapter) Name() string { return "kde-config" }
func (a *memoryKDEAdapter) Capabilities(context.Context) []capability.Capability {
	return []capability.Capability{{ID: "kde.color-scheme", Supported: true}}
}
func (a *memoryKDEAdapter) Read(context.Context, string) (string, bool, error) {
	return a.value, a.exists, nil
}
func (a *memoryKDEAdapter) Set(_ context.Context, key, value string) error {
	if key != "color-scheme" { return ErrUnsupported }
	a.value, a.exists = value, true
	return nil
}
func (a *memoryKDEAdapter) Unset(_ context.Context, key string) error {
	if key != "color-scheme" { return ErrUnsupported }
	a.value, a.exists = "", false
	return nil
}

func TestKDEColorSchemeVerticalSlicePreviewApplyRollback(t *testing.T) {
	ctx := context.Background()
	kde := &memoryKDEAdapter{value: "BreezeLight", exists: true}
	registry, err := adapter.NewRegistry(kde)
	if err != nil { t.Fatal(err) }
	store, err := NewStore(t.TempDir())
	if err != nil { t.Fatal(err) }
	engine, err := NewEngine(registry, store)
	if err != nil { t.Fatal(err) }
	plan := Plan{Changes: []Change{{Adapter: "kde-config", Key: "color-scheme", Action: ActionSet, Value: "BreezeDark"}}}

	preview, err := engine.Preview(ctx, plan)
	if err != nil { t.Fatal(err) }
	if !preview.Supported || len(preview.Changes) != 1 { t.Fatalf("preview = %#v", preview) }
	change := preview.Changes[0]
	if !change.Before.Exists || change.Before.Value != "BreezeLight" || !change.After.Exists || change.After.Value != "BreezeDark" || change.Noop { t.Fatalf("preview change = %#v", change) }

	tx, err := engine.Apply(ctx, plan)
	if err != nil { t.Fatal(err) }
	if tx.Status != StatusApplied || !kde.exists || kde.value != "BreezeDark" { t.Fatalf("apply = %#v, adapter=%q/%v", tx, kde.value, kde.exists) }

	tx, err = engine.Rollback(ctx, tx.ID)
	if err != nil { t.Fatal(err) }
	if tx.Status != StatusRolledBack || !kde.exists || kde.value != "BreezeLight" { t.Fatalf("rollback = %#v, adapter=%q/%v", tx, kde.value, kde.exists) }

	tx2, err := engine.Rollback(ctx, tx.ID)
	if err != nil { t.Fatal(err) }
	if tx2.Status != StatusRolledBack || kde.value != "BreezeLight" { t.Fatalf("second rollback = %#v, adapter=%q", tx2, kde.value) }
}
