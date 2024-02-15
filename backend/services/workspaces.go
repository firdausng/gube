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
