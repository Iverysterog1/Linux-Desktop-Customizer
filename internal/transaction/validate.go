package transaction

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const (
	MaxPlanChanges  = 256
	MaxAdapterLength = 128
	MaxKeyLength     = 512
	MaxValueLength   = 256 * 1024
)

// ValidatePlan rejects ambiguous or unbounded plans before adapter reads or mutations.
func ValidatePlan(plan Plan) error {
	if len(plan.Changes) == 0 {
		return fmt.Errorf("transaction: plan must contain at least one change")
	}
	if len(plan.Changes) > MaxPlanChanges {
		return fmt.Errorf("transaction: plan has %d changes; maximum is %d", len(plan.Changes), MaxPlanChanges)
	}

	seen := make(map[string]struct{}, len(plan.Changes))
	for i, change := range plan.Changes {
		if err := validateChange(change); err != nil {
			return fmt.Errorf("transaction: change %d: %w", i, err)
		}
		if !utf8.ValidString(change.Adapter) || !utf8.ValidString(change.Key) || !utf8.ValidString(change.Value) {
			return fmt.Errorf("transaction: change %d contains invalid UTF-8", i)
		}
		if strings.IndexByte(change.Adapter, 0) >= 0 || strings.IndexByte(change.Key, 0) >= 0 || strings.IndexByte(change.Value, 0) >= 0 {
			return fmt.Errorf("transaction: change %d contains NUL", i)
		}
		if len(change.Adapter) > MaxAdapterLength {
			return fmt.Errorf("transaction: change %d adapter name too long", i)
		}
		if len(change.Key) > MaxKeyLength {
			return fmt.Errorf("transaction: change %d key too long", i)
		}
		if len(change.Value) > MaxValueLength {
			return fmt.Errorf("transaction: change %d value too large", i)
		}

		target := change.Adapter + "\x00" + change.Key
		if _, ok := seen[target]; ok {
			return fmt.Errorf("transaction: duplicate target %q/%q", change.Adapter, change.Key)
		}
		seen[target] = struct{}{}
	}
	return nil
}
