# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**gh-repo-url** is a GitHub CLI extension that retrieves and outputs the current repository's URL. It queries the GitHub repository metadata from the current git directory and returns either an HTTPS or SSH format URL.

## Architecture

The project follows a simple CLI structure:

- **Entry Point**: `main.go` - Minimal entry point that calls `cmd.Execute()`
- **CLI Framework**: Uses [Cobra](https://github.com/spf13/cobra) for command structure and flag parsing
- **GitHub Integration**: Uses `github.com/cli/go-gh/v2` to access GitHub CLI's repository context via the `repository.Current()` API
- **Single Command**: Root command in `cmd/root.go` with optional `--ssh` flag to switch output format

### Command Behavior

The root command:
- Takes no arguments (`cobra.ExactArgs(0)`)
- Fetches the current repository using `repository.Current()`
- Outputs HTTPS URL by default: `https://{host}/{owner}/{name}`
- Outputs SSH URL with `--ssh` flag: `git@{host}:{owner}/{name}.git`

## Build & Run

```bash
# Build the extension
go build -o gh-repo-url

# Install as a GitHub CLI extension (for local development)
gh extension install .

# Run the extension
gh repo-url
gh repo-url --ssh  # or -s for short form
```

## Dependencies

Key dependencies:
- `github.com/cli/go-gh/v2` - GitHub CLI SDK for repository access
- `github.com/spf13/cobra` - CLI framework
- `github.com/spf13/pflag` - Flag parsing

## Testing

No tests are currently included. The CLI can be tested manually or through GitHub's extension testing mechanisms.

## Release Process

The project uses GitHub Actions for automated releases:
- **Trigger**: Tags matching `v*` (e.g., `v0.0.1`)
- **Action**: `cli/gh-extension-precompile@v2` precompiles binaries for multiple platforms
- **Artifacts**: Includes build attestations for security

To release:
```bash
git tag v0.0.X
git push origin v0.0.X
```

## Key Files

- `main.go` - Entry point
- `cmd/root.go` - Root command with all business logic
- `.github/workflows/release.yml` - Release automation
- `go.mod` - Go module definition (requires Go 1.25.5+)
