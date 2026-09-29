package profile

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

const SchemaVersion = 1

type Action string

const (
	ActionSet   Action = "set"
	ActionUnset Action = "unset"
)

var allowedAdapters = map[string]struct{}{
	"local-file": {},
	"kde-config": {},
}

type Operation struct {
	Adapter string `json:"adapter"`
	Action  Action `json:"action"`
	Key     string `json:"key"`
	Value   string `json:"value,omitempty"`
}

type Profile struct {
	Version    int         `json:"version"`
	Name       string      `json:"name"`
	Operations []Operation `json:"operations"`
}

// Parse accepts only the declarative profile schema. Unknown fields are
// rejected so executable/plugin payloads cannot be smuggled through generic
// JSON and interpreted later by a more permissive layer.
func Parse(data []byte) (Profile, error) {
	var p Profile
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&p); err != nil {
		return Profile{}, fmt.Errorf("profile: decode: %w", err)
	}
	if err := ensureEOF(dec); err != nil {
		return Profile{}, err
	}
	if err := Validate(p); err != nil {
		return Profile{}, err
	}
	return p, nil
}

func Validate(p Profile) error {
	if p.Version != SchemaVersion {
		return fmt.Errorf("profile: unsupported schema version %d", p.Version)
	}
	name := strings.TrimSpace(p.Name)
	if name == "" {
		return fmt.Errorf("profile: name must not be empty")
	}
	if len(name) > 128 {
		return fmt.Errorf("profile: name exceeds 128 characters")
	}
	if len(p.Operations) == 0 {
		return fmt.Errorf("profile: at least one operation is required")
	}
	if len(p.Operations) > 256 {
		return fmt.Errorf("profile: operation count exceeds 256")
	}

	for i, op := range p.Operations {
		if _, ok := allowedAdapters[op.Adapter]; !ok {
			return fmt.Errorf("profile: operation %d uses unsupported adapter %q", i, op.Adapter)
		}
		if strings.TrimSpace(op.Key) == "" {
			return fmt.Errorf("profile: operation %d key must not be empty", i)
		}
		if strings.ContainsRune(op.Key, '\x00') {
			return fmt.Errorf("profile: operation %d key contains NUL", i)
		}
		switch op.Action {
		case ActionSet:
			// Empty values are valid settings; the adapter remains responsible
			// for validating adapter-specific keys and values.
		case ActionUnset:
			if op.Value != "" {
				return fmt.Errorf("profile: operation %d unset must not include a value", i)
			}
		default:
			return fmt.Errorf("profile: operation %d uses unsupported action %q", i, op.Action)
		}
	}
	return nil
}

func ensureEOF(dec *json.Decoder) error {
	var extra any
	if err := dec.Decode(&extra); err == io.EOF {
		return nil
	} else if err != nil {
		return fmt.Errorf("profile: trailing data: %w", err)
	}
	return fmt.Errorf("profile: multiple JSON values are not allowed")
}
