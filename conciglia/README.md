# Conciglia - Story 6.4: Color Configuration Implementation

This directory contains the implementation of Story 6.4 (FR26): Color Configuration support for the Conciglia shell.

## Implementation Summary

### Features Implemented

✅ **Config File Support** - `{"output": {"colors": false}}` disables colors  
✅ **CONCIGLIA_NO_COLOR Support** - `CONCIGLIA_NO_COLOR=1` disables colors  
✅ **NO_COLOR Support** - Respects the `NO_COLOR` environment variable (de facto standard)  
✅ **TTY Detection** - Default: true when connected to a TTY, false otherwise  

### Configuration Priority

The configuration system follows this priority order (highest to lowest):

1. **NO_COLOR** environment variable (any value disables colors)
2. **CONCIGLIA_NO_COLOR=1** environment variable (only "1" disables colors)
3. Configuration file setting (`output.colors`)
4. TTY detection (default)

### Package Structure

```
conciglia/
├── internal/
│   └── config/
│       ├── config.go       # Configuration loading and management
│       └── config_test.go  # Comprehensive test suite
├── go.mod
└── go.sum
```

## Usage Examples

### Loading Configuration

```go
import "github.com/LucaLanziani/bmad-testing/conciglia/internal/config"

// Load from file (or use defaults if file doesn't exist)
cfg, err := config.Load("/path/to/config.json")
if err != nil {
    log.Fatal(err)
}

// Check if colors are enabled
if cfg.ColorsEnabled() {
    fmt.Println("\033[32mGreen text!\033[0m")
} else {
    fmt.Println("Plain text")
}
```

### Configuration File Format

Create a `config.json` file:

```json
{
  "output": {
    "colors": false
  }
}
```

### Environment Variables

```bash
# Disable colors using NO_COLOR (standard)
NO_COLOR=1 conciglia

# Disable colors using CONCIGLIA_NO_COLOR
CONCIGLIA_NO_COLOR=1 conciglia
```

## Demo

Run the interactive demo to see color configuration in action:

```bash
# Default behavior (colors enabled in TTY)
go run examples/demo.go

# Disable colors with NO_COLOR
NO_COLOR=1 go run examples/demo.go

# Disable colors with CONCIGLIA_NO_COLOR
CONCIGLIA_NO_COLOR=1 go run examples/demo.go
```

## Testing

Run the test suite:

```bash
cd conciglia
go test ./internal/config -v
```

### Test Coverage

- ✅ Default configuration with TTY detection
- ✅ Loading configuration from JSON file (colors enabled)
- ✅ Loading configuration from JSON file (colors disabled)
- ✅ Non-existent config file (uses defaults)
- ✅ CONCIGLIA_NO_COLOR=1 override
- ✅ CONCIGLIA_NO_COLOR with other values (should not disable)
- ✅ NO_COLOR environment variable
- ✅ Priority ordering (NO_COLOR > CONCIGLIA_NO_COLOR)

## Acceptance Criteria Status

- ✅ Config `{"output": {"colors": false}}` disables colors
- ✅ `CONCIGLIA_NO_COLOR=1` disables colors
- ✅ `NO_COLOR` env var is respected (standard)

## Technical Implementation Details

### TTY Detection

The implementation uses `golang.org/x/term` package to detect if stdout is connected to a terminal:

```go
func isTerminal() bool {
    return term.IsTerminal(int(os.Stdout.Fd()))
}
```

### Environment Variable Processing

1. **NO_COLOR**: Any non-empty value disables colors (per https://no-color.org/ standard)
2. **CONCIGLIA_NO_COLOR**: Only the value "1" disables colors (app-specific convention)

### Configuration Loading

The `Load()` function:
1. Creates default configuration (with TTY detection)
2. Loads from JSON file if it exists
3. Applies environment variable overrides
4. Returns the final configuration

## Dependencies

- `golang.org/x/term` - Terminal detection (cross-platform)

## Future Enhancements

Potential improvements for future stories:

- Color scheme customization (different color palettes)
- Per-output-type color settings (errors, warnings, success)
- Force-color mode for non-TTY scenarios (CI/CD)
