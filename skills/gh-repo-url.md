# gh-repo-url Extension Skill

## What This Tool Does

The **gh-repo-url** extension is a GitHub CLI command that outputs the full HTTPS URL of the current repository in the format `https://host/owner/name`.

## When to Use This Tool

Use this extension when:

- A user asks for the URL of their repository
- You need to construct URLs for GitHub API calls
- You need to create direct links to the repository on GitHub
- You're building automation that requires the repository URL
- You need to determine the repository's host (for GitHub Enterprise support)

## Prerequisites

Before using this tool, verify:

1. The user is currently in a Git repository directory
2. The repository is connected to GitHub (or a GitHub Enterprise instance)
3. The GitHub CLI (`gh` command) is installed and configured with proper authentication
4. The user has the gh-repo-url extension installed as a GitHub CLI extension

## Installation

If not already installed:

```bash
gh extension install pierskarsenbarg/gh-repo-url
```

## Usage

### Basic Command

Run from any directory within a Git repository:

```bash
gh repo-url
```

Or directly with the built binary:

```bash
./gh-repo-url
```

### Expected Output

The command outputs a single line with the repository URL:

```
https://github.com/owner/repo-name
```

### Output Format Details

- **Host**: The GitHub domain (e.g., `github.com` for public GitHub, custom domain for Enterprise)
- **Owner**: The repository owner (user or organization)
- **Name**: The repository name

### Capturing Output

In shell scripts, capture the output:

```bash
repo_url=$(gh repo-url)
echo "Repository URL: $repo_url"
```

## Error Scenarios

### Not in a Git Repository

**Error**: The command fails silently or outputs nothing
**User Communication**: "You're not currently in a Git repository. Please run this command from within a Git repository directory."

### No GitHub Context

**Error**: The command fails or outputs an error message
**User Communication**: "This repository doesn't have a GitHub context. Make sure you're in a repository that's connected to GitHub."

### Missing Extension

**Error**: Command not found
**User Communication**: "The gh-repo-url extension isn't installed. Install it with: `gh extension install pierskarsenbarg/gh-repo-url`"

## Integration Examples

### Creating Repository Links in Automation

```bash
repo_url=$(gh repo-url)
# Create link to specific file
file_link="$repo_url/blob/main/README.md"
echo "Documentation: $file_link"
```

### Integration with CI/CD Workflows

Use in GitHub Actions or other CI systems to dynamically get the repository URL:

```bash
gh repo-url  # Outputs the URL for logging or downstream tasks
```

## Troubleshooting Tips for AI Assistants

1. **Always verify the user is in a Git repository** before suggesting this command
2. **Confirm GitHub CLI is installed** with `gh --version` if commands fail
3. **Check authentication** with `gh auth status` if there are permission issues
4. **For Enterprise GitHub**, the output will reflect the custom domain, which is expected behavior
5. **The output is a simple URL string** - guide users on how to use it in their specific context
