package main

import (
	"context"
	"os/exec"
)

// App struct
type App struct {
	ctx context.Context
}

type GenerictResult[T any] struct {
	Data         T
	ErrorMessage string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) RunCommand(command string) GenerictResult[string] {
	//cmd := exec.Command("powershell", "/c", "Get-ChildItem")
	cmd := exec.Command("powershell", "/c", command)
	out, err := cmd.Output()
	if err != nil {
		result := GenerictResult[string]{ErrorMessage: err.Error()}
		return result
	}
	result := GenerictResult[string]{Data: string(out)}
	return result
}
