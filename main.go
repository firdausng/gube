package main

import (
	"context"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gube/backend"
	"gube/backend/services"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	go startServer()

	// Create an instance of the app structure
	app := backend.NewApp()
	contextService := services.NewContextService()
	workspaceService := services.NewWorkspaceService()
	podService := services.NewPodService()
	namespaceService := services.NewNamespaceService()

	AppMenu := menu.NewMenu()
	FileMenu := AppMenu.AddSubmenu("File")
	//FileMenu.AddText("&Open", keys.CmdOrCtrl("o"), openFile)
	FileMenu.AddSeparator()
	FileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		runtime.Quit(app.Ctx)
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
		OnStartup: func(ctx context.Context) {
			contextService.SetContext(ctx)
			workspaceService.SetContext(ctx)
			podService.SetContext(ctx, contextService, workspaceService)
			namespaceService.SetContext(ctx, contextService)
			app.Startup(ctx)
		},
		Bind: []interface{}{
			contextService,
			workspaceService,
			podService,
			namespaceService,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func changeTheme(a *backend.App, data *menu.CallbackData) {
	runtime.EventsEmit(a.Ctx, "change-theme")
}

func startServer() {
	//server := gin.Default()
	//server.Use(cors.Default())
	//
	//server.GET("/", func(context *gin.Context) {
	//	context.JSON(http.StatusCreated, gin.H{
	//		"message": "hi there",
	//	})
	//})
	////routes.Routes(*app)
	//
	//server.Run(":3001")
}
