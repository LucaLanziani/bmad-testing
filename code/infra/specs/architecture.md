---
stepsCompleted: [1, 2, 3, 4, 5, 6, 7, 8]
status: complete
completedAt: '2025-12-04'
inputDocuments:
  - specs/prd.md
  - specs/analysis/product-brief-Project101-2025-12-04.md
workflowType: 'architecture'
lastStep: 0
project_name: 'Conciglia'
user_name: 'Nearformer'
date: '2025-12-04'
---

# Architecture Decision Document

_This document builds collaboratively through step-by-step discovery. Sections are appended as we work through each architectural decision together._

---

## Project Context Analysis

### Requirements Overview

**Functional Requirements (26 FRs in 7 capability areas):**

| Area | FRs | Architectural Implication |
|------|-----|---------------------------|
| Command Execution | FR1-4 | POSIX command executor, process management |
| JSON Pipeline Ops | FR5-10 | Stream-based JSON parser, operator pipeline |
| JSON Literals | FR11-13 | Lexer/parser for JSON detection |
| Interactive Shell | FR14-17 | Readline library, history persistence |
| Output Formatting | FR18-20 | Formatter module, TTY detection |
| Error Handling | FR21-23 | Error type system, exit code mapping |
| Configuration | FR24-26 | Config loader, XDG paths |

**Non-Functional Requirements:**

| Category | Key Constraint | Impact |
|----------|----------------|--------|
| Performance | < 50ms startup | No reflection, minimal init |
| Performance | Stream-based JSON | Can't load entire JSON into memory |
| Compatibility | 95% POSIX | Need robust command passthrough |
| Portability | Single static binary | No cgo, pure Go |
| Portability | Linux x86_64 + ARM64 | Cross-compilation |

### Scale & Complexity

- **Primary domain:** CLI/Systems programming
- **Complexity level:** Medium — Well-defined scope, parsing/execution non-trivial
- **Estimated architectural components:** 6-8 major modules

### Technical Constraints & Dependencies

- **Language:** Go (specified in PRD)
- **Dependencies:** Go stdlib only — no external dependencies for v1.0
- **Binary:** Single static binary, no cgo
- **Platforms:** Linux x86_64, Linux ARM64

### Cross-Cutting Concerns Identified

1. **JSON Detection** — Determining when output is JSON vs text affects multiple components
2. **Error Handling** — Consistent error messages and exit codes throughout
3. **Output Formatting** — TTY detection and format selection spans pipeline
4. **Stream Processing** — Memory-efficient handling required in parser and operators

---

## Starter Template Evaluation

### Primary Technology Domain

**CLI Tool / Shell** — written in Go (specified in PRD)

Go CLI projects use established **project structure conventions** rather than opinionated starter templates.

### Project Structure

```
conciglia/
├── cmd/
│   └── conciglia/        # Main entry point
│       └── main.go
├── internal/             # Private packages (not importable)
│   ├── lexer/            # Tokenization
│   ├── parser/           # AST generation
│   ├── executor/         # Command execution
│   ├── pipeline/         # JSON pipeline operators
│   ├── shell/            # Interactive shell (REPL)
│   └── config/           # Configuration loading
├── pkg/                  # Public packages (if needed)
├── go.mod
├── go.sum
├── Makefile
├── .goreleaser.yaml      # Release automation
└── .golangci.yaml        # Linter config
```

### Tooling Decisions

| Tool | Purpose | Rationale |
|------|---------|-----------|
| **Go modules** | Dependency management | Standard since Go 1.11 |
| **golangci-lint** | Code quality | Aggregates 50+ linters |
| **goreleaser** | Cross-platform builds & releases | Single binary for Linux x86_64 + ARM64 |
| **Makefile** | Build automation | Standard for Go projects |
| **Go testing** | Unit/integration tests | Built-in, no external framework |

### Initialization Command

```bash
mkdir conciglia && cd conciglia
go mod init github.com/youruser/conciglia
mkdir -p cmd/conciglia internal/{lexer,parser,executor,pipeline,shell,config}
touch cmd/conciglia/main.go
```

**Note:** Project initialization should be the first implementation story.

---

## Core Architectural Decisions

### Decision Priority Analysis

**Critical Decisions (Block Implementation):**
- Parser architecture (goyacc)
- Execution model (hybrid)
- Pipeline data flow (Value interface)

**Important Decisions (Shape Architecture):**
- JSON streaming strategy
- Error handling approach
- Configuration format

**Deferred Decisions (Post-MVP):**
- Plugin system architecture
- Advanced scripting (functions, control flow)

### Parser Architecture

**Decision:** goyacc parser generator

**Rationale:** Formal grammar definition with YACC-style syntax. Generates pure Go code that integrates cleanly with the codebase. Grammar file serves as documentation.

**Implementation:** Grammar file in `internal/parser/grammar.y`, generated parser in same package.

### Execution Model

**Decision:** Hybrid passthrough + intercept

**Rationale:** POSIX commands execute via fork/exec to the OS, preserving full compatibility. JSON operators (`where`, `select`, `map`) execute in-process for type safety and streaming.

**Flow:**
```
Input → Lexer → Parser → AST
                          ↓
              ┌───────────┴───────────┐
              ↓                       ↓
        POSIX Command            JSON Operator
        (fork/exec)              (in-process)
              ↓                       ↓
           stdout                   Value
              └───────────┬───────────┘
                          ↓
                    Next Stage
```

### Pipeline Data Flow

**Decision:** Value interface with streaming support

**Rationale:** A unified `Value` interface can represent text, JSON object, JSON array (streamable), or error. Enables memory-efficient processing of large datasets.

```go
type Value interface {
    Type() ValueType
    AsJSON() (any, error)
    AsText() (string, error)
    Stream() <-chan Value  // For arrays
}
```

### JSON Streaming Strategy

**Decision:** Stream top-level arrays only

**Rationale:** Most JSON pipeline use cases filter arrays (e.g., `curl api | where .status == "active"`). Array elements stream one-by-one. Objects and nested structures parse fully. Covers 90% of use cases with manageable complexity.

### Interactive Shell (Readline)

**Decision:** Custom minimal readline using `golang.org/x/term`

**Rationale:** Maintains stdlib-only constraint. Implement basic line editing, history navigation, and Ctrl+R search. `golang.org/x/term` is quasi-stdlib (maintained by Go team).

**Scope for v1.0:**
- Arrow key navigation (left/right within line, up/down for history)
- Ctrl+A/E (beginning/end of line)
- Ctrl+R (reverse history search)
- History persistence to file

### Error Handling

**Decision:** Errors as Value type

**Rationale:** Errors flow through pipeline as a special Value type. Enables graceful degradation — one bad record doesn't stop the entire pipeline. Users can filter or inspect errors explicitly.

```
curl api | where .x > 5
# If item 3 has no .x field:
# → Warning: .x not found on item 3 (item skipped)
# → Continues processing remaining items
```

### Configuration

**Decision:** JSON format at `~/.config/conciglia/config.json`

**Rationale:** Stdlib support via `encoding/json`. Consistent with Conciglia's JSON-native identity. Users already comfortable with JSON.

```json
{
  "output": {
    "defaultFormat": "pretty",
    "colors": true
  },
  "history": {
    "file": "~/.local/share/conciglia/history",
    "maxEntries": 10000
  }
}
```

### Decision Impact Analysis

**Implementation Sequence:**
1. Lexer/Parser (goyacc grammar)
2. Value interface and types
3. POSIX command executor
4. JSON operators (where, select, map)
5. Pipeline orchestration
6. Interactive shell (readline)
7. Configuration loading

**Cross-Component Dependencies:**
- Parser depends on lexer token types
- Executor depends on Value interface
- Pipeline depends on both executor and Value
- Shell depends on pipeline and config

---

## Implementation Patterns & Consistency Rules

### Pattern Categories Defined

**Critical Conflict Points Identified:** 5 areas where AI agents could make different choices

### Go Naming Conventions

**Package Names:**
- Lowercase, no underscores: `lexer`, `parser`, `pipeline`, `executor`
- Descriptive but concise

**Functions:**
- Public: PascalCase (`ParseJSON`, `NewValue`, `Execute`)
- Private: camelCase (`parseToken`, `nextChar`, `handleError`)

**Files:**
- snake_case: `value_type.go`, `json_parser.go`, `pipeline_stage.go`
- Tests co-located: `value_test.go` next to `value.go`

**Variables:**
- camelCase: `currentToken`, `inputReader`, `outputValue`
- Short names OK in small scopes: `i`, `v`, `err`

### Error Message Format

**User-Facing Errors MUST Include:**
1. Error type/category
2. Location (line/column when applicable)
3. Context (what was being parsed/executed)
4. Suggestion (when determinable)

**Format:**
```
error: <category> at <location>: <detail> — <suggestion>
```

**Examples:**

```
error: invalid JSON at line 3, column 15: unexpected '}' — did you mean ']'?
error: unknown operator 'filter' — did you mean 'where'?
error: type mismatch in 'where': expected number, got string at .price
```

**Internal Errors (Logged):**
- Simple format: `level: message`
- Include stack context in debug mode

### JSON Field Access Syntax

**jq-Compatible Syntax:**

| Syntax | Meaning | Example |
|--------|---------|---------|
| `.field` | Access field | `.name` → `"Alice"` |
| `.field.nested` | Nested access | `.user.email` |
| `.[n]` | Array index | `.[0]` → first element |
| `.[]` | Iterate array | `.users[]` → each user |
| `.["key"]` | Special key | `.["my-key"]` |

**Chaining:**
```
.users[].orders[].total
```

### Null/Empty Handling (Permissive Mode)

| Scenario | Behavior | Output |
|----------|----------|--------|
| `where` on `[]` | Return empty | `[]` |
| `.x` on missing key | Return null + warn | `null` |
| `where .x > 5` on string | Skip item + warn | (item omitted) |
| `select .x` on `null` | Return null | `null` |
| Operator on error Value | Propagate error | error Value |

**Warning Format:**
```
warning: .price not found on item 3 (skipped)
warning: type mismatch at item 7: .count is string, expected number (skipped)
```

### Output Formatting Rules

| Context | Format | Rationale |
|---------|--------|-----------|
| Interactive (TTY) | Pretty JSON (indented, colored) | Human readability |
| Piped to command | Compact JSON (NDJSON) | Machine processing |
| `--pretty` flag | Force pretty | User override |
| `--json` flag | Force compact | User override |
| `--text` flag | Plain text | Legacy compat |
| Text passthrough | Unchanged | POSIX commands |

**TTY Detection:**
```go
if term.IsTerminal(int(os.Stdout.Fd())) {
    // Pretty output
} else {
    // Compact output
}
```

### Enforcement Guidelines

**All AI Agents MUST:**
1. Follow Go naming conventions exactly as specified
2. Use the error message format for all user-facing errors
3. Implement jq-compatible field access syntax
4. Handle null/empty permissively with warnings
5. Respect TTY detection for output formatting

**Pattern Verification:**
- `golangci-lint` enforces naming conventions
- Integration tests verify error message format
- Test cases cover null/empty edge cases

### Anti-Patterns to Avoid

**❌ Don't:**
- Use underscores in package names (`json_ops`)
- Return generic errors (`error: something went wrong`)
- Crash on missing JSON fields
- Assume output format without checking TTY
- Invent new field access syntax

**✅ Do:**
- Follow standard Go idioms
- Provide helpful, actionable error messages
- Fail gracefully with warnings
- Auto-detect context for formatting
- Stay jq-compatible for familiar syntax

---

## Project Structure & Boundaries

### Complete Project Directory Structure

```
conciglia/
├── cmd/
│   └── conciglia/
│       └── main.go                 # Entry point, flag parsing
├── internal/
│   ├── lexer/
│   │   ├── lexer.go                # Tokenizer
│   │   ├── token.go                # Token types
│   │   └── lexer_test.go
│   ├── parser/
│   │   ├── grammar.y               # goyacc grammar file
│   │   ├── parser.go               # Generated + helpers
│   │   ├── ast.go                  # AST node types
│   │   └── parser_test.go
│   ├── value/
│   │   ├── value.go                # Value interface
│   │   ├── json.go                 # JSON value types
│   │   ├── text.go                 # Text value type
│   │   ├── error.go                # Error value type
│   │   ├── stream.go               # Streaming array iterator
│   │   └── value_test.go
│   ├── executor/
│   │   ├── executor.go             # Command execution
│   │   ├── posix.go                # POSIX fork/exec
│   │   ├── builtin.go              # Built-in operators
│   │   └── executor_test.go
│   ├── pipeline/
│   │   ├── pipeline.go             # Pipeline orchestration
│   │   ├── stage.go                # Pipeline stage interface
│   │   └── pipeline_test.go
│   ├── operators/
│   │   ├── where.go                # where operator
│   │   ├── select.go               # select operator
│   │   ├── map.go                  # map operator
│   │   ├── operators_test.go
│   │   └── field_access.go         # jq-style field access
│   ├── shell/
│   │   ├── shell.go                # Interactive REPL
│   │   ├── readline.go             # Custom readline impl
│   │   ├── history.go              # History management
│   │   └── shell_test.go
│   ├── output/
│   │   ├── formatter.go            # Output formatting
│   │   ├── pretty.go               # Pretty JSON printer
│   │   ├── color.go                # Terminal colors
│   │   └── formatter_test.go
│   └── config/
│       ├── config.go               # Config loading
│       ├── defaults.go             # Default values
│       └── config_test.go
├── go.mod
├── go.sum
├── Makefile
├── .goreleaser.yaml
├── .golangci.yaml
├── .github/
│   └── workflows/
│       ├── ci.yml                  # Lint + test
│       └── release.yml             # goreleaser
├── README.md
├── LICENSE
└── docs/
    ├── grammar.md                  # Grammar documentation
    └── operators.md                # Operator reference
```

### Architectural Boundaries

| Boundary | Interface | Direction |
|----------|-----------|-----------|
| **Lexer → Parser** | `[]Token` | Lexer produces tokens, parser consumes |
| **Parser → Executor** | `AST` | Parser produces AST, executor walks it |
| **Executor → Pipeline** | `Value` | Commands produce Values |
| **Pipeline → Output** | `Value` | Pipeline produces final Value |
| **Config → All** | `Config struct` | Read-only config access |

### Data Flow

```
User Input
    ↓
┌─────────┐
│  Lexer  │ → []Token
└────┬────┘
     ↓
┌─────────┐
│ Parser  │ → AST
└────┬────┘
     ↓
┌──────────┐
│ Executor │ → Walks AST, dispatches to:
└────┬─────┘
     ├─────────────────┐
     ↓                 ↓
┌─────────┐     ┌──────────┐
│  POSIX  │     │ Operators │
│  Exec   │     │ (where,  │
│         │     │  select) │
└────┬────┘     └────┬─────┘
     ↓                ↓
     └────────┬───────┘
              ↓
        ┌──────────┐
        │ Pipeline │ → Orchestrates stages
        └────┬─────┘
             ↓
        ┌──────────┐
        │  Output  │ → Format & display
        └──────────┘
```

### Requirements to Structure Mapping

| FR Category | Package | Key Files |
|-------------|---------|-----------|
| Command Execution (FR1-4) | `executor/` | `posix.go`, `executor.go` |
| JSON Pipeline (FR5-10) | `pipeline/`, `operators/` | `pipeline.go`, `where.go`, `select.go`, `map.go` |
| JSON Literals (FR11-13) | `lexer/`, `parser/` | `lexer.go`, `grammar.y`, `ast.go` |
| Interactive Shell (FR14-17) | `shell/` | `shell.go`, `readline.go`, `history.go` |
| Output Formatting (FR18-20) | `output/` | `formatter.go`, `pretty.go`, `color.go` |
| Error Handling (FR21-23) | `value/` | `error.go` |
| Configuration (FR24-26) | `config/` | `config.go`, `defaults.go` |

### Package Dependencies

```
cmd/conciglia
    └── internal/shell
            ├── internal/pipeline
            │       ├── internal/executor
            │       │       ├── internal/parser
            │       │       │       └── internal/lexer
            │       │       └── internal/value
            │       └── internal/operators
            │               └── internal/value
            ├── internal/config
            └── internal/output
                    └── internal/value
```

### Integration Points

**Internal Communication:**
- All packages communicate via Go interfaces
- `Value` interface is the universal data type
- No global state; configuration passed explicitly

**External Integrations:**
- OS process execution via `os/exec`
- Terminal handling via `golang.org/x/term`
- File system for history and config

### Development Workflow

**Build:**
```bash
make build      # Build binary
make test       # Run all tests
make lint       # Run golangci-lint
make generate   # Run goyacc
```

**Release:**
```bash
make release    # goreleaser (creates Linux binaries)
```

---

## Architecture Validation Results

### Coherence Validation ✅

**Decision Compatibility:**
All technology choices work together without conflicts:
- Go + goyacc = native code generation
- stdlib-only + custom readline = consistent constraint
- Value interface + streaming = unified data model

**Pattern Consistency:**
Implementation patterns align with Go conventions:
- Package naming follows Go standards
- Error handling uses established Go patterns
- Testing follows Go idioms (table-driven, co-located)

**Structure Alignment:**
Project structure supports all architectural decisions:
- `internal/` packages enforce boundaries
- Clear separation of concerns across packages
- Data flow follows package dependencies

### Requirements Coverage Validation ✅

**Functional Requirements Coverage:**

| FR Category | Package | Status |
|-------------|---------|--------|
| Command Execution (FR1-4) | `executor/` | ✅ Covered |
| JSON Pipeline (FR5-10) | `pipeline/`, `operators/` | ✅ Covered |
| JSON Literals (FR11-13) | `lexer/`, `parser/` | ✅ Covered |
| Interactive Shell (FR14-17) | `shell/` | ✅ Covered |
| Output Formatting (FR18-20) | `output/` | ✅ Covered |
| Error Handling (FR21-23) | `value/` | ✅ Covered |
| Configuration (FR24-26) | `config/` | ✅ Covered |

**Non-Functional Requirements Coverage:**

| NFR | Architectural Support | Status |
|-----|----------------------|--------|
| < 50ms startup | No reflection, minimal init, static binary | ✅ |
| Stream-based JSON | `value/stream.go`, top-level array streaming | ✅ |
| 95% POSIX compat | `executor/posix.go` fork/exec passthrough | ✅ |
| Single binary | goreleaser static linking, no cgo | ✅ |
| < 50MB memory | Streaming design, no full JSON load | ✅ |

### Implementation Readiness Validation ✅

**Decision Completeness:**
- ✅ All critical decisions documented
- ✅ Technology stack fully specified (Go, goyacc, stdlib)
- ✅ Integration patterns defined (Value interface)

**Structure Completeness:**
- ✅ All packages and files defined
- ✅ Package dependencies mapped
- ✅ Data flow documented

**Pattern Completeness:**
- ✅ Naming conventions established
- ✅ Error message format specified
- ✅ JSON field access syntax defined (jq-compatible)
- ✅ Null/empty handling rules documented

### Gap Analysis Results

**No Critical Gaps**

**Minor Gaps (Non-Blocking):**

| Gap | Priority | Resolution Path |
|-----|----------|-----------------|
| goyacc grammar details | Medium | Define during parser implementation |
| Value interface method signatures | Medium | Refine during value package implementation |
| Operator expression syntax | Low | Extend as operators are implemented |

### Architecture Completeness Checklist

**✅ Requirements Analysis**
- [x] Project context thoroughly analyzed
- [x] Scale and complexity assessed (Medium)
- [x] Technical constraints identified (stdlib-only, < 50ms)
- [x] Cross-cutting concerns mapped (JSON detection, error handling)

**✅ Architectural Decisions**
- [x] Parser: goyacc
- [x] Execution: Hybrid (POSIX passthrough + built-in operators)
- [x] Data flow: Value interface with streaming
- [x] JSON: Stream top-level arrays
- [x] Readline: Custom (golang.org/x/term)
- [x] Errors: Errors as Value
- [x] Config: JSON format

**✅ Implementation Patterns**
- [x] Go naming conventions
- [x] Error message format (helpful with suggestions)
- [x] JSON field access syntax (jq-compatible)
- [x] Null/empty handling (permissive with warnings)
- [x] Output formatting (TTY detection)

**✅ Project Structure**
- [x] Complete directory structure defined
- [x] All packages specified with purpose
- [x] Package dependencies mapped
- [x] FR to package mapping complete

### Architecture Readiness Assessment

**Overall Status:** ✅ READY FOR IMPLEMENTATION

**Confidence Level:** High

**Key Strengths:**
1. Clean separation of concerns across packages
2. Unified Value interface for data flow
3. jq-compatible syntax reduces learning curve
4. Permissive error handling improves UX
5. stdlib-only ensures portability

**Areas for Future Enhancement:**
1. Plugin system (post-MVP)
2. Advanced scripting (functions, control flow)
3. Additional operators (`sort`, `unique`, `count`)
4. Shell completions

### Implementation Handoff

**AI Agent Guidelines:**
- Follow all architectural decisions exactly as documented
- Use implementation patterns consistently across all components
- Respect package boundaries (no circular imports)
- Refer to this document for all architectural questions
- Use jq-compatible syntax for field access
- Implement permissive error handling with warnings

**Implementation Sequence:**
1. Project initialization (go.mod, structure)
2. Lexer package (tokenization)
3. Parser package (goyacc grammar, AST)
4. Value package (interface and types)
5. Executor package (POSIX + built-in dispatch)
6. Operators package (where, select, map)
7. Pipeline package (stage orchestration)
8. Output package (formatting)
9. Shell package (interactive REPL)
10. Config package (JSON config loading)

---

## Architecture Completion Summary

### Workflow Completion

**Architecture Decision Workflow:** COMPLETED ✅
**Total Steps Completed:** 8
**Date Completed:** 2025-12-04
**Document Location:** specs/architecture.md

### Final Architecture Deliverables

**📋 Complete Architecture Document**
- All architectural decisions documented with specific rationale
- Implementation patterns ensuring AI agent consistency
- Complete project structure with all files and directories
- Requirements to architecture mapping
- Validation confirming coherence and completeness

**🏗️ Implementation Ready Foundation**
- 7 core architectural decisions made
- 5 implementation patterns defined
- 10 internal packages specified
- 26 functional requirements fully supported

**📚 AI Agent Implementation Guide**
- Technology stack: Go + goyacc + stdlib
- Consistency rules preventing implementation conflicts
- Project structure with clear package boundaries
- jq-compatible syntax for field access

### Implementation Handoff

**For AI Agents:**
This architecture document is your complete guide for implementing Conciglia. Follow all decisions, patterns, and structures exactly as documented.

**First Implementation Priority:**
```bash
mkdir conciglia && cd conciglia
go mod init github.com/youruser/conciglia
mkdir -p cmd/conciglia internal/{lexer,parser,value,executor,pipeline,operators,shell,output,config}
```

**Development Sequence:**
1. Initialize project structure
2. Implement lexer (tokenization)
3. Implement parser (goyacc grammar, AST)
4. Implement value package (Value interface)
5. Implement executor (POSIX + dispatch)
6. Implement operators (where, select, map)
7. Implement pipeline (orchestration)
8. Implement output (formatting)
9. Implement shell (interactive REPL)
10. Implement config (JSON loading)

---

**Architecture Status:** READY FOR IMPLEMENTATION ✅

