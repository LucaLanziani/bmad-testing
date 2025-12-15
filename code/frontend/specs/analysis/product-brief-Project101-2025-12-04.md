---
stepsCompleted: [1, 2, 3, 4, 5]
inputDocuments: []
workflowType: 'product-brief'
lastStep: 4
project_name: 'Conciglia'
user_name: 'Nearformer'
date: '2025-12-04'
---

# Product Brief: Conciglia

**Date:** 2025-12-04
**Author:** Nearformer

---

## Executive Summary

Conciglia is a next-generation shell that combines POSIX compatibility with modern structured data handling. Built in Go for performance and portability, it treats JSON as a first-class citizen while remaining familiar to users of traditional shells. Conciglia eliminates the friction between command-line text processing and modern data formats, making developers more productive without abandoning decades of Unix wisdom.

---

## Core Vision

### Problem Statement

Traditional shells like bash and zsh treat everything as text strings, forcing developers to constantly parse, grep, awk, and sed their way through data. This creates fragile pipelines, cryptic one-liners, and scripts that break on edge cases. Meanwhile, modern development is JSON-everywhere—APIs, configs, logs, infrastructure-as-code—yet shells remain stuck in the 1970s.

### Problem Impact

- Developers waste hours writing fragile text-parsing pipelines
- Scripts break silently on unexpected input formats
- JSON manipulation requires external tools (`jq`) or escaping to other languages
- The gap between "quick shell command" and "real script" is unnecessarily wide
- New developers struggle with arcane syntax while experienced ones work around limitations

### Why Existing Solutions Fall Short

| Shell | Issue |
|-------|-------|
| **Bash/Zsh** | Text-only, archaic syntax, poor error handling |
| **Fish** | Better UX, but still text-based, not POSIX |
| **Nushell** | Great structured data, but not POSIX-compatible (adoption barrier) |
| **PowerShell** | Object pipelines, but verbose syntax, Windows-centric history |
| **Xonsh** | Python dependency, slow startup |

### Proposed Solution

**Conciglia** bridges the gap: a shell that speaks POSIX for compatibility but thinks in JSON for power. Developers can use familiar commands while seamlessly working with structured data. Written in Go for instant startup, single-binary deployment, and true cross-platform consistency.

### Key Differentiators

1. **POSIX + JSON**: Run existing scripts AND work with structured data natively
2. **Go-powered**: Fast startup, no runtime dependencies, single binary
3. **JSON pipelines**: Pipe structured data between commands, not just text
4. **Gradual adoption**: Use as much or as little of the new features as you want
5. **Helpful errors**: Modern error messages that actually help you fix problems

---

## Target Users

### Primary Users

#### Power User ("Max")
**Profile:** Senior developer or tech lead with 8+ years experience. Lives in the terminal, maintains dotfiles, has strong opinions on tooling. Expert in bash/zsh but constantly frustrated by limitations.

**Pain Points:**
- Writes complex pipelines that break on edge cases
- Knows `jq`, `awk`, `sed` by heart but wishes they didn't need to
- Switches to Python/Node for anything beyond simple text processing
- Maintains a folder of "good enough" bash scripts

**Success Vision:** Clean, readable JSON pipelines that feel natural. Commands like `curl api | where .active | select .email` instead of `curl api | jq '.[] | select(.active==true) | .email'`.

#### Full-Stack Engineer ("Priya")
**Profile:** Mid-to-senior engineer working across frontend, backend, and infrastructure. Constantly context-switching between React, Node APIs, Docker, and Terraform. JSON is her native data format.

**Pain Points:**
- Copies JSON to VS Code to inspect it because terminal tools are awkward
- Has a bookmark folder of "bash cheatsheets"
- Dreads writing deployment scripts
- Wishes shell was as ergonomic as modern programming languages

**Success Vision:** Zero-config start (like Fish), JSON just works without learning `jq`, errors actually explain what went wrong, existing bash knowledge transfers seamlessly.

### Secondary Users

N/A for initial release. Focus on individual developers first.

### User Journey

**Adoption Model: Individual Developer Champion**

1. **Discovery:** Developer sees a cool Conciglia one-liner on Twitter/Reddit/HN, thinks "that's so much cleaner than bash"
2. **Trial:** Downloads single binary, tries it on a real task, immediately productive due to POSIX compatibility
3. **Aha Moment:** First time piping JSON between commands without `jq` - "why isn't every shell like this?"
4. **Daily Driver:** Gradually replaces default shell, starts writing scripts in Conciglia
5. **Evangelism:** Shares one-liners with team, writes a team utility script in Conciglia
6. **Team Adoption:** Organic spread as teammates see the benefits firsthand

---

## Success Metrics

### User Success Metrics

| Metric | Description | Target |
|--------|-------------|--------|
| **Shell Adoption** | User sets Conciglia as default shell | Qualitative signal of commitment |
| **JSON Pipeline Success** | User executes JSON pipeline without external tools | First-session achievement |
| **Daily Active Usage** | User returns to Conciglia daily | Retention indicator |
| **Script Migration** | User rewrites a bash script in Conciglia | Deep adoption signal |

**Aha Moment:** User successfully pipes JSON through 3+ commands in their first session without reaching for `jq` - realizes "this is how shells should work."

### Business Objectives

**Primary Goal:** Build awareness and adoption in the developer community as a credible, useful open-source project.

| Objective | Metric | v1.0 Target |
|-----------|--------|-------------|
| **Visibility** | GitHub Stars | 100 ⭐ |
| **Adoption** | Downloads/Installs | Track growth |
| **Community** | Forks & Contributors | Organic growth |
| **Engagement** | Issues & Discussions | Active community |

**Monetization:** Not a goal. Pure open-source community project.

### Key Performance Indicators

| KPI | Measurement | Timeframe |
|-----|-------------|-----------|
| **GitHub Stars** | Primary success metric | 100 by v1.0 release |
| **First-week downloads** | Release momentum | Track per release |
| **Weekly active users** | Retention signal | Post-launch tracking |
| **Community contributions** | PRs from non-maintainers | Ongoing |

**Success Definition:** Reaching 100 GitHub stars indicates meaningful developer interest and validates the product-market fit for a POSIX-compatible, JSON-native shell.

---

## MVP Scope

### Core Features

| Feature | Description | Why Essential |
|---------|-------------|---------------|
| **POSIX Command Execution** | Run existing shell commands seamlessly | Compatibility is key - users keep their workflows |
| **JSON Literal Support** | `{"key": "value"}` works directly in pipelines | Core differentiator |
| **JSON-Aware Pipe** | Commands can output/receive structured data | The "aha!" moment |
| **Basic JSON Operators** | `select`, `where`, `map` for filtering/transforming | Replaces `jq` for common ops |
| **Pretty JSON Output** | Auto-format JSON with syntax highlighting | Developer experience |
| **Helpful Error Messages** | Clear, actionable errors (not bash's cryptic output) | Reduces friction |
| **Interactive Mode** | Basic REPL with command history | Daily driver readiness |
| **Single Binary Install** | `curl \| sh` or direct download → run | Zero-friction adoption |

### Out of Scope for MVP

| Feature | Rationale for Deferring |
|---------|------------------------|
| Plugin System | Adds complexity; build community first |
| Custom Themes/Prompts | Fish-like customization can wait |
| Remote Execution | SSH integration is v2+ territory |
| Shell Scripting Extensions | Focus on interactive use first |
| Completions Ecosystem | Requires community momentum |
| Windows Native Support | Linux/macOS first, Windows later |

### MVP Success Criteria

| Criterion | Validation Method |
|-----------|-------------------|
| **Core Functionality Works** | User can pipe JSON through 3+ commands |
| **POSIX Compatibility** | Existing scripts run without modification |
| **Installation Success** | Single binary installs in < 30 seconds |
| **Daily Driver Viable** | User can set as default shell |
| **Community Interest** | 100 GitHub stars achieved |

**Go/No-Go Decision Point:** After reaching 100 stars and gathering user feedback, evaluate:
- Which features are most requested?
- What's blocking daily driver adoption?
- Is the JSON-native approach resonating?

### Future Vision

**Year 1-2 (Post-MVP):**
- Plugin system for extensibility
- Rich completions ecosystem
- Advanced JSON operators (joins, aggregations)
- Shell scripting with structured data

**Year 2-3 (Growth):**
- Windows native support
- Remote execution capabilities
- Team/enterprise features (shared configs)
- IDE integrations

**Long-term Vision:**
Conciglia becomes the default shell for developers who work with APIs, microservices, and modern toolchains - the shell that finally speaks the same language as modern software.

