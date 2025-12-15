# Epic Template

Use this template format when presenting epic structure to users.

## Epic Display Format

```markdown
## Epic {N}: {Epic Title}

**Goal:** {What users can accomplish after this epic is complete}

**User Value:** {Specific benefit to end users}

**FRs Covered:**
- FR{X}: {Description}
- FR{Y}: {Description}
- ...

**Dependencies:** {List of prerequisite epics, or "None" for first epic}

**Story Count:** {Number of stories in this epic}

---
```

## GitHub Milestone Format

When creating the GitHub milestone, use this description format:

```markdown
## Goal
{Epic goal statement}

## User Value
{What users can do after this epic is complete}

## Functional Requirements Covered
- FR{X}: {Brief description}
- FR{Y}: {Brief description}

## Stories in this Epic
{count} stories will be created as issues assigned to this milestone.

---
*Created by: create-epics-and-stories-github workflow*
```

## Example Epic

```markdown
## Epic 2: User Authentication & Profile Management

**Goal:** Enable users to securely create accounts, authenticate, and manage their profiles.

**User Value:** Users can register, login, reset passwords, and maintain their profile information.

**FRs Covered:**
- FR3: User Registration
- FR4: User Authentication
- FR5: Password Recovery
- FR6: Profile Management

**Dependencies:** Epic 1 (Foundation)

**Story Count:** 6 stories
```

