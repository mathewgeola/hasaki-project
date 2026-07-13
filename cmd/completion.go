package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion <bash|zsh|fish|powershell>",
	Short: "生成命令自动补全脚本",
	Long: `为你的终端生成自动补全脚本。

【快速配置指南】

Bash 用户:
  进入环境: bash (通常大部分 Linux/Mac 系统已内置)
  临时体验: source <(hp completion bash)
  永久生效: 
    mkdir -p ~/.local/share/bash-completion/completions
    hp completion bash > ~/.local/share/bash-completion/completions/hp

Zsh 用户:
  安装环境: sudo apt install zsh (Mac 默认已内置)
  进入环境: zsh
  临时体验: source <(hp completion zsh)
  永久生效: 
    mkdir -p ~/.zfunc
    hp completion zsh > ~/.zfunc/_hp
    echo 'fpath=(~/.zfunc $fpath)' >> ~/.zshrc
    echo 'compinit' >> ~/.zshrc

Fish 用户:
  安装环境: sudo apt install fish
  进入环境: fish
  临时体验: hp completion fish | source
  永久生效: 
    hp completion fish > ~/.config/fish/completions/hp.fish

PowerShell 用户:
  安装环境: sudo snap install powershell --classic (Windows 默认已内置)
  进入环境: pwsh
  临时体验: hp completion powershell | Out-String | Invoke-Expression
  永久生效: 
    hp completion powershell >> $PROFILE
`,
	Example: `  # 临时加载 Bash 的自动补全 (当前会话生效)
  source <(hp completion bash)

  # 生成 Bash 补全脚本并保存到本地目录 (永久生效)
	mkdir -p ~/.local/share/bash-completion/completions
    hp completion bash > ~/.local/share/bash-completion/completions/hp`,
	ValidArgs: []string{"bash", "zsh", "fish", "powershell"},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			// true 参数代表生成带描述信息的补全提示，体验更好
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletion(os.Stdout)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
