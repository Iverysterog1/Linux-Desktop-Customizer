package app

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/adapter"
	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/platform"
	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/transaction"
)

type Runtime struct {
	Registry   *adapter.Registry
	Engine     *transaction.Engine
	ManagedDir string
	StateDir   string
}

func New() (*Runtime, error) {
	managedDir, err := managedDir()
	if err != nil {
		return nil, err
	}
	stateDir, err := stateDir()
	if err != nil {
		return nil, err
	}

	fileAdapter, err := adapter.NewFileAdapter(managedDir)
	if err != nil {
		return nil, err
	}
	adapters := []adapter.Adapter{fileAdapter}

	kdeStatus := platform.ProbeKDE(os.Getenv, exec.LookPath)
	if kdeRuntimeReady(kdeStatus) {
		kdeAdapter, err := adapter.NewKDEConfigAdapter()
		if err != nil {
			return nil, fmt.Errorf("app: initialize KDE adapter after successful readiness detection: %w", err)
		}
		adapters = append(adapters, kdeAdapter)
	}

	registry, err := adapter.NewRegistry(adapters...)
	if err != nil {
		return nil, err
	}
	store, err := transaction.NewStore(filepath.Join(stateDir, "transactions"))
	if err != nil {
		return nil, err
	}
	engine, err := transaction.NewEngine(registry, store)
	if err != nil {
		return nil, err
	}
	return &Runtime{
		Registry:   registry,
		Engine:     engine,
		ManagedDir: managedDir,
		StateDir:   stateDir,
	}, nil
}

func kdeRuntimeReady(status platform.KDEStatus) bool {
	if !status.Detected {
		return false
	}
	readAvailable := false
	writeAvailable := false
	for _, command := range status.Commands {
		switch command.Name {
		case "kreadconfig6":
			readAvailable = command.Available
		case "kwriteconfig6":
			writeAvailable = command.Available
		}
	}
	return readAvailable && writeAvailable
}

func managedDir() (string, error) {
	if v := os.Getenv("LTC_MANAGED_DIR"); v != "" {
		return v, nil
	}
	if v := os.Getenv("XDG_CONFIG_HOME"); v != "" {
		return filepath.Join(v, "linux-desktop-customizer", "managed"), nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("app: determine home directory: %w", err)
	}
	return filepath.Join(home, ".config", "linux-desktop-customizer", "managed"), nil
}

func stateDir() (string, error) {
	if v := os.Getenv("LTC_STATE_DIR"); v != "" {
		return v, nil
	}
	if v := os.Getenv("XDG_STATE_HOME"); v != "" {
		return filepath.Join(v, "linux-desktop-customizer"), nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("app: determine home directory: %w", err)
	}
	return filepath.Join(home, ".local", "state", "linux-desktop-customizer"), nil
}
