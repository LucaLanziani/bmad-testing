---
name: 'step-08-summary'
description: 'Present completion summary with GitHub URLs and cleanup options'

# Path Definitions
workflow_path: '{project-root}/bmad-custom/workflows/create-epics-and-stories-github'

# File References
thisStepFile: '{workflow_path}/steps/step-08-summary.md'
workflowFile: '{workflow_path}/instructions.md'
sidecarFile: '{output_folder}/.github-sync-state.yaml'
---

# Step 8: Completion Summary

## STEP GOAL:

Present a comprehensive summary of everything created on GitHub, provide direct links for easy access, and offer cleanup options for the sidecar file.

## MANDATORY EXECUTION RULES (READ FIRST):

### Universal Rules:

- 📖 CRITICAL: Read the complete step file before taking any action
- 🔄 CRITICAL: Provide clickable GitHub URLs
- 📋 This is the FINAL step - workflow completes here

### Role Reinforcement:

- ✅ You are a helpful completion assistant
- ✅ We celebrate success and provide useful next steps
- ✅ You ensure user has all information needed

### Step-Specific Rules:

- 🎯 Present complete summary with URLs
- 💬 Congratulate user on completion
- 📋 Offer sidecar file cleanup option
- 🎉 This is the workflow finale!

## SEQUENCE OF INSTRUCTIONS:

### 1. Completion Banner

"
🎉 **WORKFLOW COMPLETE!**

Your epics and stories have been successfully created on GitHub!
"

### 2. Repository Summary

"**Repository:** [{owner}/{repo}](https://github.com/{owner}/{repo})

**Direct Links:**
- 📋 [View all Issues](https://github.com/{owner}/{repo}/issues)
- 🎯 [View all Milestones](https://github.com/{owner}/{repo}/milestones)
- 🏷️ [View all Labels](https://github.com/{owner}/{repo}/labels)"

### 3. Created Items Summary

Read from sidecar file and present:

"**What Was Created:**

### Labels ({count})
| Label | Purpose |
|-------|---------|
| epic | Epic-level work items |
| story | User stories |
| priority:high | High priority items |
| priority:medium | Medium priority items |
| priority:low | Low priority items |

### Milestones ({count})
| # | Milestone | Stories | Link |
|---|-----------|---------|------|
| 1 | {Title} | {count} | [View](https://github.com/{owner}/{repo}/milestone/1) |
| 2 | {Title} | {count} | [View](https://github.com/{owner}/{repo}/milestone/2) |
| ... | ... | ... | ... |

### Issues ({count})
| Epic | Issue Range | Milestone |
|------|-------------|-----------|
| 1 | #{start}-#{end} | {Milestone Title} |
| 2 | #{start}-#{end} | {Milestone Title} |
| ... | ... | ... |

**Quick Links by Epic:**"

For each epic:
"- **{Epic Title}**: [View {count} issues](https://github.com/{owner}/{repo}/issues?q=is%3Aissue+milestone%3A%22{encoded_milestone_name}%22)"

### 4. Statistics

"**Statistics:**

| Metric | Count |
|--------|-------|
| Total Epics/Milestones | {count} |
| Total Stories/Issues | {count} |
| API Calls Made | ~{count} |
| Time Taken | {duration} |

**FR Coverage:** 100% - All {fr_count} functional requirements mapped to stories"

### 5. Next Steps

"**Recommended Next Steps:**

1. **Review on GitHub** - Check that all issues look correct
2. **Prioritize** - Use GitHub Projects to organize your sprint
3. **Assign** - Assign issues to team members
4. **Start Sprint** - Begin development with Story 1.1!

**Sprint Planning Tip:**
Your issues are already organized by milestone (epic). Consider:
- Sprint 1: Epic 1 (Foundation)
- Sprint 2: Epic 2 stories
- etc."

### 6. Sidecar File Options

"**State File Options:**

The workflow state file (`.github-sync-state.yaml`) contains a record of everything created.

**Options:**
[K] Keep the state file (useful for reference)
[D] Delete the state file (clean up)

What would you like to do?"

Handle response:
- IF K: "✅ State file kept at: {sidecarFile}"
- IF D: Delete the sidecar file, "✅ State file deleted"

### 7. Final Message

"
---

**Thank you for using the GitHub Epic & Story Creator!**

Your project is now set up on GitHub and ready for development.

🚀 **Happy building, {user_name}!**

---

*Workflow: create-epics-and-stories-github*
*Completed: {timestamp}*
"

---

## 🚨 SYSTEM SUCCESS/FAILURE METRICS

### ✅ SUCCESS:

- Complete summary presented with all URLs
- Direct GitHub links provided
- Statistics displayed
- Next steps recommended
- Sidecar cleanup option offered
- User congratulated on completion

### ❌ SYSTEM FAILURE:

- Not providing GitHub URLs
- Missing summary information
- Not offering cleanup option
- Abrupt ending without celebration

**Master Rule:** This is the final step - ensure user leaves with all information needed and a positive experience.

