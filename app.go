package main

import (
	"context"
	"fmt"
	"toolbox-gui2/backend/terminal"
)

// App struct
type App struct {
	ctx      context.Context
	terminal *terminal.Terminal
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// 初始化终端管理器
	term, err := terminal.NewTerminal()
	if err != nil {
		fmt.Printf("初始化终端管理器失败: %v\n", err)
		return
	}
	a.terminal = term
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// OpenTerminal 打开终端并设置环境变量
func (a *App) OpenTerminal(terminalType string) error {
	if a.terminal == nil {
		return fmt.Errorf("终端管理器未初始化")
	}
	return a.terminal.OpenTerminal(terminalType)
}
