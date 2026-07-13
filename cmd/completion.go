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
  安装环境: sudo apt install bash-completion
  进入环境: bash (通常大部分 Linux/Mac 系统已内置)
  临时体验: source <(hp completion bash)
  永久生效: 
    mkdir -p ~/.local/share/bash-completion/completions
    hp completion bash > ~/.local/share/bash-completion/completions/hp
    echo 'source /usr/share/bash-completion/bash_completion' >> ~/.bashrc
    source ~/.bashrc

Zsh 用户:
  安装环境: sudo apt install zsh
  进入环境: zsh (Mac 默认已内置)
  临时体验: source <(hp completion zsh)
  永久生效: 
    mkdir -p ~/.zfunc
    hp completion zsh > ~/.zfunc/_hp
    echo 'fpath=(~/.zfunc $fpath)' >> ~/.zshrc
    echo 'compinit' >> ~/.zshrc
    source ~/.zshrc

Fish 用户:
  安装环境: sudo apt install fish
  进入环境: fish
  临时体验: hp completion fish | source
  永久生效: 
    mkdir -p ~/.config/fish/completions
    hp completion fish > ~/.config/fish/completions/hp.fish
    source ~/.config/fish/completions/hp.fish

PowerShell 用户:
  安装环境: sudo snap install powershell --classic
  进入环境: pwsh (Windows 默认已内置)
  临时体验: hp completion powershell | Out-String | Invoke-Expression
  永久生效: 
    New-Item -Type Directory -Force -Path (Split-Path -Parent $PROFILE)
    hp completion powershell >> $PROFILE
    . $PROFILE
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
