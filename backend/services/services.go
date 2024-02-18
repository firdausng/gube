package services

type AppService struct {
	Context   *ContextService
	Pod       *PodService
	Namespace *NamespaceService
}

func NewServiceContainer() *AppService {
	contextService := newContextService()
	container := &AppService{
		Context:   contextService,
		Pod:       newPodService(contextService),
		Namespace: newNamespaceService(contextService),
	}
	return container
}
