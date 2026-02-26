package cmd

import (
	"fmt"
	"os"

	"github.com/cli/go-gh/v2/pkg/repository"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "gh-repo-url",
	Args:         cobra.ExactArgs(0),
	Short:        "GitHub CLI extension for getting current repository's URL",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, err := repository.Current()
		if err != nil {
			return err
		}
		ssh, err := cmd.Flags().GetBool("ssh")
		if err != nil {
			return err
		}
		formatString := "https://%s/%s/%s"
		if ssh {
			fmt.Fprintf(cmd.OutOrStdout(), "git@%s:%s/%s.git", repo.Host, repo.Owner, repo.Name)
		} else {
			fmt.Fprintf(cmd.OutOrStdout(), formatString, repo.Host, repo.Owner, repo.Name)
		}
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("ssh", "s", false, "Return repo url in format for ssh. i.e. git@github.com:pierskarsenbarg/gh-repo-url.git")
}
