---
stepsCompleted: [1, 2, 3, 4, 7, 8, 9, 10, 11]
status: complete
inputDocuments:
  - specs/analysis/product-brief-Project101-2025-12-04.md
workflowType: 'prd'
lastStep: 0
project_name: 'Conciglia'
user_name: 'Nearformer'
date: '2025-12-04'
---

# Product Requirements Document - Conciglia

**Author:** Nearformer
**Date:** 2025-12-04

---

## Executive Summary

**Conciglia** is a modern shell that brings structured data handling to the command line without abandoning decades of Unix wisdom. It's the shell for developers who are tired of wrestling with text parsing but can't afford to learn an entirely new tool.

### The Problem

Traditional shells treat everything as text strings. In a world where APIs return JSON, configs are YAML, and logs are structured, developers spend countless hours writing fragile parsing pipelines with `grep`, `awk`, `sed`, and `jq`. The tools exist, but they don't speak the same language.

Meanwhile, modern shells like Nushell offer structured data—but demand you abandon POSIX compatibility and learn something entirely new. For busy developers, that's a non-starter.

### The Solution

Conciglia bridges the gap: **POSIX-compatible with JSON-native pipelines**. Run your existing scripts unchanged while gaining the ability to work with structured data naturally. Built in Go for instant startup and single-binary deployment.

### What Makes This Special

| Differentiator | Why It Matters |
|----------------|----------------|
| **Zero Migration Cost** | POSIX compatibility means your existing scripts, muscle memory, and workflows transfer instantly. Unlike Nushell, you don't have to choose. |
| **JSON as First-Class Citizen** | `curl api \| where .active \| select .email` — no `jq`, no escaping, no friction. |
| **Go-Powered Performance** | Instant startup, no runtime dependencies, single binary for any platform. |
| **Gradual Adoption** | Use 1% or 100% of the new features. Conciglia meets you where you are. |
| **Helpful Errors** | Modern error messages that explain what went wrong and how to fix it. |

## Project Classification

**Technical Type:** CLI Tool (Shell)
**Domain:** Developer Tooling
**Complexity:** Low

*A command-line shell targeting power users and full-stack engineers who work with JSON-heavy workflows. No regulatory or compliance requirements. Standard open-source security considerations apply.*

---

## Success Criteria

### User Success

| Milestone | Definition | Measurement |
|-----------|------------|-------------|
| **First JSON Pipeline** | User pipes JSON through 3+ commands without `jq` | First-session achievement |
| **Daily Driver** | User sets Conciglia as default shell | Commitment signal |
| **Script Migration** | User rewrites at least one bash script in Conciglia | Deep adoption |
| **Evangelism** | User shares a Conciglia one-liner with teammates | Organic growth indicator |

**"Aha!" Moment:** User runs `curl api | where .status == "active" | select .name, .email` and realizes they didn't need `jq`, `grep`, or `awk`.

### Business Success

| Objective | Metric | v1.0 Target |
|-----------|--------|-------------|
| **Visibility** | GitHub Stars | 100 ⭐ |
| **Adoption** | Downloads | Track growth |
| **Community** | Contributors | Any external PRs |
| **Engagement** | Issues opened | Active feedback |

**Monetization:** Not a goal. Pure open-source community project.

### Technical Success

| Criterion | Target | Rationale |
|-----------|--------|-----------|
| **Startup Time** | < 50ms | Faster than zsh, competitive with bash |
| **POSIX Compatibility** | 95% of bash scripts work unchanged | Zero migration friction |
| **Binary Size** | < 20MB | Single-binary distribution |
| **Platform Support** | Linux (v1.0), macOS (v1.1) | Linux first, macOS fast-follow |
| **Memory Usage** | < 50MB idle | Comparable to existing shells |

---

## Product Scope

### MVP - Minimum Viable Product

**Must ship for v1.0:**

| Feature | Success Criterion |
|---------|-------------------|
| POSIX command execution | 95% bash script compatibility |
| JSON literal parsing | `{"key": "value"}` works in pipelines |
| JSON-aware piping | Structured data flows between commands |
| Basic JSON operators | `select`, `where`, `map` functional |
| Pretty JSON output | Syntax-highlighted, auto-formatted |
| Helpful error messages | Clear, actionable, context-aware |
| Interactive REPL | History, basic line editing |
| Single binary | < 20MB, zero dependencies |
| Linux support | x86_64 and ARM64 |

### Growth Features (Post-MVP)

| Feature | Target Release |
|---------|----------------|
| macOS support | v1.1 |
| Tab completions | v1.2 |
| Custom prompts | v1.2 |
| Advanced JSON operators (joins, aggregations) | v1.3 |
| Shell scripting extensions | v2.0 |

### Vision (Future)

| Capability | Timeline |
|------------|----------|
| Windows native support | Year 2 |
| Plugin system | Year 2 |
| Remote execution (SSH) | Year 2 |
| IDE integrations | Year 2-3 |
| Team/enterprise features | Year 3+ |

**Long-term:** Conciglia becomes the default shell for developers who work with APIs, microservices, and modern toolchains.

---

## User Journeys

### Journey 1: Max - The Terminal Purist's Epiphany

Max is a senior backend engineer who's spent 12 years mastering the command line. His dotfiles are a work of art, his `.zshrc` is 400 lines of carefully tuned aliases, and he genuinely enjoys writing `awk` one-liners that would make most developers weep.

But lately, Max has been frustrated. Every microservice speaks JSON. Every API returns JSON. Every config file is JSON or YAML. And yet, his beloved shell still thinks it's 1975 and everything is a string. He's got `jq` installed, but chaining it with `grep` and `sed` feels like duct-taping three different languages together.

One evening, scrolling through Hacker News, Max sees a post: *"Conciglia: A POSIX shell that speaks JSON natively."* He's skeptical — he's seen "shell replacements" come and go. Fish was nice but not POSIX. Nushell looked cool but he'd have to relearn everything.

He downloads the binary, drops it in `/usr/local/bin`, and types his first command:

```
curl https://api.github.com/users/maxdev/repos | where .stargazers_count > 10 | select .name, .language
```

Max stares at his terminal. Clean output. No `jq`. No escaping. No bullshit.

He spends the next two hours converting his ugliest bash scripts. The one that parses Kubernetes pod status? 40 lines becomes 6. The AWS inventory script with nested `jq` calls? It's now readable by humans.

By midnight, Max has changed his default shell. By the end of the week, he's posted three Conciglia one-liners to Twitter that go mildly viral in dev circles. He files his first issue on GitHub: a feature request, not a bug report.

**Capabilities Revealed:**
- POSIX command execution (existing tools work)
- JSON piping without external tools
- `where` and `select` operators
- Single binary installation
- Interactive REPL

---

### Journey 2: Priya - From Context-Switching Hell to Flow State

Priya is a full-stack engineer at a Series B startup. On any given day, she touches React, Node, PostgreSQL, Redis, Docker, Terraform, and at least three different AWS services. Her browser has 47 tabs open. Her VS Code has 12 files. Her terminal has 8 tmux panes.

The constant context-switching is killing her flow. The worst part? Every time she needs to inspect JSON from an API response, she copies it to VS Code, formats it, searches for the field she needs, then switches back to the terminal. It takes 30 seconds but breaks her concentration for 5 minutes.

A colleague mentions Conciglia in Slack. Priya is intrigued but hesitant — she doesn't have time to learn a new shell.

During her lunch break, she installs it and runs one command:

```
curl localhost:3000/api/users | where .role == "admin"
```

The output is pretty-printed, syntax-highlighted, and shows exactly what she needs. No copy-paste. No VS Code. No broken flow.

Over the next week, Priya discovers she's not learning a new shell — she's using her existing shell with superpowers. Her bash muscle memory works. Her aliases work. But now she can pipe JSON like she pipes text.

The breakthrough comes during an incident. A service is returning malformed data. Instead of her usual dance between terminal, VS Code, and Postman, Priya debugs entirely in Conciglia:

```
curl api.internal/health | where .services[].status != "healthy" | select .services[].name, .services[].error
```

She finds the broken service in 10 seconds. Her team lead asks how she diagnosed it so fast. She shares the command. Three teammates install Conciglia that afternoon.

**Capabilities Revealed:**
- Pretty JSON output with syntax highlighting
- Zero learning curve for POSIX users
- Array access in pipelines
- Gradual adoption path
- Team viral adoption through visible productivity

---

### Journey Requirements Summary

| Capability | Max's Journey | Priya's Journey |
|------------|---------------|-----------------|
| POSIX command execution | ✅ Aliases/scripts work | ✅ Muscle memory transfers |
| JSON piping | ✅ Core epiphany | ✅ Core value |
| `where` filtering | ✅ First command | ✅ Daily use |
| `select` projection | ✅ First command | ✅ Incident response |
| Pretty output | — | ✅ Flow state enabler |
| Single binary install | ✅ Friction-free trial | ✅ Easy team adoption |
| Array/nested access | — | ✅ Complex debugging |
| Error messages | — | (Implied for debugging) |

---

## CLI Tool Specific Requirements

### Project-Type Overview

Conciglia is a **command-line shell** that operates in two primary modes:

| Mode | Description | Key Features |
|------|-------------|--------------|
| **Interactive** | Daily driver REPL | History, line editing, syntax highlighting, pretty output |
| **Scriptable** | Automation & pipelines | POSIX compatibility, machine-readable output, exit codes |

### Command Structure

**Core Command Patterns:**

```
# POSIX commands pass through unchanged
ls -la
cat file.txt | grep pattern

# JSON-aware piping with built-in operators
curl api | where .status == "active" | select .name, .email

# Mixed pipelines
kubectl get pods -o json | where .status.phase != "Running" | select .metadata.name
```

**Built-in Operators (MVP):**

| Operator | Purpose | Example |
|----------|---------|---------|
| `where` | Filter JSON by condition | `where .age > 21` |
| `select` | Project specific fields | `select .name, .email` |
| `map` | Transform values | `map .name \| uppercase` |

### Output Formats

| Format | Flag | Default | Use Case |
|--------|------|---------|----------|
| **Pretty JSON** | `--pretty` / `-p` | ✅ Interactive | Human inspection |
| **Compact JSON** | `--json` / `-j` | — | Piping to tools |
| **Plain Text** | `--text` / `-t` | ✅ POSIX commands | Legacy compatibility |
| **Table** | `--table` | — | List display (v1.2+) |

**Auto-detection:** Interactive terminal → pretty output. Piped → compact output.

### Configuration Schema

**Config File Location:** `~/.config/conciglia/config.toml`

```toml
# ~/.config/conciglia/config.toml

[output]
default_format = "pretty"  # pretty | compact | text
colors = true
max_depth = 10

[history]
file = "~/.local/share/conciglia/history"
max_entries = 10000

[compatibility]
posix_mode = true
strict_json = false
```

**Environment Variable Overrides:**
- `CONCIGLIA_OUTPUT_FORMAT` — Override default output format
- `CONCIGLIA_NO_COLOR` — Disable colors
- `CONCIGLIA_HISTORY_FILE` — Custom history location

### Scripting Support

**Script Extension:** `.conch`

**Shebang:**
```bash
#!/usr/bin/env conciglia
```

**Exit Codes:**
| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | General error |
| 2 | Invalid JSON in pipeline |
| 126 | Command not executable |
| 127 | Command not found |

### Implementation Considerations

**Performance Targets:**
- Startup: < 50ms
- JSON parsing: Stream large files, don't load all into memory
- Memory: < 50MB idle

**Platform Targets (MVP):**
- Linux x86_64
- Linux ARM64

**Dependencies:**
- None (single static binary)
- Go standard library only for v1.0

---

## Project Scoping & Phased Development

### MVP Strategy & Philosophy

**MVP Approach:** Problem-Solving MVP — Solve the JSON-in-shell problem with minimal features

**Rationale:** The core insight is that developers need JSON pipelines to "just work" without learning a new shell language. Everything else can wait.

**Resource Requirements:** Solo developer or small team (1-2 people)

### MVP Feature Set (Phase 1 — v1.0)

**Core User Journeys Supported:**
- Max: Parse API responses without `jq`
- Priya: Filter Kubernetes output in one command

**Must-Have Capabilities:**

| Feature | Why Essential |
|---------|---------------|
| POSIX command execution | Zero migration cost |
| JSON literal parsing | `{"key": "value"}` in pipelines |
| JSON-aware pipe operator | The "aha!" moment |
| `where` operator | Filter JSON arrays |
| `select` operator | Project specific fields |
| `map` operator | Transform values |
| Basic readline/history | Usable as daily driver |
| Helpful error messages | Distinguish from bash |

**Explicit Out of Scope (v1.0):**
- Custom syntax/grammar beyond POSIX + operators
- Plugin/extension system
- Advanced scripting (functions, control flow)
- IDE/editor integration
- Windows support
- macOS support (fast-follow)
- Table output format
- Shell completions

### Post-MVP Features

**Phase 2 (v1.1-1.2) — Polish:**
- macOS support
- Table output format
- Shell completions (bash, zsh, fish)
- `sort`, `unique`, `count` operators
- YAML support

**Phase 3 (v2.0) — Expansion:**
- Functions and control flow
- Plugin system
- IDE integration
- Windows support
- Custom prompt theming

### Risk Mitigation Strategy

**Technical Risks:**
| Risk | Mitigation |
|------|------------|
| POSIX compatibility gaps | Use established POSIX test suites, prioritize common commands |
| JSON parsing performance | Stream-based parsing, lazy evaluation |
| Startup time > 50ms | Profile early, avoid reflection, static linking |

**Market Risks:**
| Risk | Mitigation |
|------|------------|
| "Why not just use jq?" | Demo the pipeline experience, not individual features |
| Nushell comparison | Emphasize POSIX compat as key differentiator |

**Resource Risks:**
| Risk | Mitigation |
|------|------------|
| Scope creep | MVP list is frozen; nice-to-haves go to Phase 2 |
| Time to first release | Focus on Linux-only for v1.0 |

---

## Functional Requirements

### Command Execution

- **FR1:** User can execute any POSIX-compliant command unchanged
- **FR2:** User can pipe output between POSIX commands
- **FR3:** User can execute commands in background with `&`
- **FR4:** User can redirect output to files with `>` and `>>`

### JSON Pipeline Operations

- **FR5:** User can pipe JSON output from any command
- **FR6:** User can filter JSON arrays with `where` operator
- **FR7:** User can project specific fields with `select` operator
- **FR8:** User can transform values with `map` operator
- **FR9:** User can chain multiple JSON operators in a single pipeline
- **FR10:** System auto-detects JSON vs text output

### JSON Literal Support

- **FR11:** User can type JSON literals directly in commands
- **FR12:** User can use JSON literals as input to operators
- **FR13:** System validates JSON syntax and reports clear errors

### Interactive Shell

- **FR14:** User can navigate command history with arrow keys
- **FR15:** User can edit current line with readline controls
- **FR16:** User can search history with Ctrl+R
- **FR17:** System saves history across sessions

### Output Formatting

- **FR18:** System displays pretty JSON in interactive mode
- **FR19:** System displays compact JSON when piped
- **FR20:** User can force output format with flags (`--pretty`, `--json`, `--text`)

### Error Handling

- **FR21:** System displays helpful error messages for invalid JSON
- **FR22:** System displays helpful error messages for invalid operators
- **FR23:** System returns appropriate exit codes

### Configuration

- **FR24:** User can configure default output format
- **FR25:** User can configure history file location
- **FR26:** User can enable/disable colors

---

## Non-Functional Requirements

### Performance

| Requirement | Target | Measurement |
|-------------|--------|-------------|
| **Startup Time** | < 50ms | Cold start to prompt |
| **Command Execution** | < 10ms overhead | vs direct execution |
| **JSON Parsing** | Stream-based | Handle files > 100MB without loading to memory |
| **Memory Usage (idle)** | < 50MB | RSS after startup |
| **Memory Usage (large JSON)** | < 2x file size | During pipeline processing |

### Compatibility

| Requirement | Target | Measurement |
|-------------|--------|-------------|
| **POSIX Commands** | 95% compatibility | Bash test suite pass rate |
| **Exit Codes** | POSIX-compliant | Standard exit code semantics |
| **Signal Handling** | Standard signals | SIGINT, SIGTERM, SIGPIPE |

### Reliability

| Requirement | Target | Measurement |
|-------------|--------|-------------|
| **Crash Recovery** | Graceful | History preserved on crash |
| **Invalid JSON** | Non-fatal | Clear error, continue pipeline where possible |
| **Binary Stability** | Static linking | No runtime dependencies |

### Portability

| Requirement | Target | Measurement |
|-------------|--------|-------------|
| **Linux x86_64** | Full support | Primary target |
| **Linux ARM64** | Full support | Primary target |
| **Single Binary** | Yes | No external dependencies |
| **Config Location** | XDG Base Directory | `~/.config/conciglia/` |

