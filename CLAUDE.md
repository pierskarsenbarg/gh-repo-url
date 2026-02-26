# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**gh-repo-url** is a GitHub CLI extension written in Go that outputs the URL of the current repository in the format `https://host/owner/name`. The project also includes Pulumi infrastructure code to manage the GitHub repository configuration.

## Skills Documentation

Agent instructions and usage guidance for this extension are available in the `/skills/` directory:
- `skills/gh-repo-url.md` - Comprehensive skill documentation for AI agents on when and how to use this extension, including error scenarios, integration examples, and troubleshooting tips.

## Architecture

### Main Components

1. **Go CLI Tool** (`main.go`): The core extension that uses the `github.com/cli/go-gh/v2` library to:
   - Detect the current repository context from the GitHub CLI environment
   - Extract repository metadata (host, owner, name)
   - Output the formatted repository URL

2. **Pulumi Infrastructure** (`infra/`): TypeScript-based infrastructure-as-code that:
   - Manages the GitHub repository configuration itself
   - Sets up repository features (issues, projects, wiki, forking)
   - Configures branch protection and default branch
   - Manages repository topics (tags)
   - Uses GitHub provider for infrastructure management
   - Built to be run against the same repository it manages

## Common Commands

### Building
```bash
go build -o gh-repo-url
```

### Running the extension
```bash
./gh-repo-url
# or if installed as gh extension:
gh repo-url
```

### Testing locally
From within a Git repository:
```bash
./gh-repo-url
# Should output: https://github.com/owner/repo-name
```

### Infrastructure management (Pulumi)
```bash
cd infra

# Install dependencies
bun install  # or npm install

# Preview infrastructure changes
pulumi preview

# Deploy infrastructure changes
pulumi up

# Show stack outputs
pulumi stack output
```

## Release Process

The project uses GitHub Actions for automated releases. When a tag matching `v*` is pushed:
1. GitHub Actions workflow (`.github/workflows/release.yml`) triggers
2. Uses `cli/gh-extension-precompile@v2` to build binaries for multiple platforms
3. Generates build attestations for security

To release:
```bash
git tag v1.0.0
git push origin v1.0.0
```

## Key Dependencies

- **Go**: `github.com/cli/go-gh/v2` - GitHub CLI integration library
- **Pulumi**: `@pulumi/pulumi`, `@pulumi/github` - Infrastructure-as-code framework
- **TypeScript**: Dev dependency for infra code type checking

## Pulumi Stack Configuration

- Stack name: `dev` (default)
- Config files:
  - `Pulumi.yaml` - Project metadata
  - `Pulumi.dev.yaml` - Stack-specific configuration
- Both files use protected resources to prevent accidental deletion

## Notable Implementation Details

- The Go tool leverages the `repository.Current()` function from go-gh to automatically detect the current repository context from the GitHub CLI environment
- Pulumi resources are marked with `protect: true` to ensure the GitHub repository configuration isn't accidentally deleted
- The infra code assumes it's running with credentials that have GitHub API access (via Pulumi configuration)
