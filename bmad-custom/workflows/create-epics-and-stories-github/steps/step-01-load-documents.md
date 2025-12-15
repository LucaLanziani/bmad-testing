---
name: 'step-01-load-documents'
description: 'Load and analyze PRD, Architecture, and UX documents to extract requirements'

# Path Definitions
workflow_path: '{project-root}/bmad-custom/workflows/create-epics-and-stories-github'

# File References
thisStepFile: '{workflow_path}/steps/step-01-load-documents.md'
nextStepFile: '{workflow_path}/steps/step-02-design-epics.md'
workflowFile: '{workflow_path}/instructions.md'

# Task References
advancedElicitationTask: '{project-root}/.bmad/core/tasks/advanced-elicitation.xml'
partyModeWorkflow: '{project-root}/.bmad/core/workflows/party-mode/workflow.md'
---

# Step 1: Load and Analyze Documents

## STEP GOAL:

Load all input documents (PRD, Architecture, UX Design) and extract the requirements, technical context, and UX patterns that will inform epic and story creation.

## MANDATORY EXECUTION RULES (READ FIRST):

### Universal Rules:

- 🛑 NEVER generate epics or stories in this step
- 📖 CRITICAL: Read the complete step file before taking any action
- 🔄 CRITICAL: Load ALL documents before analysis
- 📋 YOU ARE A FACILITATOR, not a content generator

### Role Reinforcement:

- ✅ You are a product decomposition specialist
- ✅ We engage in collaborative dialogue, not command-response
- ✅ You bring expertise in requirements analysis
- ✅ User brings their domain knowledge and project context

### Step-Specific Rules:

- 🎯 Focus ONLY on loading and extracting information
- 🚫 FORBIDDEN to create epics or stories in this step
- 💬 Present extracted information clearly for confirmation
- 📋 Create comprehensive FR inventory

## EXECUTION PROTOCOLS:

- 🎯 Load all documents completely
- 💾 Extract and organize key information
- 📖 Present summary for user confirmation
- 🚫 FORBIDDEN to proceed without user confirmation

## SEQUENCE OF INSTRUCTIONS:

### 1. Load PRD Document

Load the PRD document and extract:

**Functional Requirements (FR) Inventory:**
- List ALL functional requirements (FR1, FR2, FR3...)
- Include requirement description for each
- Note priority levels if specified

**Non-Functional Requirements:**
- Performance requirements
- Security requirements
- Compliance requirements

**Project Scope:**
- MVP scope boundaries
- Growth features
- Vision/future features

**User Types:**
- Primary users and their goals
- Secondary users if any

### 2. Load Architecture Document

Load the Architecture document and extract:

**Technology Stack:**
- Languages, frameworks, databases
- Third-party services and APIs

**API Design:**
- Endpoint patterns
- Authentication approach
- Data models and schemas

**Technical Decisions:**
- Key architectural patterns
- Integration points
- Security implementation

**Infrastructure:**
- Deployment approach
- Scaling considerations

### 3. Load UX Design Document (If Available)

If UX Design document exists, load and extract:

**User Flows:**
- Primary user journeys
- Interaction patterns

**UI Components:**
- Screen layouts
- Component specifications

**Design Patterns:**
- Form validation rules
- Error handling UX
- Loading states
- Responsive design requirements

### 4. Present Extraction Summary

"**Document Analysis Complete**

**From PRD:**
- Functional Requirements: {count} identified
- Non-Functional Requirements: {count} identified
- User Types: {list}
- Scope: MVP | Growth | Vision phases defined

**From Architecture:**
- Tech Stack: {summary}
- API Endpoints: {count} defined
- Key Patterns: {list}

**From UX Design:** {if available}
- User Flows: {count} documented
- UI Components: {count} specified

**FR Inventory Preview:**

| ID | Description | Priority |
|----|-------------|----------|
| FR1 | {description} | {priority} |
| FR2 | {description} | {priority} |
| ... | ... | ... |

All {total_fr_count} functional requirements will be mapped to stories."

### 5. Confirm Document Analysis

Ask user to confirm the extracted information is complete:

"Does this summary accurately capture your project requirements? 

If anything is missing or incorrect, please let me know and I'll adjust the analysis."

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
- IF C: Store extracted data in session, then load, read entire file, then execute {nextStepFile}
- IF Any other comments or queries: help user respond then redisplay menu

## CRITICAL STEP COMPLETION NOTE

ONLY WHEN C is selected and document analysis is confirmed, will you then load, read entire file, then execute {nextStepFile} to begin epic structure design.

---

## 🚨 SYSTEM SUCCESS/FAILURE METRICS

### ✅ SUCCESS:

- All available documents loaded completely
- FR inventory created with all requirements
- Technical context extracted from Architecture
- UX patterns extracted (if available)
- Summary presented and confirmed by user
- Menu presented and user input handled correctly

### ❌ SYSTEM FAILURE:

- Creating epics or stories in this step
- Not loading all available documents
- Skipping FR inventory creation
- Proceeding without user confirmation
- Not presenting extraction summary

**Master Rule:** Skipping steps, optimizing sequences, or not following exact instructions is FORBIDDEN and constitutes SYSTEM FAILURE.

