# gh-repo-url

GitHub CLI extension to return the url for the repository associated with the current directory.

## Installation

You need to have the GitHub CLI already installed: <https://cli.github.com> and you need to be logged in as well: `gh login`.

To install this extension run:

```console
gh extension install pierskarsenbarg/gh-repo-url
```

## Usage

To get the URL, run the following: `gh repo-url`:

```console
❯ gh repo-url
https://github.com/pierskarsenbarg/gh-repo-url
```

To get the URL in a format for SSH (if you use git over ssh): `gh repo-url --ssh`:

```console
❯ gh repo-url --ssh
git@github.com:pierskarsenbarg/gh-repo-url.git
```

For all options, including shortened flags:

```console
❯ gh repo-url --help
GitHub CLI extension for getting current repository's URL

Usage:
  gh-repo-url [flags]

Flags:
  -h, --help   help for gh-repo-url
  -s, --ssh    Return repo url in format for ssh. i.e. git@github.com:pierskarsenbarg/gh-repo-url.git
```
