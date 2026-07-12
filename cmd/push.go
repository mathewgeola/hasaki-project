package cmd

import (
	"hasaki-project/internal/gitops"

	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push [*dirPath]",
	Short: "git push",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var dirPath string
		if len(args) == 1 {
			dirPath = args[0]
		}
		gitops.Push(dirPath)
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}
