package gitops

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func ensureGitEnv(dirPath string) error {
	if dirPath == "" {
		currentDir, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("无法获取当前工作目录: %w", err)
		}
		dirPath = currentDir
	}

	if strings.HasPrefix(dirPath, "~") {
		home, err := os.UserHomeDir()
		if err == nil {
			dirPath = strings.Replace(dirPath, "~", home, 1)
		}
	}

	info, err := os.Stat(dirPath)
	if err != nil || !info.IsDir() {
		return fmt.Errorf("目录 '%s' 不存在或不是一个有效目录", dirPath)
	}

	if err := os.Chdir(dirPath); err != nil {
		return fmt.Errorf("无法进入目录 '%s': %w", dirPath, err)
	}

	if err := exec.Command("git", "rev-parse", "--is-inside-work-tree").Run(); err != nil {
		return fmt.Errorf("该目录 '%s' 不是一个有效的 Git 仓库", dirPath)
	}

	return nil
}

func runGitCmd(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	out, err := cmd.Output()
	return strings.TrimSpace(string(out)), err
}

func runGitCmdWithOutput(args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Commit(dirPath string) {
	if err := ensureGitEnv(dirPath); err != nil {
		fmt.Printf("❌ 环境校验失败: %v\n", err)
		os.Exit(1)
	}

	now := time.Now().Format("2006-01-02 15:04:05")

	baseHash, err := runGitCmd("merge-base", "HEAD", "origin/main")
	if err != nil || baseHash == "" {
		fmt.Println("❌ 错误: 无法找到 HEAD 与 origin/main 的 merge-base。")
		fmt.Println("请确认远程是否存在 main 分支，并且本地已执行过 git fetch。")
		os.Exit(1)
	}

	if _, err := runGitCmd("reset", baseHash); err != nil {
		fmt.Println("❌ 错误: git reset 执行失败")
		os.Exit(1)
	}

	if _, err := runGitCmd("add", "-A"); err != nil {
		fmt.Println("❌ 错误: git add 执行失败")
		os.Exit(1)
	}

	if err := runGitCmdWithOutput("commit", "-m", now); err != nil {
		fmt.Println("❌ 错误: git commit 执行失败")
		os.Exit(1)
	}
}

func VersionCommit(dirPath string) {
	if err := ensureGitEnv(dirPath); err != nil {
		fmt.Printf("❌ 环境校验失败: %v\n", err)
		os.Exit(1)
	}

	data, err := os.ReadFile("VERSION")
	if err != nil {
		fmt.Printf("❌ 错误: 无法读取 VERSION 文件，请确保目录下存在该文件。\n详细信息: %v\n", err)
		os.Exit(1)
	}

	rawVersion := strings.TrimSpace(string(data))
	if rawVersion == "" {
		fmt.Println("❌ 错误: VERSION 文件内容为空。")
		os.Exit(1)
	}

	commitMsg := "v" + rawVersion

	baseHash, err := runGitCmd("merge-base", "HEAD", "origin/main")
	if err != nil || baseHash == "" {
		fmt.Println("❌ 错误: 无法找到 HEAD 与 origin/main 的 merge-base。")
		os.Exit(1)
	}

	if _, err := runGitCmd("reset", baseHash); err != nil {
		fmt.Println("❌ 错误: git reset 执行失败")
		os.Exit(1)
	}

	if _, err := runGitCmd("add", "-A"); err != nil {
		fmt.Println("❌ 错误: git add 执行失败")
		os.Exit(1)
	}

	if err := runGitCmdWithOutput("commit", "-m", commitMsg); err != nil {
		fmt.Println("❌ 错误: git commit 执行失败")
		os.Exit(1)
	}
}

func CreateTag(dirPath string) {
	if err := ensureGitEnv(dirPath); err != nil {
		fmt.Printf("❌ 环境校验失败: %v\n", err)
		os.Exit(1)
	}

	data, err := os.ReadFile("VERSION")
	if err != nil {
		fmt.Printf("❌ 错误: 无法读取 VERSION 文件，请确保目录下存在该文件。\n详细信息: %v\n", err)
		os.Exit(1)
	}

	rawVersion := strings.TrimSpace(string(data))
	if rawVersion == "" {
		fmt.Println("❌ 错误: VERSION 文件内容为空。")
		os.Exit(1)
	}

	tagName := "v" + rawVersion

	if err := runGitCmdWithOutput("tag", tagName); err != nil {
		fmt.Printf("❌ 错误: 创建本地 Tag '%s' 失败，可能该版本号对应的 Tag 已存在。\n", tagName)
		os.Exit(1)
	}
}

func Push(dirPath string) {
	if err := ensureGitEnv(dirPath); err != nil {
		fmt.Printf("❌ 环境校验失败: %v\n", err)
		os.Exit(1)
	}

	if err := runGitCmdWithOutput("push", "-f", "origin", "main"); err != nil {
		fmt.Println("❌ 错误: git push 代码执行失败")
		os.Exit(1)
	}

	if err := runGitCmdWithOutput("push", "origin", "--tags"); err != nil {
		fmt.Println("❌ 错误: git push tags 执行失败")
		os.Exit(1)
	}
}
