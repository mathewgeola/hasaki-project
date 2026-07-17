package cmd

import (
	"github.com/mathewgeola/hasaki-project/internal/gitops"

	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit [dirPath]",
	Short: "执行标准化的代码提交",
	Long: `执行标准化的 Git 提交工作流。

如果未提供 [dirPath] 参数，默认在当前工作目录 (.) 执行提交操作。`,
	Example: `  # 提交当前目录的变更 (默认行为)
  hp commit

  # 提交指定相对路径的项目
  hp commit ./my-project

  # 提交指定绝对路径的项目
  hp commit /path/to/my-project`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dirPath := "."
		if len(args) == 1 {
			dirPath = args[0]
		}

		gitops.Commit(dirPath)
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
