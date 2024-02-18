package services

import (
	"gube/backend/models"
	"k8s.io/client-go/tools/clientcmd/api"
)

type NamespaceService struct {
	contextService *ContextService
}

func newNamespaceService(contextService *ContextService) *NamespaceService {
	container := &NamespaceService{
		contextService: contextService,
	}
	return container
}

func (service NamespaceService) GetNamespaceList() models.GenerictResult[[]*api.Context] {
	var contexts []*api.Context
	for _, context := range service.contextService.Config.Contexts {
		contexts = append(contexts, context)
	}
	result := models.GenerictResult[[]*api.Context]{Data: contexts}
	return result
}
