package services

import (
	"context"
	"gube/backend/models"
)

type WorkspaceService struct {
	ctx context.Context
}

func NewWorkspaceService() *WorkspaceService {
	return &WorkspaceService{}
}

func (service *WorkspaceService) SetContext(ctx context.Context) {
	service.ctx = ctx
}

func newWorkspaceService(ctx context.Context, contextService *ContextService) *WorkspaceService {
	container := &WorkspaceService{
		ctx: ctx,
		//contextService: contextService,
	}
	return container
}

func (workspace *WorkspaceService) GetActiveWorkspace() models.GenerictResult[models.Workspace] {
	activeWorkspace, err := models.GetActiveWorkspace()
	if err != nil {
		result := models.GenerictResult[models.Workspace]{ErrorMessage: err.Error()}
		return result
	}
	result := models.GenerictResult[models.Workspace]{Data: *activeWorkspace}
	return result
}

func (workspace *WorkspaceService) GetAllWorkspace() models.GenerictResult[[]models.Workspace] {
	list, err := models.GetAllWorkspaces()
	if err != nil {
		result := models.GenerictResult[[]models.Workspace]{ErrorMessage: err.Error()}
		return result
	}
	result := models.GenerictResult[[]models.Workspace]{Data: list}
	return result
}
