package main

import (
	"context"
	_ "embed"
	"os"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed layout_left_icon_264872.ico
var iconBytes []byte

type TrayManager struct {
	ctx context.Context
}

func NewTrayManager() *TrayManager {
	return &TrayManager{}
}

func (t *TrayManager) SetContext(ctx context.Context) {
	t.ctx = ctx
}

func (t *TrayManager) onReady() {
	systray.SetIcon(iconBytes)
	systray.SetTitle("Toolbox")
	systray.SetTooltip("Toolbox")

	mShow := systray.AddMenuItem("显示", "显示主窗口")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("退出", "退出应用")

	go func() {
		for {
			select {
			case <-mShow.ClickedCh:
				runtime.WindowShow(t.ctx)
			case <-mQuit.ClickedCh:
				systray.Quit()
				os.Exit(0)
				return
			}
		}
	}()
}

func (t *TrayManager) onExit() {
	// 清理工作
}

func (t *TrayManager) Run() {
	go systray.Run(t.onReady, t.onExit)
}
