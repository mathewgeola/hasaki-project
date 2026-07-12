package cmd

import (
	"hasaki-project/internal/gitops"

	"github.com/spf13/cobra"
)

var vcommitCmd = &cobra.Command{
	Use:   "vcommit [*dirPath]",
	Short: "git version commit",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var dirPath string
		if len(args) == 1 {
			dirPath = args[0]
		}
		gitops.VersionCommit(dirPath)
	},
}

func init() {
	rootCmd.AddCommand(vcommitCmd)
}
