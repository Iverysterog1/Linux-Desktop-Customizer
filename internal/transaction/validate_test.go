package transaction

import (
	"strings"
	"testing"
)

func TestValidatePlanRejectsEmpty(t *testing.T) {
	t.Parallel()
	if err := ValidatePlan(Plan{}); err == nil {
		t.Fatal("empty plan should be rejected")
	}
}

func TestValidatePlanRejectsDuplicateTarget(t *testing.T) {
	t.Parallel()
	plan := Plan{Changes: []Change{
		{Adapter: "local-file", Key: "theme/name", Action: ActionSet, Value: "a"},
		{Adapter: "local-file", Key: "theme/name", Action: ActionSet, Value: "b"},
	}}
	if err := ValidatePlan(plan); err == nil {
		t.Fatal("duplicate target should be rejected")
	}
}

func TestValidatePlanRejectsOversizedAndNULValues(t *testing.T) {
	t.Parallel()
	cases := []Plan{
		{Changes: []Change{{Adapter: "local-file", Key: "x", Action: ActionSet, Value: strings.Repeat("a", MaxValueLength+1)}}},
		{Changes: []Change{{Adapter: "local-file", Key: "x\x00y", Action: ActionSet, Value: "ok"}}},
	}
	for _, plan := range cases {
		if err := ValidatePlan(plan); err == nil {
			t.Fatalf("plan should be rejected: %+v", plan)
		}
	}
}

func TestValidatePlanAcceptsDistinctTargets(t *testing.T) {
	t.Parallel()
	plan := Plan{Changes: []Change{
		{Adapter: "local-file", Key: "theme/name", Action: ActionSet, Value: "dark"},
		{Adapter: "local-file", Key: "theme/icon", Action: ActionUnset},
	}}
	if err := ValidatePlan(plan); err != nil {
		t.Fatalf("valid plan rejected: %v", err)
	}
}
