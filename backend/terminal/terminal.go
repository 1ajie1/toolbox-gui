package terminal

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// Terminal 终端管理器结构体
type Terminal struct {
	toolsPath string
}

// NewTerminal 创建新的终端管理器
func NewTerminal() (*Terminal, error) {
	// 获取当前工作目录
	workDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("获取工作目录失败: %v", err)
	}

	// 获取tools目录的绝对路径
	toolsPath := filepath.Join(workDir, "tools")

	// 确保tools目录存在
	if err := os.MkdirAll(toolsPath, 0755); err != nil {
		return nil, fmt.Errorf("创建tools目录失败: %v", err)
	}

	return &Terminal{
		toolsPath: toolsPath,
	}, nil
}

// OpenTerminal 打开指定类型的终端
func (t *Terminal) OpenTerminal(terminalType string) error {
	if runtime.GOOS != "windows" {
		return fmt.Errorf("当前只支持Windows系统")
	}

	var cmd *exec.Cmd

	switch terminalType {
	case "powershell":
		// 使用 cmd.exe start 命令启动 PowerShell
		psCommand := fmt.Sprintf(
			"$env:PATH = $env:PATH + ';%s'; "+
				"Write-Host '已将工具目录添加到PATH环境变量: %s'; "+
				"Set-Location '%s'; "+
				"$Host.UI.RawUI.WindowTitle = '运维工具箱 - PowerShell'",
			t.toolsPath, t.toolsPath, t.toolsPath,
		)

		// 使用 cmd.exe /c start 来启动新窗口
		cmd = exec.Command("cmd.exe", "/c", "start", "powershell.exe", "-NoExit", "-Command", psCommand)
		log.Printf("启动PowerShell，设置环境变量PATH=%s\n", t.toolsPath)

	case "cmd":
		// 使用直接的CMD启动方式
		cmd = exec.Command("cmd.exe", "/c", "start", "cmd.exe", "/k",
			fmt.Sprintf("set PATH=%%PATH%%;%s& title CMD",
				t.toolsPath),
		)

		log.Printf("启动CMD，命令如下：%s\n", cmd)

	default:
		return fmt.Errorf("不支持的终端类型: %s", terminalType)
	}

	// 启动进程
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("启动终端失败: %v", err)
	}

	return nil
}

// contains 检查字符串是否包含子串（不区分大小写）
func contains(s, substr string) bool {
	s = strings.ToLower(s)
	substr = strings.ToLower(substr)
	return strings.Contains(s, substr)
}
