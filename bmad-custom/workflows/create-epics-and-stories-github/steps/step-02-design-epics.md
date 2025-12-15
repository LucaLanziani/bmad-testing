---
name: 'step-02-design-epics'
description: 'Design the epic structure based on extracted requirements'

# Path Definitions
workflow_path: '{project-root}/bmad-custom/workflows/create-epics-and-stories-github'

# File References
thisStepFile: '{workflow_path}/steps/step-02-design-epics.md'
nextStepFile: '{workflow_path}/steps/step-03-generate-stories.md'
workflowFile: '{workflow_path}/instructions.md'

# Task References
advancedElicitationTask: '{project-root}/.bmad/core/tasks/advanced-elicitation.xml'
partyModeWorkflow: '{project-root}/.bmad/core/workflows/party-mode/workflow.md'

# Template References
epicTemplate: '{workflow_path}/templates/epic-template.md'
---

# Step 2: Design Epic Structure

## STEP GOAL:

Design the epic structure that organizes related functional requirements into user-value-delivering epics. Each epic should represent a coherent set of functionality that users can benefit from.

## MANDATORY EXECUTION RULES (READ FIRST):

### Universal Rules:

- 🛑 NEVER generate detailed stories in this step
- 📖 CRITICAL: Read the complete step file before taking any action
- 🔄 CRITICAL: Epics must deliver USER VALUE, not just technical capability
- 📋 YOU ARE A FACILITATOR, not a content generator

### Role Reinforcement:

- ✅ You are a product decomposition specialist
- ✅ We engage in collaborative dialogue, not command-response
- ✅ You bring expertise in Agile epic design
- ✅ User brings their domain knowledge and priorities

### Step-Specific Rules:

- 🎯 Focus ONLY on epic-level structure
- 🚫 FORBIDDEN to create detailed story content in this step
- 💬 Design epics collaboratively with user input
- 📋 Each epic must map to specific FRs

## EXECUTION PROTOCOLS:

- 🎯 Design epic structure based on requirements
- 💾 Map FRs to proposed epics
- 📖 Get user approval on epic structure
- 🚫 FORBIDDEN to proceed without epic approval

## EPIC DESIGN PRINCIPLES:

### User-Value First

Each epic MUST deliver USER VALUE, not just technical capability:

✅ **CORRECT - User Value:**
- Epic: "User Authentication & Profile Management" (users can register, login, manage profiles)
- Epic: "Content Creation & Publishing" (users can create and share content)

❌ **WRONG - Technical Layers:**
- Epic: "Database Schema Setup" (no direct user value)
- Epic: "API Endpoints" (no direct user value)

### Foundation Epic Exception

Epic 1 may include technical foundation work that enables all subsequent user-facing features:
- Project setup and build system
- Core infrastructure and deployment
- Database schema foundation
- Authentication framework
- API framework setup

## SEQUENCE OF INSTRUCTIONS:

### 1. Analyze Requirements for Epic Grouping

Review the FR inventory from Step 1 and identify natural groupings:

"Based on your {total_fr_count} functional requirements, I'll propose epic groupings that deliver incremental user value..."

### 2. Propose Epic Structure

Present proposed epic structure:

"**Proposed Epic Structure:**

**Epic 1: Foundation & Core Setup**
- Goal: Establish technical foundation enabling all subsequent features
- FRs Covered: Infrastructure, deployment, core services
- User Value: Platform is operational and deployable

**Epic 2: {User-Value Title}**
- Goal: {What users can accomplish}
- FRs Covered: {FR list}
- User Value: {Specific benefit to users}
- Dependencies: Epic 1

**Epic 3: {User-Value Title}**
- Goal: {What users can accomplish}
- FRs Covered: {FR list}
- User Value: {Specific benefit to users}
- Dependencies: {list}

... continue for all epics ..."

### 3. FR Coverage Matrix

Present coverage verification:

"**FR Coverage Matrix:**

| FR | Description | Epic |
|----|-------------|------|
| FR1 | {desc} | Epic 2 |
| FR2 | {desc} | Epic 2 |
| FR3 | {desc} | Epic 3 |
| ... | ... | ... |

**Coverage Status:** {covered_count}/{total_count} FRs mapped to epics"

Ensure ALL FRs are covered.

### 4. Epic Summary Table

"**Epic Summary:**

| # | Epic Title | FR Count | Dependencies |
|---|------------|----------|--------------|
| 1 | Foundation & Core Setup | {n} | None |
| 2 | {Title} | {n} | Epic 1 |
| 3 | {Title} | {n} | Epic 1, 2 |
| ... | ... | ... | ... |

**Total Epics:** {count}
**All FRs Covered:** ✅ Yes / ❌ No"

### 5. Collaborative Refinement

Ask for user feedback:

"Does this epic structure make sense for your project?

**Consider:**
- Are the epic boundaries logical?
- Does each epic deliver clear user value?
- Are the dependencies correct?
- Should any epics be split or combined?

Please share your thoughts or adjustments."

### 6. Present MENU OPTIONS

Display: **Select an Option:** [A] Advanced Elicitation [P] Party Mode [C] Continue

#### EXECUTION RULES:

- ALWAYS halt and wait for user input after presenting menu
- ONLY proceed to next step when user selects 'C'
- After other menu items execution, return to this menu
- User can chat or ask questions - always respond and then redisplay menu

#### Menu Handling Logic:

- IF A: Execute {advancedElicitationTask}
- IF P: Execute {partyModeWorkflow}
- IF C: Store approved epic structure, then load, read entire file, then execute {nextStepFile}
- IF Any other comments or queries: help user respond then redisplay menu

## CRITICAL STEP COMPLETION NOTE

ONLY WHEN C is selected and epic structure is approved, will you then load, read entire file, then execute {nextStepFile} to begin detailed story generation.

---

## 🚨 SYSTEM SUCCESS/FAILURE METRICS

### ✅ SUCCESS:

- Epic structure designed with user-value focus
- ALL FRs mapped to epics (100% coverage)
- Dependencies clearly identified
- User approved the epic structure
- Menu presented and user input handled correctly

### ❌ SYSTEM FAILURE:

- Creating technical-layer epics (not user-value)
- Missing FR coverage
- Creating detailed stories in this step
- Proceeding without user approval
- Not presenting FR coverage matrix

**Master Rule:** Skipping steps, optimizing sequences, or not following exact instructions is FORBIDDEN and constitutes SYSTEM FAILURE.

