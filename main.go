package main

import (
	"embed"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	AppMenu := menu.NewMenu()
	FileMenu := AppMenu.AddSubmenu("File")
	//FileMenu.AddText("&Open", keys.CmdOrCtrl("o"), openFile)
	FileMenu.AddSeparator()
	FileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		runtime.Quit(app.ctx)
	})

	SettingMenu := AppMenu.AddSubmenu("Settings")
	SettingMenu.AddText("&Switch Theme", keys.CmdOrCtrl("t"), func(data *menu.CallbackData) {
		changeTheme(app, data)
	})

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Gube",
		Width:  1024,
		Height: 768,
		Menu:   AppMenu, // reference the menu above
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 255},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func changeTheme(a *App, data *menu.CallbackData) {
	runtime.EventsEmit(a.ctx, "change-theme")
}
