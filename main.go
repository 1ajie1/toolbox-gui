package main

import (
	"context"
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Ltime)
}

func main() {
	// Create an instance of the app structure
	app := NewApp()
	tray := NewTrayManager()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "toolbox-gui2",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
			tray.SetContext(ctx)
			tray.Run()
			// 最小化窗口
			runtime.WindowHide(ctx)
		},
		OnBeforeClose: func(ctx context.Context) bool {
			// 返回true会阻止窗口关闭，这里我们隐藏窗口然后阻止关闭
			runtime.WindowHide(ctx)
			return true
		},
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			WindowIsTranslucent: true,
			BackdropType:        windows.Auto,
			DisableWindowIcon:   false,
			Theme:               windows.Dark,
		},
		OnDomReady: func(ctx context.Context) {
			runtime.WindowSetTitle(ctx, "Toolbox")
		},
		MinWidth:  800,
		MinHeight: 600,
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
