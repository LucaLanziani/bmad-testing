---
name: 'step-06-push-milestones'
description: 'Create GitHub milestones for each epic'

# Path Definitions
workflow_path: '{project-root}/bmad-custom/workflows/create-epics-and-stories-github'

# File References
thisStepFile: '{workflow_path}/steps/step-06-push-milestones.md'
nextStepFile: '{workflow_path}/steps/step-07-push-issues.md'
workflowFile: '{workflow_path}/instructions.md'
sidecarFile: '{output_folder}/.github-sync-state.yaml'
---

# Step 6: Push Milestones to GitHub

## STEP GOAL:

Create GitHub milestones for each epic. Each epic becomes a milestone that stories (issues) will be assigned to.

## MANDATORY EXECUTION RULES (READ FIRST):

### Universal Rules:

- 🛑 ALWAYS update sidecar file after each successful milestone
- 📖 CRITICAL: Read the complete step file before taking any action
- 🔄 CRITICAL: Store milestone IDs for issue assignment
- 📋 Handle API errors gracefully

### Role Reinforcement:

- ✅ You are a GitHub integration engineer
- ✅ We show progress clearly to user
- ✅ You handle errors gracefully with retry options

### Step-Specific Rules:

- 🎯 Create one milestone per epic
- 🚫 FORBIDDEN to create issues in this step
- 💬 Show progress and confirm after each milestone
- 📋 Store milestone IDs for next step

## EXECUTION PROTOCOLS:

- 🎯 Create milestones via GitHub REST API
- 💾 Update sidecar file after each milestone
- 📖 Show progress and request confirmation
- 🚫 Handle errors with retry/skip options

## SEQUENCE OF INSTRUCTIONS:

### 1. Initialize Milestone Creation

"**Creating Milestones on GitHub**

Repository: {owner}/{repo}

I'll create {milestone_count} milestones (one for each epic)..."

### 2. Update Sidecar Checkpoint

Update sidecar file:
```yaml
checkpoint:
  step: "push-milestones"
  epic_index: 0
  message: "Starting milestone creation"
```

### 3. Create Each Milestone

For each epic:

"**Creating Milestone {N}/{total}: {Epic Title}**"

POST `https://api.github.com/repos/{owner}/{repo}/milestones`
```json
{
  "title": "{Epic Title}",
  "description": "{Epic Goal}\n\nFRs Covered: {FR list}\n\nStories: {story_count}",
  "state": "open"
}
```

If successful (201 response):
"✅ Created milestone: {Epic Title} (ID: {milestone_number})"

Update sidecar file:
```yaml
created:
  milestones:
    - epic_id: {N}
      github_milestone_number: {number from response}
      title: "{Epic Title}"
      state: "created"
checkpoint:
  epic_index: {N}
  message: "Completed milestone {N}: {Epic Title}"
```

If error:
"❌ Failed to create milestone '{Epic Title}': {error_message}"
Ask user: "[R] Retry | [S] Skip | [A] Abort"

### 4. Progress After Each Milestone

Show progress after each milestone:

```
🎯 Milestone Progress:

[✅] Epic 1: Foundation & Core Setup (ID: 1)
[✅] Epic 2: User Authentication (ID: 2)
[⏳] Epic 3: Content Management
[ ] Epic 4: Discovery & Search

Progress: 2/4 complete
```

### 5. Confirmation Between Milestones

After each milestone, ask:
"**Continue to next milestone?** [Y] Yes | [N] Pause here"

If N: Save state and inform user they can resume later.

### 6. Milestones Summary

After all milestones created:

"**Milestone Creation Complete**

| Epic | Milestone Title | GitHub ID | Status |
|------|-----------------|-----------|--------|
| 1 | {Title} | #{id} | ✅ Created |
| 2 | {Title} | #{id} | ✅ Created |
| 3 | {Title} | #{id} | ✅ Created |
| ... | ... | ... | ... |

**Summary:**
- Total Milestones: {count}
- Successfully Created: {count}
- Failed: {count}

**Milestone IDs stored for issue assignment.**"

### 7. Update Final Checkpoint

Update sidecar file:
```yaml
checkpoint:
  step: "push-milestones"
  message: "All milestones complete. Ready for issues."
```

### 8. Confirm and Proceed to Issues

Display: "**Milestones complete. Ready to create {issue_count} issues?** [Y] Yes, continue | [N] No, stop here"

#### Menu Handling Logic:

- IF Y: Load, read entire file, then execute {nextStepFile}
- IF N: Save current state, inform user they can resume later
- IF Any questions: respond and redisplay options

## CRITICAL STEP COMPLETION NOTE

ONLY WHEN user confirms continuation will you load, read entire file, then execute {nextStepFile} to begin issue creation.

---

## 🚨 SYSTEM SUCCESS/FAILURE METRICS

### ✅ SUCCESS:

- All milestones created successfully
- Milestone IDs stored in sidecar file
- Progress displayed clearly with confirmations
- Errors handled gracefully
- User confirmed continuation to issues

### ❌ SYSTEM FAILURE:

- Not storing milestone IDs (needed for issue assignment)
- Not updating sidecar file after each milestone
- Not confirming between milestones
- Creating issues in this step
- Proceeding without user confirmation

**Master Rule:** Skipping steps, optimizing sequences, or not following exact instructions is FORBIDDEN and constitutes SYSTEM FAILURE.

