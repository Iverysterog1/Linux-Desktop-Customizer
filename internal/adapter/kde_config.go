package adapter

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"unicode/utf8"

	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/capability"
)

const kdeMissingSentinel = "__LTC_KDE_VALUE_NOT_SET_6f39b26d__"

// KDERunner is the narrow process boundary used by KDEConfigAdapter. Keeping
// command execution behind this interface makes the adapter deterministic in
// tests and prevents profiles from supplying executable names or arguments.
type KDERunner interface {
	Run(context.Context, string, ...string) (string, error)
}

type execKDERunner struct{}

func (execKDERunner) Run(ctx context.Context, name string, args ...string) (string, error) {
	out, err := exec.CommandContext(ctx, name, args...).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s: %w: %s", name, err, strings.TrimSpace(string(out)))
	}
	return strings.TrimSuffix(string(out), "\n"), nil
}

type kdeConfigTarget struct {
	file  string
	group string
	key   string
}

// kdeConfigTargets is intentionally an allowlist. Profile keys never become
// filenames, groups, command names, flags, or arbitrary command arguments.
var kdeConfigTargets = map[string]kdeConfigTarget{
	"color-scheme": {file: "kdeglobals", group: "General", key: "ColorScheme"},
}

// KDEConfigAdapter owns a small, reviewed Plasma 6 configuration surface.
// The first slice supports only the global color-scheme key; further settings
// must be added explicitly to kdeConfigTargets after review and tests.
type KDEConfigAdapter struct {
	readCommand  string
	writeCommand string
	runner       KDERunner
}

func NewKDEConfigAdapter() (*KDEConfigAdapter, error) {
	readPath, err := exec.LookPath("kreadconfig6")
	if err != nil {
		return nil, fmt.Errorf("kde config adapter: kreadconfig6 unavailable: %w", err)
	}
	writePath, err := exec.LookPath("kwriteconfig6")
	if err != nil {
		return nil, fmt.Errorf("kde config adapter: kwriteconfig6 unavailable: %w", err)
	}
	return newKDEConfigAdapter(readPath, writePath, execKDERunner{})
}

func newKDEConfigAdapter(readCommand, writeCommand string, runner KDERunner) (*KDEConfigAdapter, error) {
	if readCommand == "" || writeCommand == "" {
		return nil, errors.New("kde config adapter: read and write commands are required")
	}
	if runner == nil {
		return nil, errors.New("kde config adapter: runner is required")
	}
	return &KDEConfigAdapter{readCommand: readCommand, writeCommand: writeCommand, runner: runner}, nil
}

func (a *KDEConfigAdapter) Name() string { return "kde-config" }

func (a *KDEConfigAdapter) Capabilities(context.Context) []capability.Capability {
	return []capability.Capability{{
		ID:        "kde.color-scheme",
		Supported: true,
		Reason:    "Plasma 6 color scheme can be changed through the reviewed KDE config adapter",
	}}
}

func (a *KDEConfigAdapter) Read(ctx context.Context, key string) (string, bool, error) {
	target, err := kdeTarget(key)
	if err != nil {
		return "", false, err
	}
	out, err := a.runner.Run(ctx, a.readCommand,
		"--file", target.file,
		"--group", target.group,
		"--key", target.key,
		"--default", kdeMissingSentinel,
	)
	if err != nil {
		return "", false, fmt.Errorf("kde config adapter: read %q: %w", key, err)
	}
	out = strings.TrimSuffix(out, "\r")
	if out == kdeMissingSentinel {
		return "", false, nil
	}
	return out, true, nil
}

func (a *KDEConfigAdapter) Set(ctx context.Context, key, value string) error {
	target, err := kdeTarget(key)
	if err != nil {
		return err
	}
	if err := validateKDEValue(value); err != nil {
		return fmt.Errorf("kde config adapter: set %q: %w", key, err)
	}

	current, exists, err := a.Read(ctx, key)
	if err != nil {
		return err
	}
	if exists && current == value {
		return nil
	}
	if _, err := a.runner.Run(ctx, a.writeCommand,
		"--file", target.file,
		"--group", target.group,
		"--key", target.key,
		value,
	); err != nil {
		return fmt.Errorf("kde config adapter: set %q: %w", key, err)
	}
	return nil
}

func (a *KDEConfigAdapter) Unset(ctx context.Context, key string) error {
	target, err := kdeTarget(key)
	if err != nil {
		return err
	}
	_, exists, err := a.Read(ctx, key)
	if err != nil {
		return err
	}
	if !exists {
		return nil
	}
	if _, err := a.runner.Run(ctx, a.writeCommand,
		"--file", target.file,
		"--group", target.group,
		"--key", target.key,
		"--delete",
	); err != nil {
		return fmt.Errorf("kde config adapter: unset %q: %w", key, err)
	}
	return nil
}

func kdeTarget(key string) (kdeConfigTarget, error) {
	target, ok := kdeConfigTargets[key]
	if !ok {
		return kdeConfigTarget{}, fmt.Errorf("kde config adapter: unsupported key %q", key)
	}
	return target, nil
}

func validateKDEValue(value string) error {
	if value == "" {
		return errors.New("value must not be empty")
	}
	if len(value) > 256 {
		return errors.New("value exceeds 256 bytes")
	}
	if !utf8.ValidString(value) {
		return errors.New("value must be valid UTF-8")
	}
	if strings.ContainsAny(value, "\x00\r\n") {
		return errors.New("value contains a forbidden control character")
	}
	if value == kdeMissingSentinel {
		return errors.New("value is reserved")
	}
	return nil
}
