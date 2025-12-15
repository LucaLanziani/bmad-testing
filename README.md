# BMAD Multi-Repo Configuration

This repository contains the centralized BMAD (BMad Method) configuration shared across multiple code repositories. By keeping BMAD config in a dedicated repo, all your project repos can use the same agents, workflows, and settings.

The current repo comes with BMAD v6 pre-installed with the bmm and bmb modules.

This repo is only used to store specs and share agents, the project code is in the `./code` folder and that's BMAD should use the content of that folder to evaluate the status of the project.

## Prerequisites

- Git
- A BMAD-compatible IDE (Cursor, VS Code with extensions, Windsurf, Claude Code)
- For Cursor users: Cursor IDE with rules support

## Initial Setup

### 1. Clone All Repositories

Ensure all your repositories are in the same parent folder:

```bash
mkdir your-project && cd your-project

# Clone or create the BMAD config repo
git clone https://github.com/nearform/bmad-starter-kit project-specs

# Clone your code repositories
git clone <your-backend-repo-url>
git clone <your-frontend-repo-url>
git clone <your-infra-repo-url>
```
                                                                      
### 2. Link BMAD to Each Repository and backwards

Run the setup script to create symlinks between all repos:

```bash
./project-specs/setup-bmad-links.js
```

<details>
<summary>I you want to learn more:</summary>

The script will:
- Auto-detect repositories in the parent folder
- Prompt for confirmation before creating each symlink
- Create `.bmad` symlinks in each repo (BMAD configuration)
- Create `specs` symlinks in each repo (project specifications)
- Create `.cursor/rules/bmad` symlinks for Cursor IDE
- Create reverse links in `code/` folder for quick access to all linked repos
- Skip repos that don't exist or are already directories

#### Skip Confirmations

Use the `-y` flag to skip all confirmations:

```bash
./project-specs/setup-bmad-links.js -y
```

#### Custom Repositories

To link specific repos (instead of auto-detected ones):

```bash
./project-specs/setup-bmad-links.js my-api my-web my-mobile

# Or with auto-confirm
./project-specs/setup-bmad-links.js -y my-api my-web my-mobile
```
</details>

### 3. Start using BMAD

In the project-specs folder using your preferred Agentic system:

```
*init-workflow
```


### 4. Customize BMAD config and *install custom agents*

We encourage you to run the commmand:

```bash
cd project-specs
npx bmad-method@alpha install
```

and select `Modify BMAD Installation (Confirm or change each setting)` to customize the installation to your need.

You can also install custom agents using:

```bash
cd project-specs
npx bmad-method@alpha agent-install
```

## Windows Users

The setup script is a Node.js script and works cross-platform. Run it with:

```bash
node project-specs/setup-bmad-links.js
```

Alternatively, create junction points manually in Command Prompt (as Administrator):

```cmd
mklink /J backend\.bmad project-specs\.bmad
mklink /J backend\specs project-specs\specs
mkdir backend\.cursor\rules
mklink /J backend\.cursor\rules\bmad project-specs\.cursor\rules\bmad
mklink /J project-specs\code\backend ..\backend
```

## Resources

- [BMAD Method Quick Start](.bmad/bmm/docs/quick-start.md)
- [Agents Guide](.bmad/bmm/docs/agents-guide.md)
- [Cursor Instructions](.bmad/docs/cursor-instructions.md)
- [BMAD GitHub](https://github.com/bmad-code-org/BMAD-METHOD)
- [Discord Community](https://discord.gg/gk8jAdXWmj)


