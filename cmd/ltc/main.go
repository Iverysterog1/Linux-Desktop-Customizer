package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/app"
	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/transaction"
	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/version"
)

func main() {
	if err := run(context.Background(), os.Args[1:], os.Stdout, os.Stderr); err != nil {
		fmt.Fprintln(os.Stderr, "ltc:", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, args []string, stdout, stderr io.Writer) error {
	if len(args) == 0 {
		printUsage(stderr)
		return flag.ErrHelp
	}
	switch args[0] {
	case "version":
		fmt.Fprintln(stdout, version.Value)
		return nil
	case "status":
		rt, err := app.New()
		if err != nil {
			return err
		}
		type adapterStatus struct {
			Name         string      `json:"name"`
			Capabilities interface{} `json:"capabilities"`
		}
		status := struct {
			Version    string          `json:"version"`
			StateDir   string          `json:"state_dir"`
			ManagedDir string          `json:"managed_dir"`
			Adapters   []adapterStatus `json:"adapters"`
		}{Version: version.Value, StateDir: rt.StateDir, ManagedDir: rt.ManagedDir}
		for _, a := range rt.Registry.All() {
			status.Adapters = append(status.Adapters, adapterStatus{
				Name: a.Name(), Capabilities: a.Capabilities(ctx),
			})
		}
		return writeJSON(stdout, status)
	case "preview", "apply":
		fs := flag.NewFlagSet(args[0], flag.ContinueOnError)
		fs.SetOutput(stderr)
		profile := fs.String("profile", "", "path to a declarative profile JSON file")
		if err := fs.Parse(args[1:]); err != nil {
			return err
		}
		if *profile == "" {
			return fmt.Errorf("--profile is required")
		}
		plan, err := loadPlan(*profile)
		if err != nil {
			return err
		}
		rt, err := app.New()
		if err != nil {
			return err
		}
		if args[0] == "preview" {
			preview, err := rt.Engine.Preview(ctx, plan)
			if err != nil {
				return err
			}
			return writeJSON(stdout, preview)
		}
		tx, err := rt.Engine.Apply(ctx, plan)
		if err != nil {
			if tx.ID != "" {
				_ = writeJSON(stdout, tx)
			}
			if errors.Is(err, transaction.ErrUnsupported) {
				return fmt.Errorf("%w; run preview for details", err)
			}
			return err
		}
		return writeJSON(stdout, tx)
	case "rollback", "unapply":
		fs := flag.NewFlagSet(args[0], flag.ContinueOnError)
		fs.SetOutput(stderr)
		id := fs.String("transaction", "", "transaction id to restore")
		if err := fs.Parse(args[1:]); err != nil {
			return err
		}
		if *id == "" {
			return fmt.Errorf("--transaction is required")
		}
		rt, err := app.New()
		if err != nil {
			return err
		}
		tx, err := rt.Engine.Rollback(ctx, *id)
		if err != nil {
			return err
		}
		return writeJSON(stdout, tx)
	default:
		printUsage(stderr)
		return fmt.Errorf("unknown command %q", args[0])
	}
}

func loadPlan(path string) (transaction.Plan, error) {
	f, err := os.Open(path)
	if err != nil {
		return transaction.Plan{}, fmt.Errorf("open profile: %w", err)
	}
	defer f.Close()

	dec := json.NewDecoder(io.LimitReader(f, 1<<20))
	dec.DisallowUnknownFields()
	var plan transaction.Plan
	if err := dec.Decode(&plan); err != nil {
		return transaction.Plan{}, fmt.Errorf("decode profile: %w", err)
	}
	var extra any
	if err := dec.Decode(&extra); err != io.EOF {
		if err == nil {
			return transaction.Plan{}, fmt.Errorf("decode profile: trailing JSON content")
		}
		return transaction.Plan{}, fmt.Errorf("decode profile: %w", err)
	}
	return plan, nil
}

func writeJSON(w io.Writer, value any) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(value)
}

func printUsage(w io.Writer) {
	fmt.Fprintln(w, "Linux Desktop Customizer foundation")
	fmt.Fprintln(w, "usage:")
	fmt.Fprintln(w, "  ltc version")
	fmt.Fprintln(w, "  ltc status")
	fmt.Fprintln(w, "  ltc preview --profile profile.json")
	fmt.Fprintln(w, "  ltc apply --profile profile.json")
	fmt.Fprintln(w, "  ltc rollback --transaction ID")
	fmt.Fprintln(w, "  ltc unapply --transaction ID")
}
