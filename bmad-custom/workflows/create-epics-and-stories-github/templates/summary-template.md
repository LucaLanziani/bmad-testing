# Summary Template

Use this template format for the final completion summary.

## Completion Summary Format

```markdown
# 🎉 GitHub Epic & Story Creation Complete!

## Repository
**{owner}/{repo}**

### Quick Links
- 📋 [All Issues](https://github.com/{owner}/{repo}/issues)
- 🎯 [All Milestones](https://github.com/{owner}/{repo}/milestones)
- 🏷️ [All Labels](https://github.com/{owner}/{repo}/labels)

---

## Created Items Summary

### Labels ({count})
| Label | Color | Purpose |
|-------|-------|---------|
| epic | 🟣 #7057ff | Epic-level work items |
| story | 🟢 #0e8a16 | User stories |
| priority:high | 🔴 #d93f0b | High priority |
| priority:medium | 🟡 #fbca04 | Medium priority |
| priority:low | 🟢 #0e8a16 | Low priority |

### Milestones ({count})
| # | Title | Stories | Status |
|---|-------|---------|--------|
| 1 | {Epic 1 Title} | {count} | Open |
| 2 | {Epic 2 Title} | {count} | Open |
| ... | ... | ... | ... |

### Issues ({count})
| Epic | Issue Numbers | Count |
|------|---------------|-------|
| {Epic 1} | #{start}-#{end} | {count} |
| {Epic 2} | #{start}-#{end} | {count} |
| ... | ... | ... |

---

## Statistics

| Metric | Value |
|--------|-------|
| Epics Created | {count} |
| Stories Created | {count} |
| Labels Created | {count} |
| FRs Covered | {count}/{total} (100%) |

---

## Next Steps

1. **Review** - Check issues on GitHub
2. **Prioritize** - Organize in GitHub Projects
3. **Assign** - Assign to team members
4. **Sprint** - Start with Epic 1, Story 1.1

---

## Milestone Links

{For each epic:}
- **{Epic Title}**: [View {count} issues](https://github.com/{owner}/{repo}/issues?q=is%3Aissue+milestone%3A%22{url_encoded_title}%22)

---

*Workflow: create-epics-and-stories-github*
*Completed: {timestamp}*
*User: {user_name}*
```

## URL Encoding Reference

For milestone filter URLs, encode special characters:
- Space → `%20` or `+`
- `&` → `%26`
- `#` → `%23`

Example:
- Milestone: "User Authentication & Profile"
- Encoded: `User%20Authentication%20%26%20Profile`
- Full URL: `https://github.com/owner/repo/issues?q=is%3Aissue+milestone%3A%22User%20Authentication%20%26%20Profile%22`

