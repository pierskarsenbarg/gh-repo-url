package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func newTestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "gh-repo-url",
		Args:         cobra.ExactArgs(0),
		Short:        "GitHub CLI extension for getting current repository's URL",
		SilenceUsage: true,
		RunE:         rootCmd.RunE,
	}
	cmd.Flags().BoolP("ssh", "s", false, "Return repo url in format for ssh. i.e. git@github.com:pierskarsenbarg/gh-repo-url.git")
	return cmd
}

func TestRootCommand_DefaultHTTPSFormat(t *testing.T) {
	cmd := newTestCmd()
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs([]string{})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("command failed: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "https://") {
		t.Errorf("expected HTTPS output, got: %s", output)
	}
	if strings.Contains(output, "git@") {
		t.Errorf("expected HTTPS format but got SSH format: %s", output)
	}
}

func TestRootCommand_SSHFormat(t *testing.T) {
	cmd := newTestCmd()
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs([]string{"--ssh"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("command failed: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "git@") {
		t.Errorf("expected SSH format, got: %s", output)
	}
	if !strings.Contains(output, ".git") {
		t.Errorf("expected SSH format to end with .git, got: %s", output)
	}
}

func TestRootCommand_SSHFlagShortForm(t *testing.T) {
	cmd := newTestCmd()
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs([]string{"-s"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("command failed: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "git@") {
		t.Errorf("expected SSH format with -s flag, got: %s", output)
	}
}

func TestRootCommand_HTTPSFormatContainsHostAndRepo(t *testing.T) {
	cmd := newTestCmd()
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs([]string{})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("command failed: %v", err)
	}

	output := buf.String()
	// Should contain at least one slash after domain and have repo name
	parts := strings.Split(output, "/")
	if len(parts) < 3 {
		t.Errorf("HTTPS URL should have at least 3 parts separated by /, got: %s", output)
	}
}

func TestRootCommand_SSHFormatContainsHostAndRepo(t *testing.T) {
	cmd := newTestCmd()
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs([]string{"--ssh"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("command failed: %v", err)
	}

	output := buf.String()
	// Should contain git@, :, and /
	if !strings.Contains(output, "git@") || !strings.Contains(output, ":") {
		t.Errorf("SSH URL should contain git@ and :, got: %s", output)
	}
	if !strings.Contains(output, "/") {
		t.Errorf("SSH URL should contain owner/repo, got: %s", output)
	}
}

func TestRootCommand_RejectsArguments(t *testing.T) {
	cmd := newTestCmd()
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs([]string{"extra-arg"})

	err := cmd.Execute()
	if err == nil {
		t.Error("expected error when providing arguments, but command succeeded")
	}
}
