package cmd

import (
	"github.com/mathewgeola/hasaki-project/internal/gitops"

	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push [dirPath]",
	Short: "推送代码到远程仓库",
	Long: `执行标准化的 Git 推送工作流。

如果未提供 [dirPath] 参数，默认在当前工作目录 (.) 执行推送操作。`,
	Example: `  # 推送当前目录的变更 (默认行为)
  hp push

  # 推送指定相对路径的项目
  hp push ./my-project

  # 推送指定绝对路径的项目
  hp push /path/to/my-project`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dirPath := "."
		if len(args) == 1 {
			dirPath = args[0]
		}

		gitops.Push(dirPath)
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}
