package adapter

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/capability"
)

const kdeMissingSentinel = "__LTC_KDE_VALUE_NOT_SET_6f39b26d__"

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

type kdeConfigTarget struct { file, group, key string }

var kdeConfigTargets = map[string]kdeConfigTarget{
	"color-scheme": {file: "kdeglobals", group: "General", key: "ColorScheme"},
}

type KDEConfigAdapter struct {
	readCommand, writeCommand string
	runner KDERunner
	dataDirs []string
}

func NewKDEConfigAdapter() (*KDEConfigAdapter, error) {
	readPath, err := exec.LookPath("kreadconfig6")
	if err != nil { return nil, fmt.Errorf("kde config adapter: kreadconfig6 unavailable: %w", err) }
	writePath, err := exec.LookPath("kwriteconfig6")
	if err != nil { return nil, fmt.Errorf("kde config adapter: kwriteconfig6 unavailable: %w", err) }
	return newKDEConfigAdapterWithDataDirs(readPath, writePath, execKDERunner{}, kdeDataDirs())
}

func newKDEConfigAdapter(readCommand, writeCommand string, runner KDERunner) (*KDEConfigAdapter, error) {
	return newKDEConfigAdapterWithDataDirs(readCommand, writeCommand, runner, kdeDataDirs())
}

func newKDEConfigAdapterWithDataDirs(readCommand, writeCommand string, runner KDERunner, dataDirs []string) (*KDEConfigAdapter, error) {
	if readCommand == "" || writeCommand == "" { return nil, errors.New("kde config adapter: read and write commands are required") }
	if runner == nil { return nil, errors.New("kde config adapter: runner is required") }
	return &KDEConfigAdapter{readCommand: readCommand, writeCommand: writeCommand, runner: runner, dataDirs: append([]string(nil), dataDirs...)}, nil
}

func kdeDataDirs() []string {
	var dirs []string
	if home, err := os.UserHomeDir(); err == nil && home != "" { dirs = append(dirs, filepath.Join(home, ".local", "share")) }
	if configured := os.Getenv("XDG_DATA_DIRS"); configured != "" {
		for _, dir := range filepath.SplitList(configured) { if dir != "" { dirs = append(dirs, dir) } }
	} else { dirs = append(dirs, "/usr/local/share", "/usr/share") }
	return dirs
}

func (a *KDEConfigAdapter) Name() string { return "kde-config" }
func (a *KDEConfigAdapter) Capabilities(context.Context) []capability.Capability {
	return []capability.Capability{{ID:"kde.color-scheme", Supported:true, Reason:"Plasma 6 color scheme can be changed through the reviewed KDE config adapter"}}
}

func (a *KDEConfigAdapter) Read(ctx context.Context, key string) (string, bool, error) {
	target, err := kdeTarget(key); if err != nil { return "", false, err }
	out, err := a.runner.Run(ctx, a.readCommand, "--file", target.file, "--group", target.group, "--key", target.key, "--default", kdeMissingSentinel)
	if err != nil { return "", false, fmt.Errorf("kde config adapter: read %q: %w", key, err) }
	out = strings.TrimSuffix(out, "\r")
	if out == kdeMissingSentinel { return "", false, nil }
	return out, true, nil
}

func (a *KDEConfigAdapter) Set(ctx context.Context, key, value string) error {
	target, err := kdeTarget(key); if err != nil { return err }
	if err := validateKDEValue(value); err != nil { return fmt.Errorf("kde config adapter: set %q: %w", key, err) }
	if key == "color-scheme" && !validKDESchemeName(value) { return fmt.Errorf("kde config adapter: set %q: invalid color scheme name", key) }
	current, exists, err := a.Read(ctx, key); if err != nil { return err }
	if exists && current == value { return nil }
	if key == "color-scheme" && !a.colorSchemeInstalled(value) { return fmt.Errorf("kde config adapter: set %q: color scheme %q is not installed", key, value) }
	if _, err := a.runner.Run(ctx, a.writeCommand, "--file", target.file, "--group", target.group, "--key", target.key, value); err != nil { return fmt.Errorf("kde config adapter: set %q: %w", key, err) }
	return nil
}

func validKDESchemeName(name string) bool {
	if name == "" || name == "." || name == ".." || filepath.Base(name) != name { return false }
	return !strings.ContainsAny(name, `/\\`)
}

func (a *KDEConfigAdapter) colorSchemeInstalled(name string) bool {
	if !validKDESchemeName(name) { return false }
	for _, dataDir := range a.dataDirs {
		if dataDir == "" { continue }
		path := filepath.Join(dataDir, "color-schemes", name+".colors")
		info, err := os.Lstat(path)
		if err == nil && info.Mode().IsRegular() && info.Mode()&os.ModeSymlink == 0 { return true }
	}
	return false
}

func (a *KDEConfigAdapter) Unset(ctx context.Context, key string) error {
	target, err := kdeTarget(key); if err != nil { return err }
	_, exists, err := a.Read(ctx, key); if err != nil { return err }
	if !exists { return nil }
	if _, err := a.runner.Run(ctx, a.writeCommand, "--file", target.file, "--group", target.group, "--key", target.key, "--delete"); err != nil { return fmt.Errorf("kde config adapter: unset %q: %w", key, err) }
	return nil
}

func kdeTarget(key string) (kdeConfigTarget, error) {
	target, ok := kdeConfigTargets[key]; if !ok { return kdeConfigTarget{}, fmt.Errorf("kde config adapter: unsupported key %q", key) }
	return target, nil
}

func validateKDEValue(value string) error {
	if value == "" { return errors.New("value must not be empty") }
	if len(value) > 256 { return errors.New("value exceeds 256 bytes") }
	if !utf8.ValidString(value) { return errors.New("value must be valid UTF-8") }
	if strings.ContainsAny(value, "\x00\r\n") { return errors.New("value contains a forbidden control character") }
	if value == kdeMissingSentinel { return errors.New("value is reserved") }
	return nil
}
