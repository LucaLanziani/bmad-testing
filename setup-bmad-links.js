#!/usr/bin/env node
/**
 * BMAD Multi-Repo Link Setup Script
 *
 * This script creates symbolic links to share BMAD configuration
 * across multiple repositories in the same parent folder.
 *
 * Usage: node setup-bmad-links.js [-y] [repo1] [repo2] ...
 *        node setup-bmad-links.js              # Auto-detects repos in parent folder
 *        node setup-bmad-links.js -y           # Skip confirmations
 *
 * Run this script from the parent folder containing all repos.
 */

const fs = require('fs');
const path = require('path');
const readline = require('readline');

const SCRIPT_DIR = __dirname;
const PARENT_DIR = path.dirname(SCRIPT_DIR);
const BMAD_CONFIG = path.basename(SCRIPT_DIR);

// Detect repos in parent folder (directories only, excluding current folder)
function detectRepos() {
  try {
    const entries = fs.readdirSync(PARENT_DIR, { withFileTypes: true });
    return entries
      .filter((entry) => {
        // Must be a directory
        if (!entry.isDirectory()) return false;
        // Skip hidden folders
        if (entry.name.startsWith('.')) return false;
        // Skip the current folder (where script is running from)
        if (entry.name === BMAD_CONFIG) return false;
        return true;
      })
      .map((entry) => entry.name);
  } catch {
    return [];
  }
}

// Symlinks to create in each repo
// linkPath: relative path in target repo (parent dirs will be created)
// sourcePath: path in project-specs to link to
const SYMLINK_DEFS = [
  { linkPath: '.bmad', sourcePath: '.bmad' },
  { linkPath: '.cursor/rules/bmad', sourcePath: '.cursor/rules/bmad' },
  { linkPath: 'specs', sourcePath: 'specs' },
];

// Folder in project-specs to create reverse links to repos
const CODE_FOLDER = 'code';

// Calculate relative path for symlink target
function getSymlinkTarget(repoPath, linkPath, sourcePath) {
  const linkDir = path.dirname(path.join(repoPath, linkPath));
  const sourceFullPath = path.join(SCRIPT_DIR, sourcePath);
  return path.relative(linkDir, sourceFullPath);
}

// Colors for output
const colors = {
  red: (text) => `\x1b[0;31m${text}\x1b[0m`,
  green: (text) => `\x1b[0;32m${text}\x1b[0m`,
  yellow: (text) => `\x1b[1;33m${text}\x1b[0m`,
};

// Parse command line arguments
function parseArgs() {
  const args = process.argv.slice(2);
  let autoConfirm = false;
  const repos = [];

  for (const arg of args) {
    if (arg === '-y' || arg === '--yes') {
      autoConfirm = true;
    } else {
      repos.push(arg);
    }
  }

  const detectedRepos = repos.length > 0 ? repos : detectRepos();
  return {
    autoConfirm,
    repos: detectedRepos,
    autoDetected: repos.length === 0,
  };
}

// Confirmation prompt
async function confirm(message, autoConfirm) {
  if (autoConfirm) {
    return true;
  }

  const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
  });

  return new Promise((resolve) => {
    rl.question(`${message} [y/N] `, (answer) => {
      rl.close();
      resolve(answer.toLowerCase() === 'y' || answer.toLowerCase() === 'yes');
    });
  });
}

// Check if path is a symlink
function isSymlink(filePath) {
  try {
    return fs.lstatSync(filePath).isSymbolicLink();
  } catch {
    return false;
  }
}

// Check if path is a directory (not following symlinks)
function isDirectory(filePath) {
  try {
    const stats = fs.lstatSync(filePath);
    return stats.isDirectory() && !stats.isSymbolicLink();
  } catch {
    return false;
  }
}

// Check if path exists (not following symlinks)
function exists(filePath) {
  try {
    fs.lstatSync(filePath);
    return true;
  } catch {
    return false;
  }
}

// Create symlink helper
function createSymlink(target, linkPath) {
  fs.symlinkSync(target, linkPath);
}

// Remove file/symlink
function remove(filePath) {
  fs.unlinkSync(filePath);
}

// Ensure directory exists
function ensureDir(dirPath) {
  fs.mkdirSync(dirPath, { recursive: true });
}

// Process a single symlink: check, confirm, create/replace
// Returns true if link was created/replaced, false if skipped
async function processSymlink({ linkPath, target, displayName, rmHint, autoConfirm }) {
  ensureDir(path.dirname(linkPath));

  if (isSymlink(linkPath)) {
    console.log(`  Existing ${displayName} symlink found`);
    if (await confirm(`  Replace ${displayName} symlink?`, autoConfirm)) {
      remove(linkPath);
      createSymlink(target, linkPath);
      console.log(colors.green(`  ✓ Linked ${displayName} -> ${target}`));
      return true;
    }
    console.log(colors.yellow(`  ⚠ Skipped ${displayName}`));
    return false;
  }

  if (isDirectory(linkPath)) {
    console.log(colors.yellow(`  ⚠ ${displayName} is a directory, not a symlink. Skipping...`));
    console.log(`    Remove it manually if you want to link: rm -rf ${rmHint}`);
    return false;
  }

  if (exists(linkPath)) {
    console.log(colors.yellow(`  ⚠ ${displayName} exists but is not a symlink. Skipping...`));
    return false;
  }

  if (await confirm(`  Create ${displayName} symlink?`, autoConfirm)) {
    createSymlink(target, linkPath);
    console.log(colors.green(`  ✓ Linked ${displayName} -> ${target}`));
    return true;
  }
  console.log(colors.yellow(`  ⚠ Skipped ${displayName}`));
  return false;
}

async function main() {
  const { autoConfirm, repos, autoDetected } = parseArgs();

  console.log('');
  console.log('==========================================');
  console.log('  BMAD Multi-Repo Link Setup');
  console.log('==========================================');
  console.log('');

  // Check if running from correct location
  const bmadPath = path.join(SCRIPT_DIR, '.bmad');
  if (!fs.existsSync(bmadPath)) {
    console.log(colors.red(`Error: .bmad folder not found in ${SCRIPT_DIR}`));
    console.log('Make sure this script is in the project-specs repository.');
    process.exit(1);
  }

  console.log(`Working directory: ${PARENT_DIR}`);
  console.log(`BMAD config: ${BMAD_CONFIG}`);
  if (repos.length === 0) {
    console.log(colors.yellow('No repos found in parent directory.'));
    process.exit(0);
  }
  console.log(`Target repos: ${repos.join(', ')}${autoDetected ? ' (auto-detected)' : ''}`);
  if (autoConfirm) {
    console.log('Mode: Auto-confirm enabled (-y)');
  } else {
    console.log('Mode: Interactive (use -y to skip confirmations)');
  }
  console.log('');

  // Track results
  let successCount = 0;
  let skipCount = 0;
  let failCount = 0;

  for (const repo of repos) {
    console.log('----------------------------------------');
    console.log(`Processing: ${repo}`);
    console.log('----------------------------------------');

    // When repos are explicitly passed, resolve relative to cwd; otherwise use PARENT_DIR
    const repoPath = autoDetected
      ? path.join(PARENT_DIR, repo)
      : path.resolve(process.cwd(), repo);

    if (!fs.existsSync(repoPath)) {
      console.log(colors.yellow('⚠ Skipping: directory not found'));
      skipCount++;
      continue;
    }

    // Don't link to self
    if (repo === BMAD_CONFIG) {
      console.log(colors.yellow('⚠ Skipping: cannot link to self'));
      skipCount++;
      continue;
    }

    // Build list of all links to process for this repo
    const codeFolderPath = path.join(SCRIPT_DIR, CODE_FOLDER);
    const links = [
      // Forward links: project-specs -> repo
      ...SYMLINK_DEFS.map(({ linkPath, sourcePath }) => ({
        linkPath: path.join(repoPath, linkPath),
        target: getSymlinkTarget(repoPath, linkPath, sourcePath),
        displayName: `${repo}/${linkPath}`,
        rmHint: `${repo}/${linkPath}`,
      })),
      // Reverse link: repo -> code/ folder
      {
        linkPath: path.join(codeFolderPath, repo),
        target: path.relative(codeFolderPath, repoPath),
        displayName: `${CODE_FOLDER}/${repo}`,
        rmHint: `${CODE_FOLDER}/${repo}`,
      },
    ];

    // Process all symlinks
    let anyLinked = false;
    for (const link of links) {
      if (await processSymlink({ ...link, autoConfirm })) {
        anyLinked = true;
      }
    }

    // Count as success only if at least one link was created
    if (anyLinked) {
      successCount++;
    } else {
      skipCount++;
    }

    console.log('');
  }

  console.log('==========================================');
  console.log('  Summary');
  console.log('==========================================');
  console.log(colors.green(`Successful: ${successCount}`));
  console.log(colors.yellow(`Skipped: ${skipCount}`));
  if (failCount > 0) {
    console.log(colors.red(`Failed: ${failCount}`));
  }
  console.log('');

  if (successCount > 0) {
    console.log('BMAD is now linked! Open your repos in Cursor and use:');
    console.log('  @bmad/bmm/agents/dev');
    console.log('  @bmad/bmm/agents/architect');
    console.log('  @bmad/core/agents/bmad-master');
    console.log('');
    console.log(`Reverse links created in ${CODE_FOLDER}/ for quick repo access.`);
    console.log('');
  }
}

main().catch((err) => {
  console.error(colors.red(`Error: ${err.message}`));
  process.exit(1);
});

