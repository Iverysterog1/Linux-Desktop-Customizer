package profile

import (
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
		{Version: 1, Name: "unset-value", Operations: []Operation{{Adapter: "local-file", Action: ActionUnset, Key: "a", Value: "unexpected"}}},
	}
	for _, p := range cases {
		if err := Validate(p); err == nil {
			t.Fatalf("expected invalid profile to be rejected: %+v", p)
		}
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
