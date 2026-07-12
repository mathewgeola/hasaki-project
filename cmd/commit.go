package cmd

import (
	"hasaki-project/internal/gitops"

	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit [*dirPath]",
	Short: "git commit",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var dirPath string
		if len(args) == 1 {
			dirPath = args[0]
		}
		gitops.Commit(dirPath)

	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
