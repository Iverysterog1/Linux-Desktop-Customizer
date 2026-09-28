package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/app"
	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/version"
)

func main() {
	rt, err := app.New()
	if err != nil {
		fmt.Fprintln(os.Stderr, "ltc-ui:", err)
		os.Exit(1)
	}

	fmt.Printf("Linux Desktop Customizer %s\n", version.Value)
	fmt.Println("UI foundation: graphical shell not implemented yet.")
	fmt.Println("The safety/transaction engine is available through the ltc CLI.")
	for _, a := range rt.Registry.All() {
		fmt.Printf("adapter %s:\n", a.Name())
		for _, cap := range a.Capabilities(context.Background()) {
			fmt.Printf("  - %s: supported=%t", cap.ID, cap.Supported)
			if cap.Reason != "" {
				fmt.Printf(" (%s)", cap.Reason)
			}
			fmt.Println()
		}
	}
}
