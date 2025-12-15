---
workflowName: create-epics-and-stories-github
targetModule: bmm
workflowType: integration
flowPattern: sequential-linear
date: 2025-12-04
user_name: Nearformer
stepsCompleted: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12]
lastStep: complete
status: complete
---

# Workflow Creation Plan: create-epics-and-stories-github

**Created:** 2025-12-04
**Author:** Nearformer
**Module:** bmm
**Type:** Integration Workflow (extends existing workflow with GitHub Projects integration)

## Executive Summary

This workflow extends the existing `create-epics-and-stories` workflow to create epics and stories directly on GitHub Projects. It will:

1. Leverage existing PRD, Architecture, and UX documents (same as base workflow)
2. Generate comprehensive epics and stories with full technical context
3. **NEW**: Push epics and stories to GitHub Projects via the GitHub API
4. Support GitHub Projects (v2) with proper field mapping

### Problem Solved

Currently, teams must manually transfer epics/stories from markdown documents to GitHub Projects. This workflow automates that process, ensuring:
- Consistent story formatting in GitHub
- Proper epic→story relationships
- Field mapping (priority, labels, status)
- No manual copy-paste errors

### Target Users

- Project Managers managing work in GitHub Projects
- Developers who want stories in their GitHub workflow
- Teams using GitHub Issues + Projects for sprint tracking

## Requirements Analysis

### Workflow Purpose

- **Problem to Solve**: Automate creation of epics and stories directly on GitHub, eliminating manual transfer from documents to issue tracker
- **Primary Users**: Project managers, developers, and teams using GitHub Issues for sprint tracking
- **Main Outcome**: GitHub Milestones (epics) and Issues (stories) with proper relationships and labels
- **Usage Frequency**: Per-project or per-sprint basis

### Workflow Classification

- **Type**: Integration (generate content + external API actions)
- **Flow Pattern**: Sequential Linear (Generate all → Review → Push all)
- **Interaction Style**: Collaborative during generation, confirmation before push
- **Instruction Style**: Intent-based (conversational)
- **Autonomy Level**: Semi-autonomous with user checkpoints

### Input Requirements

- **Required Inputs**: PRD.md, Architecture.md
- **Optional Inputs**: UX Design.md (if UI exists)
- **Prerequisites**: Completed PRD and Architecture documents
- **Runtime Config**: GitHub repository (owner/repo format)

### Output Specifications

- **Primary Output**: GitHub Milestones + Issues (NO local files)
- **Intermediate Outputs**: In-memory epic/story generation for review
- **Output Format**: GitHub API entities (milestones, issues, labels)
- **Summary Report**: List of created milestones/issues with URLs

### GitHub Integration Details

- **Authentication**: `GITHUB_TOKEN` environment variable
- **Issue Structure**: Epic → GitHub Milestone, Story → Issue assigned to milestone
- **Labels**: Auto-create defaults (`epic`, `story`, priority labels)
- **Assignees**: Not supported (by design)

### Technical Constraints

- **Dependencies**: GitHub API access, valid authentication token
- **Integrations**: GitHub REST API (or GraphQL)
- **Performance Requirements**: Handle batch creation of multiple issues
- **Standalone**: Yes, callable independently

### Target Location

- **Module**: bmm
- **Folder Name**: create-epics-and-stories-github
- **Target Path**: `{project-root}/.bmad/bmm/workflows/3-solutioning/create-epics-and-stories-github`
- **Custom Location**: `{project-root}/bmad-custom/workflows/create-epics-and-stories-github`

### Success Criteria

- **Quality Metrics**: All FRs from PRD covered by stories
- **Success Indicators**: Milestones created, Issues linked to milestones, Labels applied
- **User Satisfaction**: Clean GitHub project board ready for sprint planning

## Tool Requirements Summary

### Selected Tools

| Category | Tool | Purpose |
|----------|------|---------|
| Core BMAD | `advanced-elicitation` | Optional refinement of epics/stories |
| LLM Feature | `file-io` | Read PRD, Architecture, UX documents |
| LLM Feature | `web-browsing` | GitHub REST API calls |
| Memory | `sidecar-file` | Track progress, enable resume capability |

### Integration Approach

- **GitHub API**: Direct REST API calls using native web/fetch capabilities
- **No MCP Required**: Workflow uses built-in LLM tools only
- **Authentication**: `GITHUB_TOKEN` environment variable passed to API headers

### Sidecar Memory Schema

```yaml
# .github-sync-state.yaml
workflow_run_id: "uuid"
repository: "owner/repo"
created_at: "timestamp"
milestones:
  - epic_id: 1
    github_milestone_id: 123
    title: "Epic Title"
    state: "created"
issues:
  - story_id: "1.1"
    github_issue_number: 456
    milestone_id: 123
    title: "Story Title"
    state: "created"
last_checkpoint: "epic_2_story_3"
```

## Core Tools Configuration

### Workflows & Tasks

| Tool | Status | Integration Points |
|------|--------|-------------------|
| **Party-Mode** | ❌ Excluded | Not needed |
| **Advanced Elicitation** | ✅ Optional | Quality checkpoint before GitHub push |
| **Brainstorming** | ❌ Excluded | Not needed |

### LLM Tool Features

| Tool | Status | Integration Points |
|------|--------|-------------------|
| **Web-Browsing** | ✅ Included | GitHub REST API calls |
| **File I/O** | ✅ Included | Read input documents |
| **Sub-Agents** | ❌ Excluded | Not needed |
| **Sub-Processes** | ❌ Excluded | Not needed |

### Tool-Memory

| Tool | Status | Use Case |
|------|--------|----------|
| **Sidecar File** | ✅ Included | Track GitHub entities, enable resume |

### Workflow Step Integration Map

```
Step 0: Load documents (file-io)
Step 1: Generate epics structure (in memory)
Step 2: Generate stories per epic (in memory)
Step 3: Review checkpoint (advanced-elicitation optional)
Step 4: Create labels on GitHub (web-browsing → REST API)
Step 5: Create milestones on GitHub (web-browsing → REST API, sidecar write)
Step 6: Create issues on GitHub (web-browsing → REST API, sidecar write)
Step 7: Summary report
```

## Memory Configuration

### Memory Requirements

| Type | Status | Use Case |
|------|--------|----------|
| **Sidecar File** | ✅ Selected | Track GitHub entities, enable resume |
| **Vector Database** | ❌ Not selected | Not needed |

### Sidecar File Implementation

**File:** `.github-sync-state.yaml`

**Schema:**
```yaml
workflow_run_id: "uuid"
repository: "owner/repo"
created_at: "timestamp"
status: "in_progress|completed|failed"
milestones:
  - epic_id: 1
    github_milestone_id: 123
    title: "Epic Title"
    state: "created"
issues:
  - story_id: "1.1"
    github_issue_number: 456
    milestone_id: 123
    title: "Story Title"
    state: "created"
last_checkpoint: "step_name"
```

**Resume Logic:**
- Check for existing state file on workflow start
- If incomplete, offer resume from last checkpoint
- Skip already-created entities
- Continue from last pending item

**Cleanup:** State file persists for reference; manual deletion optional

## External Tools Configuration

### MCP Integrations

**Selected Tools:** None - using native LLM capabilities

**Rationale:** Direct GitHub REST API calls via web-browsing provide all needed functionality without external dependencies.

### GitHub API Integration

**Authentication:** `GITHUB_TOKEN` environment variable

**API Endpoints Used:**
| Endpoint | Method | Purpose |
|----------|--------|---------|
| `/repos/{owner}/{repo}/labels` | POST | Create labels |
| `/repos/{owner}/{repo}/milestones` | POST | Create milestones (epics) |
| `/repos/{owner}/{repo}/issues` | POST | Create issues (stories) |
| `/repos/{owner}/{repo}/issues/{num}` | PATCH | Assign issues to milestones |

### Installation Requirements

**None** - Workflow is self-contained using:
- Native file-io for document reading
- Native web-browsing for GitHub API calls
- Sidecar file for state persistence

## Installation Guidance Configuration

### Installation Approach

**Selected Approach:** Minimal prerequisites check at workflow start
**Tools Requiring Installation:** None
**Installation Step Placement:** Step 0 (Prerequisites validation)

### User Prerequisites

1. **GitHub Personal Access Token**
   - URL: `https://github.com/settings/tokens`
   - Scopes: `repo` (full repository access)
   - Environment: `GITHUB_TOKEN`

2. **Input Documents**
   - PRD.md (required)
   - Architecture.md (required)  
   - UX Design.md (optional)

### Prerequisites Check Logic

```
Step 0: Prerequisites Check
├── Verify GITHUB_TOKEN environment variable is set
├── Test GitHub API access (GET /user)
├── Verify target repository exists and is accessible
├── Verify input documents are present
└── If any fail → provide setup guidance and exit gracefully
```

## Final Tools Configuration Summary

### Tools Inventory

| Category | Count | Tools |
|----------|-------|-------|
| Core BMAD | 1 | advanced-elicitation (optional) |
| LLM Features | 2 | file-io, web-browsing |
| Memory | 1 | sidecar-file |
| External | 0 | None |

**Total Active Tools:** 4
**Installation Required:** None

### Integration Strategy

| Aspect | Approach |
|--------|----------|
| **User Experience** | Collaborative generation, confirmation before push |
| **Checkpoint Approach** | Review after generation, before GitHub push |
| **Performance** | Sequential processing, batch API calls |
| **Installation** | Zero setup (native LLM capabilities only) |

### Ready for Design

✅ All tools configured and ready for workflow design phase.

---

## Detailed Design

### Step Structure

**Total Steps:** 9 (Steps 0-8)

| Step | Name | Goal | Menu Type |
|------|------|------|-----------|
| **0** | Prerequisites | Verify token, test API, check resume, get repo | Auto-proceed |
| **1** | Load Documents | Read PRD, Architecture, UX; extract requirements | Standard (A/P/C) |
| **2** | Design Epics | Generate ALL epic structure | Standard (A/P/C) |
| **3** | Generate Stories | Generate ALL stories for all epics | Standard (A/P/C) |
| **4** | Pre-Push Review | Complete summary, last chance to refine | Standard (A/P/C) |
| **5** | Push Labels | Create default labels on GitHub | Confirm + progress |
| **6** | Push Milestones | Create milestones (one per epic) | Confirm + progress |
| **7** | Push Issues | Create issues (stories) per batch | Confirm + progress |
| **8** | Summary | Completion report with URLs | Final |

### Interaction Design

- **Menu Pattern**: Standard A/P/C for content steps; Y/N confirm for push steps
- **User Input Points**: Step 0 (repo), Steps 1-4 (review/refine), Steps 5-7 (confirm pushes)
- **AI Autonomy Level**: Collaborative during generation, confirmation-gated for GitHub actions
- **Decision Flow**: Sequential with checkpoints; resume capability at any push step

### Data Flow Architecture

```
Step 0: Prerequisites
├── Input: GITHUB_TOKEN (env), repository (user input)
├── Check: Existing sidecar file → Offer resume or fresh start
└── Output: github_config, resume_state

Step 1: Load Documents
├── Input: PRD.md, Architecture.md, UX.md (optional)
└── Output: fr_inventory, tech_context, ux_patterns

Steps 2-3: Generate Epics & Stories
├── Input: fr_inventory, tech_context, ux_patterns
└── Output: epics[] with stories[] (in memory)

Step 4: Review
├── Input: epics[] with stories[]
└── Output: approved_epics[]

Steps 5-7: Push to GitHub
├── Input: approved_epics[], github_config
├── Checkpoint: After each milestone, after each issue batch
└── Output: created_entities{} with sidecar writes

Step 8: Summary
├── Input: created_entities{}
└── Output: Completion report
```

### File Organization

**Core Files:**
- `workflow.yaml` - Workflow configuration
- `instructions.md` - Detailed workflow instructions

**Step Files:**
```
steps/
├── step-00-prerequisites.md
├── step-01-load-documents.md
├── step-02-design-epics.md
├── step-03-generate-stories.md
├── step-04-review.md
├── step-05-push-labels.md
├── step-06-push-milestones.md
├── step-07-push-issues.md
└── step-08-summary.md
```

**Templates:**
```
templates/
├── epic-template.md
├── story-template.md
└── summary-template.md
```

### AI Role Definition

- **Primary Role**: Product decomposition specialist & GitHub integration engineer
- **Expertise Areas**: Agile story writing, GitHub API, requirements analysis
- **Communication Style**: Direct, structured, progress-focused
- **Collaboration Approach**: Collaborative during generation, confirmation-gated for actions

### Quality Assurance

**Validation Points:**
- Step 0: GITHUB_TOKEN exists, repo accessible
- Step 1: Required documents present
- Steps 2-3: All FRs covered by stories
- Steps 5-7: API responses validated

**Error Handling:**
- Missing token → Exit with setup instructions
- Repo not accessible → Exit with permission guidance
- API rate limit → Pause, show wait time, offer retry
- API error → Log to sidecar, offer retry or skip
- Network failure → Save checkpoint, offer resume later

**Recovery Strategies:**
- Sidecar file tracks exact checkpoint
- Resume from milestone X, issue Y
- Skip already-created entities

### Resume Logic

```yaml
# Sidecar checkpoint example
checkpoint:
  step: "push-issues"
  epic_index: 2
  story_index: 3
  message: "Completed Epic 2, Story 3. Next: Epic 2, Story 4"
```

### Special Features

- **Conditional Logic**: Resume detection in Step 0
- **Branch Points**: Resume vs fresh start
- **Integrations**: GitHub REST API
- **Multi-Scenario Support**: New run, resume run, error recovery

## Implementation Plan

### Build Summary

**Build Date:** 2025-12-04
**Target Location:** `bmad-custom/workflows/create-epics-and-stories-github/`

### Files Generated

| File | Type | Description |
|------|------|-------------|
| `workflow.yaml` | Config | Workflow configuration and variables |
| `instructions.md` | Main | Workflow entry point and architecture |
| `steps/step-00-prerequisites.md` | Step | Token, repo, resume check |
| `steps/step-01-load-documents.md` | Step | Load PRD, Architecture, UX |
| `steps/step-02-design-epics.md` | Step | Generate epic structure |
| `steps/step-03-generate-stories.md` | Step | Generate all stories |
| `steps/step-04-review.md` | Step | Pre-push review |
| `steps/step-05-push-labels.md` | Step | Create GitHub labels |
| `steps/step-06-push-milestones.md` | Step | Create milestones |
| `steps/step-07-push-issues.md` | Step | Create issues |
| `steps/step-08-summary.md` | Step | Completion summary |
| `templates/epic-template.md` | Template | Epic display format |
| `templates/story-template.md` | Template | Story/issue format |
| `templates/summary-template.md` | Template | Completion summary format |

**Total Files:** 14

### File Structure

```
bmad-custom/workflows/create-epics-and-stories-github/
├── workflow.yaml
├── instructions.md
├── steps/
│   ├── step-00-prerequisites.md
│   ├── step-01-load-documents.md
│   ├── step-02-design-epics.md
│   ├── step-03-generate-stories.md
│   ├── step-04-review.md
│   ├── step-05-push-labels.md
│   ├── step-06-push-milestones.md
│   ├── step-07-push-issues.md
│   └── step-08-summary.md
└── templates/
    ├── epic-template.md
    ├── story-template.md
    └── summary-template.md
```

## Review and Validation

### Validation Results

| Category | Result |
|----------|--------|
| Configuration validation | ✅ PASSED |
| Step compliance | ✅ PASSED |
| Cross-file consistency | ✅ PASSED |
| Requirements verification | ✅ PASSED |
| Best practices | ✅ PASSED |

### Issues Found

**Critical Issues:** None
**Warnings:** None
**Suggestions:** None

### Completion Checklist

- [x] All files generated successfully
- [x] No syntax errors in YAML
- [x] All paths are correct
- [x] Variables are consistent
- [x] Design requirements met
- [x] Best practices followed
- [x] User approved final workflow

---

## Workflow Creation Complete

**Status:** ✅ COMPLETE
**Completed:** 2025-12-04
**Author:** Nearformer

### Output Location

`bmad-custom/workflows/create-epics-and-stories-github/`

### How to Use

1. Set `GITHUB_TOKEN` environment variable
2. Have PRD.md and Architecture.md ready
3. Reference workflow: `@bmad-custom/workflows/create-epics-and-stories-github`

---

## Final Configuration

### Output Files to Generate

- **workflow.yaml**: bmm workflow configuration
- **instructions.md**: Step-by-step workflow instructions
- **epics-template.md**: Template for epic/story output
- **github-integration.md**: GitHub API integration guide

### Target Location

- **Folder**: `{project-root}/.bmad/bmm/workflows/3-solutioning/create-epics-and-stories-github`
- **Module**: bmm
- **Custom Location**: `{project-root}/bmad-custom/workflows/create-epics-and-stories-github`

### Final Checklist

- [ ] All requirements documented
- [ ] Workflow designed and approved
- [ ] Files generated successfully
- [ ] Workflow tested and validated

## Ready for Implementation

When you approve this plan, I'll generate all the workflow files in the specified location with the exact structure and content outlined above.

