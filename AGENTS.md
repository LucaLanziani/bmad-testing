# AGENTS.md - AI Agent Context Guide

This document helps AI agents understand how to work with this repository.

## Repository Purpose

This is a **BMAD configuration hub**, not a code repository.

- Stores project specifications and shared agent definitions
- Actual project code lives in `./code/` folder (symlinks to external repos)
- BMAD configuration is symlinked to code repositories

## Critical Directory Structure

| Directory | Purpose |
|-----------|---------|
| `code/` | **Symlinks to actual code repos** (backend, frontend, infra, test) - look here for implementation work |
| `specs/` | Project specifications and sprint artifacts |
| `.bmad/` | BMAD core configuration (symlinked to code repos) |
| `bmad-custom/agents/` | Custom agent source definitions |
| `.cursor/rules/bmad/` | Cursor IDE agent rules |
| `.github/agents/` | GitHub Copilot agent definitions |

## Important Notes for AI Agents

1. **When evaluating project status**, examine `./code/` contents - that's where the actual code lives
2. **Specs and documentation** belong in `specs/`
3. **Custom agents** are defined in `bmad-custom/` but installed to `.bmad/custom/`
4. This repo uses symlinks - `.bmad/`, `specs/`, and `.cursor/rules/bmad/` are shared across all linked code repos


## How to Use Agents

### In Cursor IDE

Reference agents in chat using the `@` syntax:

```
@.bmad/bmm/agents/architect    # Activate architect agent
@.bmad/bmm/agents/dev          # Activate dev agent
@.bmad/custom/agents/infra-engineer  # Activate custom infra agent
@.bmad/index                   # See all available agents
```

- Rules are Manual type - only loaded when explicitly referenced
- Can combine multiple agents: `@.bmad/bmm/agents/dev @.bmad/bmm/agents/tea`

### In VS Code with GitHub Copilot

BMAD agents are installed as custom agents in `.github/agents/`.

**Using @ mention (recommended):**

```
@bmd-custom-bmm-architect How should I structure this service?
@bmd-custom-bmm-dev *develop-story
@bmd-custom-bmm-pm *workflow-status
```

**Available agents:**
- `@bmd-custom-bmm-pm` - Product Manager
- `@bmd-custom-bmm-analyst` - Business Analyst
- `@bmd-custom-bmm-architect` - System Architect
- `@bmd-custom-bmm-sm` - Scrum Master
- `@bmd-custom-bmm-dev` - Developer
- `@bmd-custom-bmm-tea` - Test Architect
- `@bmd-custom-bmm-ux-designer` - UX Designer
- `@bmd-custom-bmm-tech-writer` - Technical Writer
- `@bmd-custom-bmm-quick-flow-solo-dev` - Principal Engineer
- `@bmd-custom-core-bmad-master` - BMad Master (orchestrator)

### Running Workflows

Once an agent is activated, trigger workflows with the `*` prefix:

```
*workflow-status          # Check current state and get recommendations
*create-prd               # Create Product Requirements Document
*create-architecture      # Create system architecture
*sprint-planning          # Initialize sprint tracking
*develop-story            # Implement a story
*code-review              # Perform code review
```

Use `*workflow-status` from any agent to get context-aware recommendations on what to do next.

## Quick Reference

### Starting a New Project
1. Load PM or Analyst agent
2. Run `*workflow-init`
3. Follow recommended workflow based on project type

### Implementing Stories
1. Load SM agent → `*create-story`
2. Load DEV agent → `*develop-story`
3. Run `*code-review`
4. Mark done with `*story-done`

### Getting Oriented
- Run `*workflow-status` from any agent to understand current project state
- Check `specs/` for existing documentation
- Check `code/` for actual implementation status

