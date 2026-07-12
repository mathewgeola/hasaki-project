package cmd

import (
	"hasaki-project/internal/gitops"

	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:   "tag [dirPath]",
	Short: "git tag",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var dirPath string
		if len(args) == 1 {
			dirPath = args[0]
		}
		gitops.CreateTag(dirPath)
	},
}

func init() {
	rootCmd.AddCommand(tagCmd)
}
