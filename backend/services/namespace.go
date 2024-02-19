package services

import (
	"context"
	"gube/backend/models"
	"k8s.io/client-go/tools/clientcmd/api"
)

type NamespaceService struct {
	ctx            context.Context
	contextService *ContextService
}

func NewNamespaceService() *NamespaceService {
	return &NamespaceService{}
}

func (service *NamespaceService) SetContext(ctx context.Context, contextService *ContextService) {
	service.contextService = contextService
	service.ctx = ctx
}

func (service NamespaceService) GetNamespaceList() models.GenerictResult[[]*api.Context] {
	var contexts []*api.Context
	for _, context := range service.contextService.Config.Contexts {
		contexts = append(contexts, context)
	}
	result := models.GenerictResult[[]*api.Context]{Data: contexts}
	return result
}
