package backend

import (
	"context"
	"gube/backend/db"
	"gube/backend/models"
	"gube/backend/services"
	"k8s.io/client-go/tools/clientcmd/api"
	"os/exec"
)

// App struct
type App struct {
	Ctx    context.Context
	config *api.Config
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (app *App) Startup(ctx context.Context) {
	app.Ctx = ctx
	db.InitDB()
	err := models.SetDefaultWorkspace()
	if err != nil {
		panic(err.Error())
	}
	app.config = services.GetConfig()
}

func (app *App) GetActiveWorkspace() models.GenerictResult[models.Workspace] {
	return services.GetActiveWorkspace()
}

func (app *App) GetContextList() models.GenerictResult[[]*api.Context] {
	return services.GetContextList(app.config)
}

func (app *App) RunCommand(command string) models.GenerictResult[string] {
	//cmd := exec.Command("powershell", "/c", "Get-ChildItem")
	cmd := exec.Command("powershell", "/c", command)
	out, err := cmd.Output()
	if err != nil {
		result := models.GenerictResult[string]{ErrorMessage: err.Error()}
		return result
	}
	result := models.GenerictResult[string]{Data: string(out)}
	return result
}
