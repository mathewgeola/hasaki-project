package cmd

import (
	"hasaki-project/internal/gitops"

	"github.com/spf13/cobra"
)

var vcommitCmd = &cobra.Command{
	Use:   "vcommit [dirPath]",
	Short: "执行带版本号的代码提交",
	Long: `执行携带版本信息的标准化 Git 提交工作流。

如果未提供 [dirPath] 参数，默认在当前工作目录 (.) 执行操作。`,
	Example: `  # 在当前目录执行版本提交 (默认行为)
  hp vcommit

  # 在指定相对路径的项目执行版本提交
  hp vcommit ./my-project

  # 在指定绝对路径的项目执行版本提交
  hp vcommit /path/to/my-project`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dirPath := "."
		if len(args) == 1 {
			dirPath = args[0]
		}

		gitops.VersionCommit(dirPath)
	},
}

func init() {
	rootCmd.AddCommand(vcommitCmd)
}
