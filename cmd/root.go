package cmd

import (
	"fmt"
	"os"

	"codeberg.org/filipmnowak/spidercm/cmd/scm"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "spidercm",
	Short: "manage source using multiple SCMs at the same time",
}

func init() {
	rootCmd.AddCommand(scm.NewInitCmd())
	rootCmd.AddCommand(scm.NewCommitCmd())
	rootCmd.AddCommand(scm.NewAddCmd())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
