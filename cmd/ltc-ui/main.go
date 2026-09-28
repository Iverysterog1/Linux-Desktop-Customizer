package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/app"
	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/ui"
	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/version"
)

func main() {
	jsonOutput := flag.Bool("json", false, "print the UI foundation model as JSON")
	flag.Parse()

	rt, err := app.New()
	if err != nil {
		fmt.Fprintln(os.Stderr, "ltc-ui:", err)
		os.Exit(1)
	}
	model := ui.FoundationModel(version.Value, rt.Registry)

	if *jsonOutput {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err := enc.Encode(model); err != nil {
			fmt.Fprintln(os.Stderr, "ltc-ui:", err)
			os.Exit(1)
		}
		return
	}

	fmt.Printf("%s %s\n", model.Product, model.Version)
	fmt.Println("UI foundation status:")
	for _, screen := range model.Screens {
		state := "available"
		if !screen.Enabled {
			state = "gated"
		}
		fmt.Printf("  - %s: %s — %s", screen.Title, state, screen.Description)
		if screen.Reason != "" {
			fmt.Printf(" (%s)", screen.Reason)
		}
		fmt.Println()
	}
	for _, warning := range model.Warnings {
		fmt.Printf("warning: %s\n", warning)
	}
}
