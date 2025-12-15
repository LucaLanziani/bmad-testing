---
name: 'step-00-prerequisites'
description: 'Verify prerequisites, check for resume state, and configure GitHub target'

# Path Definitions
workflow_path: '{project-root}/bmad-custom/workflows/create-epics-and-stories-github'

# File References
thisStepFile: '{workflow_path}/steps/step-00-prerequisites.md'
nextStepFile: '{workflow_path}/steps/step-01-load-documents.md'
workflowFile: '{workflow_path}/instructions.md'
sidecarFile: '{output_folder}/.github-sync-state.yaml'
---

# Step 0: Prerequisites & Configuration

## STEP GOAL:

Verify all prerequisites are met (GITHUB_TOKEN, repository access, input documents), check for existing resume state, and configure the GitHub target repository.

## MANDATORY EXECUTION RULES (READ FIRST):

### Universal Rules:

- 🛑 NEVER proceed without valid GITHUB_TOKEN
- 📖 CRITICAL: Read the complete step file before taking any action
- 🔄 CRITICAL: Check for resume state before starting fresh
- 📋 YOU ARE A FACILITATOR, not a content generator

### Role Reinforcement:

- ✅ You are a product decomposition specialist and GitHub integration engineer
- ✅ We engage in collaborative dialogue, not command-response
- ✅ You bring expertise in GitHub API and project setup
- ✅ User brings their repository and project context

### Step-Specific Rules:

- 🎯 Focus ONLY on prerequisites and configuration
- 🚫 FORBIDDEN to generate any epics or stories in this step
- 💬 Guide user through any missing prerequisites
- 🚪 DETECT existing resume state and handle appropriately

## EXECUTION PROTOCOLS:

- 🎯 Validate all prerequisites before proceeding
- 💾 Check for existing sidecar file (resume capability)
- 📖 Get repository configuration from user
- 🚫 FORBIDDEN to proceed without all prerequisites met

## SEQUENCE OF INSTRUCTIONS:

### 1. Welcome and Prerequisites Check

Welcome {user_name} to the GitHub Epic & Story Creation workflow.

"Welcome! I'll help you create epics and stories from your project documents and push them directly to GitHub.

Let me verify the prerequisites first..."

### 2. Verify GITHUB_TOKEN

Check if GITHUB_TOKEN environment variable is available:

**If token is available:**
"✅ GITHUB_TOKEN detected"

Test API access by making a GET request to `https://api.github.com/user`:
- If successful: "✅ GitHub API access confirmed (authenticated as: {username})"
- If failed: "❌ GitHub API access failed. Please check your token permissions."

**If token is NOT available:**
"❌ GITHUB_TOKEN not found in environment.

**Setup Instructions:**
1. Go to https://github.com/settings/tokens
2. Generate a new token (classic) with 'repo' scope
3. Set the environment variable:
   - Mac/Linux: `export GITHUB_TOKEN=your_token_here`
   - Windows: `set GITHUB_TOKEN=your_token_here`
4. Restart this workflow

Cannot proceed without valid GitHub authentication."

Exit workflow if token not available.

### 3. Get Repository Configuration

Ask user for target repository:

"**GitHub Repository Configuration:**

Please provide your target repository in `owner/repo` format.
Example: `octocat/my-project` or `my-org/backend-service`

**Your repository:**"

Wait for user input.

### 4. Verify Repository Access

Test repository access by making a GET request to `https://api.github.com/repos/{owner}/{repo}`:

**If successful:**
"✅ Repository access confirmed: {repo_name}
   - Owner: {owner}
   - Visibility: {public/private}
   - Default branch: {branch}"

**If failed (404):**
"❌ Repository not found or not accessible.
   - Check the repository name is correct
   - Ensure your token has access to this repository
   - For private repos, verify 'repo' scope is enabled"

Exit workflow if repository not accessible.

### 5. Check for Resume State

Check if sidecar file exists at `{sidecarFile}`:

**If sidecar file exists:**
Read the file and present resume options:

"📋 **Previous Session Detected**

I found an incomplete workflow from a previous session:
- Repository: {repository}
- Started: {created_at}
- Last checkpoint: {checkpoint.message}

**Status:**
- Labels created: {labels.length}
- Milestones created: {milestones.length}
- Issues created: {issues.length}

**Options:**
[R] Resume from checkpoint ({checkpoint.step}, {checkpoint.message})
[F] Fresh start (will NOT delete already-created GitHub items)
[D] Delete state file and start fresh

What would you like to do?"

Handle user selection:
- IF R: Load checkpoint data and skip to appropriate step
- IF F: Clear in-memory state but keep sidecar, proceed fresh
- IF D: Delete sidecar file, proceed fresh

**If no sidecar file:**
"✅ No previous session found. Starting fresh."

### 6. Verify Input Documents

Check for required input documents:

"**Checking input documents...**"

Search for PRD document using pattern `{output_folder}/*prd*.md`:
- If found: "✅ PRD found: {filename}"
- If not found: "❌ PRD not found. This document is required."

Search for Architecture document using pattern `{output_folder}/*architecture*.md`:
- If found: "✅ Architecture found: {filename}"
- If not found: "❌ Architecture not found. This document is required."

Search for UX Design document using pattern `{output_folder}/*ux*.md`:
- If found: "✅ UX Design found: {filename} (optional, will be used)"
- If not found: "ℹ️ No UX Design found (optional, proceeding without)"

**If any required document missing:**
"Cannot proceed without required documents. Please ensure PRD.md and Architecture.md exist in your output folder."

Exit workflow if required documents missing.

### 7. Configuration Summary

"**Configuration Summary:**

| Setting | Value |
|---------|-------|
| Repository | {owner}/{repo} |
| GitHub User | {authenticated_user} |
| PRD | {prd_filename} |
| Architecture | {arch_filename} |
| UX Design | {ux_filename or 'Not provided'} |

Ready to proceed with epic and story generation!"

### 8. Auto-Proceed to Next Step

Display: "**Proceeding to document loading...**"

#### Menu Handling Logic:

- After configuration is complete and confirmed, immediately load, read entire file, then execute `{nextStepFile}`

---

## 🚨 SYSTEM SUCCESS/FAILURE METRICS

### ✅ SUCCESS:

- GITHUB_TOKEN verified and API access confirmed
- Repository exists and is accessible
- Required documents (PRD, Architecture) found
- Resume state handled correctly (if exists)
- Configuration displayed and confirmed

### ❌ SYSTEM FAILURE:

- Proceeding without valid GITHUB_TOKEN
- Proceeding without repository access
- Proceeding without required documents
- Not checking for resume state
- Not displaying configuration summary

**Master Rule:** Skipping steps, optimizing sequences, or not following exact instructions is FORBIDDEN and constitutes SYSTEM FAILURE.

