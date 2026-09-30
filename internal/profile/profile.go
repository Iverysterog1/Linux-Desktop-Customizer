package profile

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"unicode/utf8"
)

const (
	SchemaVersion   = 1
	maxProfileBytes = 256 * 1024
	maxNameBytes    = 128
	maxKeyBytes     = 1024
	maxValueBytes   = 64 * 1024
	maxJSONDepth    = 16
)

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
	if len(data) == 0 {
		return Profile{}, fmt.Errorf("profile: document must not be empty")
	}
	if len(data) > maxProfileBytes {
		return Profile{}, fmt.Errorf("profile: document exceeds %d bytes", maxProfileBytes)
	}
	if !utf8.Valid(data) {
		return Profile{}, fmt.Errorf("profile: document must be valid UTF-8")
	}
	if err := validateJSONStructure(data); err != nil {
		return Profile{}, err
	}

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
	if len(name) > maxNameBytes {
		return fmt.Errorf("profile: name exceeds %d bytes", maxNameBytes)
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
		if len(op.Key) > maxKeyBytes {
			return fmt.Errorf("profile: operation %d key exceeds %d bytes", i, maxKeyBytes)
		}
		if strings.ContainsRune(op.Key, '\x00') {
			return fmt.Errorf("profile: operation %d key contains NUL", i)
		}
		if len(op.Value) > maxValueBytes {
			return fmt.Errorf("profile: operation %d value exceeds %d bytes", i, maxValueBytes)
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

func validateJSONStructure(data []byte) error {
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	if err := walkJSONValue(dec, 0); err != nil {
		return fmt.Errorf("profile: JSON structure: %w", err)
	}
	if _, err := dec.Token(); err == io.EOF {
		return nil
	} else if err != nil {
		return fmt.Errorf("profile: trailing JSON: %w", err)
	}
	return fmt.Errorf("profile: multiple JSON values are not allowed")
}

func walkJSONValue(dec *json.Decoder, depth int) error {
	if depth > maxJSONDepth {
		return fmt.Errorf("nesting exceeds %d levels", maxJSONDepth)
	}

	tok, err := dec.Token()
	if err != nil {
		return err
	}
	delim, ok := tok.(json.Delim)
	if !ok {
		return nil
	}

	switch delim {
	case '{':
		seen := make(map[string]struct{})
		for dec.More() {
			keyToken, err := dec.Token()
			if err != nil {
				return err
			}
			key, ok := keyToken.(string)
			if !ok {
				return fmt.Errorf("object key is not a string")
			}
			if _, exists := seen[key]; exists {
				return fmt.Errorf("duplicate object key %q", key)
			}
			seen[key] = struct{}{}
			if err := walkJSONValue(dec, depth+1); err != nil {
				return err
			}
		}
		end, err := dec.Token()
		if err != nil {
			return err
		}
		if end != json.Delim('}') {
			return fmt.Errorf("object not terminated")
		}
	case '[':
		for dec.More() {
			if err := walkJSONValue(dec, depth+1); err != nil {
				return err
			}
		}
		end, err := dec.Token()
		if err != nil {
			return err
		}
		if end != json.Delim(']') {
			return fmt.Errorf("array not terminated")
		}
	default:
		return fmt.Errorf("unexpected delimiter %q", delim)
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
