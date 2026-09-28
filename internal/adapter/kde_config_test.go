package adapter

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"testing"
)

type kdeRunnerCall struct {
	name string
	args []string
}

type fakeKDERunner struct {
	outputs []string
	errs    []error
	calls   []kdeRunnerCall
}

func (f *fakeKDERunner) Run(_ context.Context, name string, args ...string) (string, error) {
	f.calls = append(f.calls, kdeRunnerCall{name: name, args: append([]string(nil), args...)})
	i := len(f.calls) - 1
	var out string
	var err error
	if i < len(f.outputs) {
		out = f.outputs[i]
	}
	if i < len(f.errs) {
		err = f.errs[i]
	}
	return out, err
}

func newTestKDEAdapter(t *testing.T, runner KDERunner) *KDEConfigAdapter {
	t.Helper()
	a, err := newKDEConfigAdapter("/usr/bin/kreadconfig6", "/usr/bin/kwriteconfig6", runner)
	if err != nil {
		t.Fatal(err)
	}
	return a
}

func TestKDEConfigReadUsesFixedAllowlistedTarget(t *testing.T) {
	r := &fakeKDERunner{outputs: []string{"BreezeDark\n"}}
	a := newTestKDEAdapter(t, r)

	got, exists, err := a.Read(context.Background(), "color-scheme")
	if err != nil {
		t.Fatal(err)
	}
	if !exists || got != "BreezeDark\n" {
		// Runner contract returns command output as provided. Production runner
		// removes the process newline; fake output deliberately proves no shell
		// parsing or value rewriting occurs in the adapter.
		t.Fatalf("Read() = %q, %v; want raw fake output and exists", got, exists)
	}
	want := kdeRunnerCall{name: "/usr/bin/kreadconfig6", args: []string{
		"--file", "kdeglobals", "--group", "General", "--key", "ColorScheme", "--default", kdeMissingSentinel,
	}}
	if len(r.calls) != 1 || !reflect.DeepEqual(r.calls[0], want) {
		t.Fatalf("calls = %#v; want %#v", r.calls, want)
	}
}

func TestKDEConfigReadMissing(t *testing.T) {
	r := &fakeKDERunner{outputs: []string{kdeMissingSentinel}}
	a := newTestKDEAdapter(t, r)
	value, exists, err := a.Read(context.Background(), "color-scheme")
	if err != nil {
		t.Fatal(err)
	}
	if exists || value != "" {
		t.Fatalf("Read() = %q, %v; want missing", value, exists)
	}
}

func TestKDEConfigRejectsUnknownKeyWithoutExecution(t *testing.T) {
	r := &fakeKDERunner{}
	a := newTestKDEAdapter(t, r)
	if _, _, err := a.Read(context.Background(), "../../evil"); err == nil {
		t.Fatal("Read() accepted unknown key")
	}
	if err := a.Set(context.Background(), "arbitrary", "value"); err == nil {
		t.Fatal("Set() accepted unknown key")
	}
	if err := a.Unset(context.Background(), "arbitrary"); err == nil {
		t.Fatal("Unset() accepted unknown key")
	}
	if len(r.calls) != 0 {
		t.Fatalf("unexpected command execution: %#v", r.calls)
	}
}

func TestKDEConfigSetIsIdempotent(t *testing.T) {
	r := &fakeKDERunner{outputs: []string{"BreezeDark"}}
	a := newTestKDEAdapter(t, r)
	if err := a.Set(context.Background(), "color-scheme", "BreezeDark"); err != nil {
		t.Fatal(err)
	}
	if len(r.calls) != 1 {
		t.Fatalf("got %d calls; want read only", len(r.calls))
	}
}

func TestKDEConfigSetWritesFixedArguments(t *testing.T) {
	r := &fakeKDERunner{outputs: []string{"BreezeLight", ""}}
	a := newTestKDEAdapter(t, r)
	if err := a.Set(context.Background(), "color-scheme", "BreezeDark"); err != nil {
		t.Fatal(err)
	}
	if len(r.calls) != 2 {
		t.Fatalf("got %d calls; want read + write", len(r.calls))
	}
	want := kdeRunnerCall{name: "/usr/bin/kwriteconfig6", args: []string{
		"--file", "kdeglobals", "--group", "General", "--key", "ColorScheme", "BreezeDark",
	}}
	if !reflect.DeepEqual(r.calls[1], want) {
		t.Fatalf("write call = %#v; want %#v", r.calls[1], want)
	}
}

func TestKDEConfigUnsetIsIdempotent(t *testing.T) {
	r := &fakeKDERunner{outputs: []string{kdeMissingSentinel}}
	a := newTestKDEAdapter(t, r)
	if err := a.Unset(context.Background(), "color-scheme"); err != nil {
		t.Fatal(err)
	}
	if len(r.calls) != 1 {
		t.Fatalf("got %d calls; want read only", len(r.calls))
	}
}

func TestKDEConfigUnsetDeletesExistingValue(t *testing.T) {
	r := &fakeKDERunner{outputs: []string{"BreezeDark", ""}}
	a := newTestKDEAdapter(t, r)
	if err := a.Unset(context.Background(), "color-scheme"); err != nil {
		t.Fatal(err)
	}
	if len(r.calls) != 2 || r.calls[1].args[len(r.calls[1].args)-1] != "--delete" {
		t.Fatalf("unexpected calls: %#v", r.calls)
	}
}

func TestKDEConfigValueValidation(t *testing.T) {
	r := &fakeKDERunner{}
	a := newTestKDEAdapter(t, r)
	for _, value := range []string{"", "bad\nvalue", "bad\x00value", strings.Repeat("x", 257), kdeMissingSentinel} {
		if err := a.Set(context.Background(), "color-scheme", value); err == nil {
			t.Fatalf("Set() accepted invalid value %q", value)
		}
	}
	if len(r.calls) != 0 {
		t.Fatalf("validation should run before command execution: %#v", r.calls)
	}
}

func TestKDEConfigPropagatesRunnerError(t *testing.T) {
	r := &fakeKDERunner{errs: []error{errors.New("boom")}}
	a := newTestKDEAdapter(t, r)
	if _, _, err := a.Read(context.Background(), "color-scheme"); err == nil || !strings.Contains(err.Error(), "boom") {
		t.Fatalf("Read() error = %v; want wrapped runner error", err)
	}
}
