package transaction

import (
	"context"
	"errors"
	"testing"

	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/adapter"
	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/capability"
)

func newTestEngine(t *testing.T, extra ...adapter.Adapter) (*Engine, *adapter.FileAdapter) {
	t.Helper()
	fileAdapter, err := adapter.NewFileAdapter(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	adapters := append([]adapter.Adapter{fileAdapter}, extra...)
	registry, err := adapter.NewRegistry(adapters...)
	if err != nil {
		t.Fatal(err)
	}
	store, err := NewStore(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	engine, err := NewEngine(registry, store)
	if err != nil {
		t.Fatal(err)
	}
	return engine, fileAdapter
}

func TestApplyAndRollback(t *testing.T) {
	t.Parallel()
	engine, managed := newTestEngine(t)
	ctx := context.Background()
	plan := Plan{Changes: []Change{{
		Adapter: "local-file",
		Key:     "theme/name",
		Action:  ActionSet,
		Value:   "dark",
	}}}

	preview, err := engine.Preview(ctx, plan)
	if err != nil {
		t.Fatal(err)
	}
	if !preview.Supported || preview.Changes[0].Noop {
		t.Fatalf("unexpected preview: %+v", preview)
	}

	tx, err := engine.Apply(ctx, plan)
	if err != nil {
		t.Fatal(err)
	}
	if tx.Status != StatusApplied {
		t.Fatalf("status = %q, want %q", tx.Status, StatusApplied)
	}
	got, exists, err := managed.Read(ctx, "theme/name")
	if err != nil || !exists || got != "dark" {
		t.Fatalf("managed state = (%q, %v, %v)", got, exists, err)
	}

	tx, err = engine.Rollback(ctx, tx.ID)
	if err != nil {
		t.Fatal(err)
	}
	if tx.Status != StatusRolledBack {
		t.Fatalf("rollback status = %q", tx.Status)
	}
	_, exists, err = managed.Read(ctx, "theme/name")
	if err != nil {
		t.Fatal(err)
	}
	if exists {
		t.Fatal("managed state was not restored")
	}

	if _, err := engine.Rollback(ctx, tx.ID); err != nil {
		t.Fatalf("second rollback failed: %v", err)
	}
}

func TestApplySameValueIsNoop(t *testing.T) {
	t.Parallel()
	engine, managed := newTestEngine(t)
	ctx := context.Background()
	if err := managed.Set(ctx, "theme/name", "dark"); err != nil {
		t.Fatal(err)
	}
	plan := Plan{Changes: []Change{{
		Adapter: "local-file",
		Key:     "theme/name",
		Action:  ActionSet,
		Value:   "dark",
	}}}
	tx, err := engine.Apply(ctx, plan)
	if err != nil {
		t.Fatal(err)
	}
	if !tx.Changes[0].Noop {
		t.Fatal("same desired value should be a no-op")
	}
}

func TestUnsupportedAdapterIsReportedBeforeMutation(t *testing.T) {
	t.Parallel()
	engine, _ := newTestEngine(t)
	ctx := context.Background()
	plan := Plan{Changes: []Change{{
		Adapter: "does-not-exist",
		Key:     "anything",
		Action:  ActionSet,
		Value:   "x",
	}}}
	preview, err := engine.Preview(ctx, plan)
	if err != nil {
		t.Fatal(err)
	}
	if preview.Supported || preview.Changes[0].Supported {
		t.Fatalf("unsupported adapter unexpectedly supported: %+v", preview)
	}
	if _, err := engine.Apply(ctx, plan); !errors.Is(err, ErrUnsupported) {
		t.Fatalf("apply error = %v, want ErrUnsupported", err)
	}
}

func TestPartialFailureRollsBackEarlierChanges(t *testing.T) {
	t.Parallel()
	fail := &failingAdapter{}
	engine, managed := newTestEngine(t, fail)
	ctx := context.Background()
	plan := Plan{Changes: []Change{
		{Adapter: "local-file", Key: "theme/name", Action: ActionSet, Value: "dark"},
		{Adapter: "fail", Key: "boom", Action: ActionSet, Value: "x"},
	}}

	tx, err := engine.Apply(ctx, plan)
	if err == nil {
		t.Fatal("expected apply failure")
	}
	if tx.Status != StatusRolledBack {
		t.Fatalf("status = %q, want %q; err=%v", tx.Status, StatusRolledBack, err)
	}
	_, exists, readErr := managed.Read(ctx, "theme/name")
	if readErr != nil {
		t.Fatal(readErr)
	}
	if exists {
		t.Fatal("first change survived a later apply failure")
	}
}

func TestValidationMismatchRollsBackCurrentAndEarlierChanges(t *testing.T) {
	t.Parallel()
	validation := &validationAdapter{before: "before", exists: true, mode: validationMismatch}
	engine, managed := newTestEngine(t, validation)
	ctx := context.Background()
	plan := Plan{Changes: []Change{
		{Adapter: "local-file", Key: "theme/name", Action: ActionSet, Value: "dark"},
		{Adapter: "validation", Key: "setting", Action: ActionSet, Value: "after"},
	}}

	tx, err := engine.Apply(ctx, plan)
	if err == nil {
		t.Fatal("expected validation mismatch")
	}
	if tx.Status != StatusRolledBack {
		t.Fatalf("status = %q, want %q; err=%v", tx.Status, StatusRolledBack, err)
	}
	if validation.value != "before" || !validation.exists {
		t.Fatalf("current change not restored: value=%q exists=%v", validation.value, validation.exists)
	}
	_, exists, readErr := managed.Read(ctx, "theme/name")
	if readErr != nil {
		t.Fatal(readErr)
	}
	if exists {
		t.Fatal("earlier change survived validation mismatch")
	}
}

func TestValidationReadFailureRollsBackCurrentAndEarlierChanges(t *testing.T) {
	t.Parallel()
	validation := &validationAdapter{before: "before", exists: true, mode: validationReadFailure}
	engine, managed := newTestEngine(t, validation)
	ctx := context.Background()
	plan := Plan{Changes: []Change{
		{Adapter: "local-file", Key: "theme/name", Action: ActionSet, Value: "dark"},
		{Adapter: "validation", Key: "setting", Action: ActionSet, Value: "after"},
	}}

	tx, err := engine.Apply(ctx, plan)
	if err == nil {
		t.Fatal("expected validation read failure")
	}
	if tx.Status != StatusRolledBack {
		t.Fatalf("status = %q, want %q; err=%v", tx.Status, StatusRolledBack, err)
	}
	if validation.value != "before" || !validation.exists {
		t.Fatalf("current change not restored: value=%q exists=%v", validation.value, validation.exists)
	}
	_, exists, readErr := managed.Read(ctx, "theme/name")
	if readErr != nil {
		t.Fatal(readErr)
	}
	if exists {
		t.Fatal("earlier change survived validation read failure")
	}
}

type failingAdapter struct{}

func (*failingAdapter) Name() string { return "fail" }
func (*failingAdapter) Capabilities(context.Context) []capability.Capability {
	return []capability.Capability{{ID: "test-failure", Supported: true}}
}
func (*failingAdapter) Read(context.Context, string) (string, bool, error) {
	return "", false, nil
}
func (*failingAdapter) Set(context.Context, string, string) error {
	return errors.New("injected failure")
}
func (*failingAdapter) Unset(context.Context, string) error {
	return errors.New("injected failure")
}

type validationMode int

const (
	validationMismatch validationMode = iota
	validationReadFailure
)

type validationAdapter struct {
	before string
	value  string
	exists bool
	set    bool
	mode   validationMode
}

func (*validationAdapter) Name() string { return "validation" }
func (*validationAdapter) Capabilities(context.Context) []capability.Capability {
	return []capability.Capability{{ID: "test-validation", Supported: true}}
}
func (a *validationAdapter) Read(context.Context, string) (string, bool, error) {
	if a.set {
		switch a.mode {
		case validationReadFailure:
			a.set = false
			return "", false, errors.New("injected validation read failure")
		case validationMismatch:
			return "wrong", true, nil
		}
	}
	if a.value == "" && a.exists {
		return a.before, true, nil
	}
	return a.value, a.exists, nil
}
func (a *validationAdapter) Set(_ context.Context, _ string, value string) error {
	if a.set {
		a.value = value
		a.exists = true
		a.set = false
		return nil
	}
	if value == a.before {
		a.value = value
		a.exists = true
		return nil
	}
	a.value = value
	a.exists = true
	a.set = true
	return nil
}
func (a *validationAdapter) Unset(context.Context, string) error {
	a.value = ""
	a.exists = false
	a.set = false
	return nil
}
