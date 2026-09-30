package profile

import (
	"bytes"
	"strings"
	"testing"
)

func TestParseValidDeclarativeProfile(t *testing.T) {
	t.Parallel()
	data := []byte(`{
		"version": 1,
		"name": "Portable KDE",
		"operations": [
			{"adapter":"kde-config","action":"set","key":"kdeglobals|General|ColorScheme","value":"BreezeDark"},
			{"adapter":"local-file","action":"unset","key":"state/old-theme"}
		]
	}`)
	p, err := Parse(data)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	if p.Name != "Portable KDE" || len(p.Operations) != 2 {
		t.Fatalf("unexpected profile: %+v", p)
	}
}

func TestParseRejectsExecutableOrUnknownFields(t *testing.T) {
	t.Parallel()
	cases := []string{
		`{"version":1,"name":"bad","operations":[{"adapter":"kde-config","action":"set","key":"a","command":"rm -rf ~"}]}`,
		`{"version":1,"name":"bad","script":"echo hi","operations":[{"adapter":"kde-config","action":"set","key":"a"}]}`,
	}
	for _, data := range cases {
		if _, err := Parse([]byte(data)); err == nil {
			t.Fatalf("expected unknown executable field to be rejected: %s", data)
		}
	}
}

func TestParseRejectsDuplicateKeys(t *testing.T) {
	t.Parallel()
	data := []byte(`{"version":1,"name":"one","name":"two","operations":[{"adapter":"kde-config","action":"set","key":"a"}]}`)
	_, err := Parse(data)
	if err == nil || !strings.Contains(err.Error(), "duplicate object key") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestParseRejectsResourceAbuse(t *testing.T) {
	t.Parallel()

	validPrefix := []byte(`{"version":1,"name":"one","operations":[{"adapter":"kde-config","action":"set","key":"a","value":"`)
	validSuffix := []byte(`"}]}`)
	oversized := append(append(validPrefix, bytes.Repeat([]byte("x"), maxProfileBytes)...), validSuffix...)
	if _, err := Parse(oversized); err == nil || !strings.Contains(err.Error(), "document exceeds") {
		t.Fatalf("oversized document error = %v", err)
	}

	deep := `{"version":1,"name":"one","operations":[{"adapter":"kde-config","action":"set","key":"a","extra":` +
		strings.Repeat("[", maxJSONDepth+2) + "0" + strings.Repeat("]", maxJSONDepth+2) + `}]}`
	if _, err := Parse([]byte(deep)); err == nil || !strings.Contains(err.Error(), "nesting exceeds") {
		t.Fatalf("deep nesting error = %v", err)
	}
}

func TestValidateRejectsUnknownAdapterAndAction(t *testing.T) {
	t.Parallel()
	cases := []Profile{
		{
			Version: SchemaVersion,
			Name:    "plugin",
			Operations: []Operation{{
				Adapter: "gnome-extension",
				Action:  ActionSet,
				Key:     "example",
			}},
		},
		{
			Version: SchemaVersion,
			Name:    "command",
			Operations: []Operation{{
				Adapter: "kde-config",
				Action:  Action("exec"),
				Key:     "example",
			}},
		},
	}
	for _, p := range cases {
		if err := Validate(p); err == nil {
			t.Fatalf("expected profile to be rejected: %+v", p)
		}
	}
}

func TestValidateRejectsUnsafeShape(t *testing.T) {
	t.Parallel()
	cases := []Profile{
		{Version: 99, Name: "wrong-version", Operations: []Operation{{Adapter: "kde-config", Action: ActionSet, Key: "a"}}},
		{Version: 1, Name: "", Operations: []Operation{{Adapter: "kde-config", Action: ActionSet, Key: "a"}}},
		{Version: 1, Name: "empty-ops"},
		{Version: 1, Name: "empty-key", Operations: []Operation{{Adapter: "kde-config", Action: ActionSet}}},
		{Version: 1, Name: "nul-key", Operations: []Operation{{Adapter: "kde-config", Action: ActionSet, Key: "a\x00b"}}},
		{Version: 1, Name: "long-key", Operations: []Operation{{Adapter: "kde-config", Action: ActionSet, Key: strings.Repeat("k", maxKeyBytes+1)}}},
		{Version: 1, Name: "long-value", Operations: []Operation{{Adapter: "kde-config", Action: ActionSet, Key: "a", Value: strings.Repeat("v", maxValueBytes+1)}}},
		{Version: 1, Name: "unset-value", Operations: []Operation{{Adapter: "local-file", Action: ActionUnset, Key: "a", Value: "unexpected"}}},
	}
	for _, p := range cases {
		if err := Validate(p); err == nil {
			t.Fatalf("expected invalid profile to be rejected: %+v", p)
		}
	}
}

func TestParseRejectsInvalidUTF8(t *testing.T) {
	t.Parallel()
	data := append([]byte(`{"version":1,"name":"one","operations":[{"adapter":"kde-config","action":"set","key":"a"}]}`), 0xff)
	_, err := Parse(data)
	if err == nil || !strings.Contains(err.Error(), "valid UTF-8") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestParseRejectsTrailingJSON(t *testing.T) {
	t.Parallel()
	data := `{"version":1,"name":"one","operations":[{"adapter":"kde-config","action":"set","key":"a"}]} {"extra":true}`
	_, err := Parse([]byte(data))
	if err == nil || !strings.Contains(err.Error(), "multiple JSON values") {
		t.Fatalf("unexpected error: %v", err)
	}
}
