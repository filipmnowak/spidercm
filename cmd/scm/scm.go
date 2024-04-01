package scm

import (
	"fmt"
	"os"

	"codeberg.org/filipmnowak/spidercm/internal/scm"
	"github.com/spf13/cobra"
)

func NewInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "initializer SCMs",
		Run: func(_ *cobra.Command, _ []string) {
			Init()
		},
	}
	return cmd
}

func Init() {
	scmCmd := scm.SCM{}
	if err := scmCmd.Init(); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", scmCmd.Results)
}

func NewCommitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "commit",
		Short: "perform a commit",
		Run: func(cmd *cobra.Command, _ []string) {
			Commit(cmd)
		},
	}
	cmd.Flags().StringP("commit_message", "m", "", "commit message")
	cmd.MarkFlagRequired("commit_message")
	return cmd
}

func Commit(cmd *cobra.Command) {
	scmCmd := scm.SCM{}
	commitMessage, _ := cmd.Flags().GetString("commit_message")
	if err := scmCmd.Commit(commitMessage); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", scmCmd.Results)
}

func NewAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "perform a add",
		Run: func(cmd *cobra.Command, _ []string) {
			Add(cmd)
		},
	}
	cmd.Flags().StringSliceP("paths", "p", []string{}, "paths to add")
	cmd.MarkFlagRequired("paths")
	return cmd
}

func Add(cmd *cobra.Command) {
	scmCmd := scm.SCM{}
	paths, _ := cmd.Flags().GetStringSlice("paths")
	if err := scmCmd.Add(paths); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", scmCmd.Results)
}
