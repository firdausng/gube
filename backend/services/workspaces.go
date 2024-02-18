package services

import (
	"gube/backend/models"
)

func GetActiveWorkspace() models.GenerictResult[models.Workspace] {
	activeWorkspace, err := models.GetActiveWorkspace()
	if err != nil {
		result := models.GenerictResult[models.Workspace]{ErrorMessage: err.Error()}
		return result
	}
	result := models.GenerictResult[models.Workspace]{Data: *activeWorkspace}
	return result
}

func GetAllWorkspace() models.GenerictResult[[]models.Workspace] {
	list, err := models.GetAllWorkspaces()
	if err != nil {
		result := models.GenerictResult[[]models.Workspace]{ErrorMessage: err.Error()}
		return result
	}
	result := models.GenerictResult[[]models.Workspace]{Data: list}
	return result
}
