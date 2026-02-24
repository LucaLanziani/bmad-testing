package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/LucaLanziani/bmad-testing/conciglia/internal/config"
)

func main() {
	fmt.Println("=== Conciglia Color Configuration Demo ===")

	// Load configuration
	cfg, err := config.Load("")
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	// Display configuration
	cfgJSON, _ := json.MarshalIndent(cfg, "", "  ")
	fmt.Printf("Current Configuration:\n%s\n\n", cfgJSON)

	// Display environment variables
	fmt.Printf("Environment Variables:\n")
	fmt.Printf("  NO_COLOR: %q\n", os.Getenv("NO_COLOR"))
	fmt.Printf("  CONCIGLIA_NO_COLOR: %q\n", os.Getenv("CONCIGLIA_NO_COLOR"))
	fmt.Printf("\n")

	// Display color status
	fmt.Printf("Colors Enabled: %v\n\n", cfg.ColorsEnabled())

	// Demonstrate color output
	if cfg.ColorsEnabled() {
		fmt.Println("Color output example:")
		fmt.Println("  \033[32m✓ Success message (green)\033[0m")
		fmt.Println("  \033[33m⚠ Warning message (yellow)\033[0m")
		fmt.Println("  \033[31m✗ Error message (red)\033[0m")
		fmt.Println("  \033[36m→ Info message (cyan)\033[0m")
	} else {
		fmt.Println("Plain output example:")
		fmt.Println("  ✓ Success message")
		fmt.Println("  ⚠ Warning message")
		fmt.Println("  ✗ Error message")
		fmt.Println("  → Info message")
	}

	fmt.Println("\n=== Try these commands: ===")
	fmt.Println("NO_COLOR=1 go run examples/demo.go")
	fmt.Println("CONCIGLIA_NO_COLOR=1 go run examples/demo.go")
}
