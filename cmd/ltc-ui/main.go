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

func localeFromEnvironment() string {
	for _, key := range []string{"LC_ALL", "LC_MESSAGES", "LANG"} {
		if value := os.Getenv(key); value != "" {
			return value
		}
	}
	return ui.LocaleEnglish
}

func main() {
	jsonOutput := flag.Bool("json", false, "print the UI foundation model as JSON")
	locale := flag.String("locale", "", "UI locale (en or pt-PT); defaults to LC_ALL, LC_MESSAGES, then LANG")
	flag.Parse()

	rt, err := app.New()
	if err != nil {
		fmt.Fprintln(os.Stderr, "ltc-ui:", err)
		os.Exit(1)
	}
	selectedLocale := *locale
	if selectedLocale == "" {
		selectedLocale = localeFromEnvironment()
	}
	model := ui.FoundationModelForLocale(version.Value, rt.Registry, selectedLocale)

	if *jsonOutput {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err := enc.Encode(model); err != nil {
			fmt.Fprintln(os.Stderr, "ltc-ui:", err)
			os.Exit(1)
		}
		return
	}

	labels := ui.LabelsForLocale(selectedLocale)
	fmt.Printf("%s %s\n", model.Product, model.Version)
	fmt.Println(labels.Status)
	for _, screen := range model.Screens {
		state := labels.Available
		if !screen.Enabled {
			state = labels.Gated
		}
		fmt.Printf("  - %s: %s — %s", screen.Title, state, screen.Description)
		if screen.Reason != "" {
			fmt.Printf(" (%s)", screen.Reason)
		}
		fmt.Println()
	}
	for _, warning := range model.Warnings {
		fmt.Printf("%s: %s\n", labels.Warning, warning)
	}
}
