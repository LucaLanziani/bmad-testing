---
name: 'step-05-push-labels'
description: 'Create default labels on GitHub repository'

# Path Definitions
workflow_path: '{project-root}/bmad-custom/workflows/create-epics-and-stories-github'

# File References
thisStepFile: '{workflow_path}/steps/step-05-push-labels.md'
nextStepFile: '{workflow_path}/steps/step-06-push-milestones.md'
workflowFile: '{workflow_path}/instructions.md'
sidecarFile: '{output_folder}/.github-sync-state.yaml'
---

# Step 5: Push Labels to GitHub

## STEP GOAL:

Create the default labels on the GitHub repository that will be used to categorize epics and stories.

## MANDATORY EXECUTION RULES (READ FIRST):

### Universal Rules:

- 🛑 ALWAYS update sidecar file after each successful operation
- 📖 CRITICAL: Read the complete step file before taking any action
- 🔄 CRITICAL: Check if labels already exist before creating
- 📋 Handle API errors gracefully

### Role Reinforcement:

- ✅ You are a GitHub integration engineer
- ✅ We show progress clearly to user
- ✅ You handle errors gracefully with retry options

### Step-Specific Rules:

- 🎯 Create only the default labels needed
- 🚫 FORBIDDEN to create milestones or issues in this step
- 💬 Show progress for each label
- 📋 Skip labels that already exist

## EXECUTION PROTOCOLS:

- 🎯 Create labels via GitHub REST API
- 💾 Update sidecar file after each successful creation
- 📖 Show progress to user
- 🚫 Handle errors with retry/skip options

## SEQUENCE OF INSTRUCTIONS:

### 1. Initialize Label Creation

"**Creating Labels on GitHub**

Repository: {owner}/{repo}

I'll now create the default labels for organizing epics and stories..."

### 2. Initialize Sidecar File

Create or update sidecar file with workflow state:

```yaml
workflow_run_id: "{generated_uuid}"
repository: "{owner}/{repo}"
created_at: "{timestamp}"
status: "in_progress"
checkpoint:
  step: "push-labels"
  message: "Starting label creation"
created:
  labels: []
  milestones: []
  issues: []
```

### 3. Create Each Label

For each label in the default labels list:

**Label: {name}**

First, check if label exists:
GET `https://api.github.com/repos/{owner}/{repo}/labels/{name}`

If label exists (200 response):
"⏭️ Label '{name}' already exists - skipping"

If label doesn't exist (404 response):
Create the label:
POST `https://api.github.com/repos/{owner}/{repo}/labels`
```json
{
  "name": "{name}",
  "color": "{color}",
  "description": "{description}"
}
```

If successful (201 response):
"✅ Created label: {name}"
Update sidecar file: add label to `created.labels` array

If error:
"❌ Failed to create label '{name}': {error_message}"
Ask user: "[R] Retry | [S] Skip | [A] Abort"

### 4. Labels Progress Display

Show progress after each label:

```
📋 Label Progress:
[✅] epic
[✅] story
[✅] priority:high
[⏳] priority:medium
[ ] priority:low

Progress: 3/5 complete
```

### 5. Labels Summary

After all labels processed:

"**Label Creation Complete**

| Label | Status |
|-------|--------|
| epic | ✅ Created / ⏭️ Existed |
| story | ✅ Created / ⏭️ Existed |
| priority:high | ✅ Created / ⏭️ Existed |
| priority:medium | ✅ Created / ⏭️ Existed |
| priority:low | ✅ Created / ⏭️ Existed |

**Summary:**
- Created: {count}
- Already existed: {count}
- Failed: {count}

Sidecar file updated with label state."

### 6. Update Sidecar Checkpoint

Update sidecar file:
```yaml
checkpoint:
  step: "push-labels"
  message: "Labels complete. Ready for milestones."
```

### 7. Confirm and Proceed

Display: "**Labels complete. Ready to create milestones?** [Y] Yes, continue | [N] No, stop here"

#### Menu Handling Logic:

- IF Y: Load, read entire file, then execute {nextStepFile}
- IF N: Save current state, inform user they can resume later
- IF Any questions: respond and redisplay options

## CRITICAL STEP COMPLETION NOTE

ONLY WHEN user confirms continuation will you load, read entire file, then execute {nextStepFile} to begin milestone creation.

---

## 🚨 SYSTEM SUCCESS/FAILURE METRICS

### ✅ SUCCESS:

- All labels created or confirmed existing
- Sidecar file updated with label state
- Progress displayed clearly
- Errors handled gracefully
- User confirmed continuation

### ❌ SYSTEM FAILURE:

- Not checking if labels exist before creating
- Not updating sidecar file
- Not handling API errors
- Creating milestones or issues in this step
- Proceeding without user confirmation

**Master Rule:** Skipping steps, optimizing sequences, or not following exact instructions is FORBIDDEN and constitutes SYSTEM FAILURE.

