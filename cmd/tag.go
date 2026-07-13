package cmd

import (
	"hasaki-project/internal/gitops"

	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:   "tag [dirPath]",
	Short: "创建标准化的 Git 标签",
	Long: `执行标准化的 Git 标签 (Tag) 创建工作流。

如果未提供 [dirPath] 参数，默认在当前工作目录 (.) 执行操作。`,
	Example: `  # 在当前目录创建标签 (默认行为)
  hp tag

  # 在指定相对路径的项目创建标签
  hp tag ./my-project

  # 在指定绝对路径的项目创建标签
  hp tag /path/to/my-project`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dirPath := "."
		if len(args) == 1 {
			dirPath = args[0]
		}

		gitops.CreateTag(dirPath)
	},
}

func init() {
	rootCmd.AddCommand(tagCmd)
}
