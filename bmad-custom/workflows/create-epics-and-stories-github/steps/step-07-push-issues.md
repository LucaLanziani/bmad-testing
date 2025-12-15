---
name: 'step-07-push-issues'
description: 'Create GitHub issues for each story and assign to milestones'

# Path Definitions
workflow_path: '{project-root}/bmad-custom/workflows/create-epics-and-stories-github'

# File References
thisStepFile: '{workflow_path}/steps/step-07-push-issues.md'
nextStepFile: '{workflow_path}/steps/step-08-summary.md'
workflowFile: '{workflow_path}/instructions.md'
sidecarFile: '{output_folder}/.github-sync-state.yaml'
---

# Step 7: Push Issues to GitHub

## STEP GOAL:

Create GitHub issues for each story, assign them to the appropriate milestone (epic), and apply relevant labels.

## MANDATORY EXECUTION RULES (READ FIRST):

### Universal Rules:

- 🛑 ALWAYS update sidecar file after each successful issue
- 📖 CRITICAL: Read the complete step file before taking any action
- 🔄 CRITICAL: Assign issues to correct milestones using stored IDs
- 📋 Handle API errors gracefully with exact checkpoint

### Role Reinforcement:

- ✅ You are a GitHub integration engineer
- ✅ We show detailed progress to user
- ✅ You handle errors gracefully with exact resume capability

### Step-Specific Rules:

- 🎯 Create one issue per story
- 💬 Show progress and confirm after each epic's issues
- 📋 Apply labels: "story" + priority label
- 🔄 Enable exact checkpoint resume (Epic X, Story Y)

## EXECUTION PROTOCOLS:

- 🎯 Create issues via GitHub REST API
- 💾 Update sidecar file after EACH issue
- 📖 Show progress with batch confirmations
- 🚫 Handle errors with retry/skip options

## SEQUENCE OF INSTRUCTIONS:

### 1. Initialize Issue Creation

"**Creating Issues on GitHub**

Repository: {owner}/{repo}

I'll create {total_issue_count} issues across {epic_count} epics..."

### 2. Update Sidecar Checkpoint

Update sidecar file:
```yaml
checkpoint:
  step: "push-issues"
  epic_index: 0
  story_index: 0
  message: "Starting issue creation"
```

### 3. Process Each Epic's Stories

For each epic:

"**Epic {N}: {Epic Title}**
Creating {story_count} issues for this epic..."

Get milestone number from sidecar: `created.milestones[epic_id={N}].github_milestone_number`

### 4. Create Each Issue

For each story in the epic:

POST `https://api.github.com/repos/{owner}/{repo}/issues`
```json
{
  "title": "Story {Epic}.{Story}: {Story Title}",
  "body": "{Full story content with acceptance criteria, technical notes, UX notes}",
  "milestone": {milestone_number},
  "labels": ["story", "{priority_label}"]
}
```

If successful (201 response):
"  ✅ #{issue_number}: {Story Title}"

Update sidecar file:
```yaml
created:
  issues:
    - story_id: "{Epic}.{Story}"
      github_issue_number: {number}
      milestone_id: {milestone_number}
      title: "{Story Title}"
      state: "created"
checkpoint:
  epic_index: {Epic}
  story_index: {Story}
  message: "Completed Epic {Epic}, Story {Story}"
```

If error:
"  ❌ Failed: {Story Title} - {error_message}"
Ask user: "[R] Retry | [S] Skip this story | [A] Abort"

### 5. Progress Display Per Epic

After each epic's issues:

```
📝 Issue Progress - Epic {N}: {Title}

[✅] #{101}: Story 1.1 - User Registration
[✅] #{102}: Story 1.2 - Email Verification
[✅] #{103}: Story 1.3 - Password Reset
[✅] #{104}: Story 1.4 - Profile Setup

Epic {N} Complete: 4/4 issues created
```

### 6. Confirmation After Each Epic

After completing each epic's issues:

"**Epic {N} issues complete ({count} created).**

Overall Progress: {completed_epics}/{total_epics} epics, {completed_issues}/{total_issues} issues

**Continue to Epic {N+1}?** [Y] Yes | [N] Pause here"

If N: Save exact checkpoint state for resume.

### 7. Handle Resume State

If resuming from checkpoint:
- Skip already-created issues (check sidecar `created.issues`)
- Start from exact checkpoint (epic_index, story_index)
- Inform user: "Resuming from Epic {X}, Story {Y}. Skipping {count} already-created issues."

### 8. Final Issues Summary

After all issues created:

"**Issue Creation Complete**

| Epic | Milestone | Issues Created | Issue Numbers |
|------|-----------|----------------|---------------|
| 1 | {Title} | {count} | #{start}-#{end} |
| 2 | {Title} | {count} | #{start}-#{end} |
| ... | ... | ... | ... |

**Summary:**
- Total Issues: {count}
- Successfully Created: {count}
- Skipped (errors): {count}
- Skipped (already existed): {count}

All issues assigned to milestones with appropriate labels."

### 9. Update Final Checkpoint

Update sidecar file:
```yaml
status: "completed"
checkpoint:
  step: "complete"
  message: "All issues created successfully"
completed_at: "{timestamp}"
```

### 10. Proceed to Summary

Display: "**All issues created! Proceeding to summary...**"

Automatically load, read entire file, then execute {nextStepFile}

---

## 🚨 SYSTEM SUCCESS/FAILURE METRICS

### ✅ SUCCESS:

- All issues created with correct milestone assignment
- Labels applied correctly
- Sidecar file updated after EACH issue
- Exact checkpoint enables resume at any point
- Progress displayed with epic-level confirmations
- Errors handled gracefully

### ❌ SYSTEM FAILURE:

- Not assigning issues to milestones
- Not applying labels
- Not updating sidecar after each issue
- Not enabling exact checkpoint resume
- Proceeding without confirmations between epics

**Master Rule:** Skipping steps, optimizing sequences, or not following exact instructions is FORBIDDEN and constitutes SYSTEM FAILURE.

