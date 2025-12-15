---
name: 'step-03-generate-stories'
description: 'Generate detailed stories for all epics with acceptance criteria'

# Path Definitions
workflow_path: '{project-root}/bmad-custom/workflows/create-epics-and-stories-github'

# File References
thisStepFile: '{workflow_path}/steps/step-03-generate-stories.md'
nextStepFile: '{workflow_path}/steps/step-04-review.md'
workflowFile: '{workflow_path}/instructions.md'

# Task References
advancedElicitationTask: '{project-root}/.bmad/core/tasks/advanced-elicitation.xml'
partyModeWorkflow: '{project-root}/.bmad/core/workflows/party-mode/workflow.md'

# Template References
storyTemplate: '{workflow_path}/templates/story-template.md'
---

# Step 3: Generate Detailed Stories

## STEP GOAL:

Generate ALL detailed stories for ALL epics, including complete acceptance criteria, technical implementation notes, and UX specifications. Stories should be implementation-ready for a developer.

## MANDATORY EXECUTION RULES (READ FIRST):

### Universal Rules:

- 🛑 NEVER push to GitHub in this step
- 📖 CRITICAL: Read the complete step file before taking any action
- 🔄 CRITICAL: Generate ALL stories for ALL epics in this step
- 📋 YOU ARE A FACILITATOR, not a content generator

### Role Reinforcement:

- ✅ You are a product decomposition specialist
- ✅ We engage in collaborative dialogue, not command-response
- ✅ You bring expertise in Agile story writing with technical depth
- ✅ User brings their domain knowledge and context

### Step-Specific Rules:

- 🎯 Focus on comprehensive story generation
- 🚫 FORBIDDEN to push to GitHub in this step
- 💬 Stories must be detailed enough for implementation
- 📋 Each story must be completable in one focused session

## STORY GENERATION PRINCIPLES:

### Story Sizing

Each story MUST be completable by a single developer in one focused session:
- If too large, break into smaller stories
- Maintain user value in each story
- Clear acceptance criteria for each

### Technical Integration

Every story should incorporate context from:
- **PRD**: Business rules, user types, constraints
- **Architecture**: API endpoints, data models, patterns
- **UX Design** (if available): UI specs, interactions, validation

### Story Format

Use the User Story format with detailed acceptance criteria:

```
**Story {Epic}.{Story}: {Title}**

As a {user type},
I want {capability},
So that {benefit}.

**Acceptance Criteria:**
Given {context}
When {action}
Then {expected result}
And {additional expectations}

**Technical Notes:**
- {API endpoint from Architecture}
- {Data model considerations}
- {Security requirements}
- {Performance requirements}

**UX Notes:** (if applicable)
- {UI component reference}
- {Interaction pattern}
- {Validation rules}
- {Error states}

**Prerequisites:** Story {X.Y}
```

## SEQUENCE OF INSTRUCTIONS:

### 1. Initialize Story Generation

"**Generating Stories for All Epics**

Based on the approved epic structure, I'll now generate detailed stories for each epic. Each story will include:
- User story format (As a... I want... So that...)
- Detailed acceptance criteria (Given/When/Then)
- Technical implementation notes from Architecture
- UX specifications (if available)

Generating stories for {total_epic_count} epics..."

### 2. Generate Stories for Each Epic

For each epic in the approved structure:

"**Epic {N}: {Epic Title}**

{Epic Goal Statement}

**Stories:**"

Generate all stories for this epic:

**Story {N}.1: {Story Title}**
As a {user type},
I want {capability},
So that {benefit}.

**Acceptance Criteria:**
- Given {context}
- When {action}
- Then {expected result}

**Technical Notes:**
- {From Architecture}

**UX Notes:** {if applicable}

**Prerequisites:** {or "None - first story in epic"}

---

Continue for all stories in the epic, then move to next epic.

### 3. Present Complete Story Overview

After generating all stories, present summary:

"**Story Generation Complete**

| Epic | Epic Title | Stories | FR Coverage |
|------|------------|---------|-------------|
| 1 | {Title} | {count} | {FR list} |
| 2 | {Title} | {count} | {FR list} |
| ... | ... | ... | ... |

**Totals:**
- Total Epics: {count}
- Total Stories: {count}
- FRs Covered: {covered}/{total} (should be 100%)

**Story Details Generated:**
- Acceptance Criteria: ✅ All stories
- Technical Notes: ✅ All stories
- UX Notes: {count} stories (where applicable)
- Prerequisites: ✅ All mapped"

### 4. Validation Check

Verify completeness:

"**Validation Checklist:**

- [ ] All FRs mapped to at least one story
- [ ] All stories have acceptance criteria
- [ ] All stories have technical notes
- [ ] Story dependencies are logical
- [ ] Each story is sized for single-session completion
- [ ] No orphaned stories (all have epic assignment)

**Issues Found:** {list any issues or "None"}"

### 5. Present MENU OPTIONS

Display: **Select an Option:** [A] Advanced Elicitation [P] Party Mode [C] Continue

#### EXECUTION RULES:

- ALWAYS halt and wait for user input after presenting menu
- ONLY proceed to next step when user selects 'C'
- After other menu items execution, return to this menu
- User can chat or ask questions - always respond and then redisplay menu

#### Menu Handling Logic:

- IF A: Execute {advancedElicitationTask}
- IF P: Execute {partyModeWorkflow}
- IF C: Store all generated stories, then load, read entire file, then execute {nextStepFile}
- IF Any other comments or queries: help user respond then redisplay menu

## CRITICAL STEP COMPLETION NOTE

ONLY WHEN C is selected and stories are approved, will you then load, read entire file, then execute {nextStepFile} for final review before GitHub push.

---

## 🚨 SYSTEM SUCCESS/FAILURE METRICS

### ✅ SUCCESS:

- ALL stories generated for ALL epics
- Each story has complete acceptance criteria
- Technical notes included from Architecture
- UX notes included where applicable
- 100% FR coverage verified
- Stories sized appropriately
- Menu presented and user input handled correctly

### ❌ SYSTEM FAILURE:

- Generating stories one epic at a time with separate approvals
- Missing acceptance criteria
- Missing technical context
- Pushing to GitHub in this step
- Proceeding without complete story generation

**Master Rule:** Skipping steps, optimizing sequences, or not following exact instructions is FORBIDDEN and constitutes SYSTEM FAILURE.

