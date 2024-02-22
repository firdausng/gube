package services

import (
	"context"
	"flag"
	"fmt"
	"gube/backend/models"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

type ContextService struct {
	ctx           context.Context
	Config        *api.Config
	contextClient map[string]*kubernetes.Clientset
}

func NewContextService() *ContextService {
	return &ContextService{}
}

func (service *ContextService) SetContext(ctx context.Context) {
	config := getConfig()
	service.Config = config
	service.contextClient = make(map[string]*kubernetes.Clientset)
	service.ctx = ctx
}

func (service *ContextService) GetContextClient(contextName string) (*kubernetes.Clientset, error) {
	if _, ok := service.contextClient[contextName]; ok {
		fmt.Printf("Client already exists for this context %s \n", contextName)
	} else {
		clientConfig := clientcmd.NewNonInteractiveClientConfig(*service.Config, contextName, nil, nil)
		restConfig, _ := clientConfig.ClientConfig()

		client, err := kubernetes.NewForConfig(restConfig)
		if err != nil {
			return nil, err
		}
		service.contextClient[contextName] = client
	}
	return service.contextClient[contextName], nil
}

func (service *ContextService) GetContextList() models.GenerictResult[[]*api.Context] {
	var contexts []*api.Context
	for _, context := range service.Config.Contexts {
		contexts = append(contexts, context)
	}
	result := models.GenerictResult[[]*api.Context]{Data: contexts}
	return result
}

func getConfig() *api.Config {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err1 := clientcmd.LoadFromFile(*kubeconfig)
	if err1 != nil {
		panic(err1.Error())
	}
	return config
}
