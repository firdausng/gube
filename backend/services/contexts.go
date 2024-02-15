package services

import (
	"flag"
	"gube/backend/models"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func GetContextList(config *api.Config) models.GenerictResult[[]*api.Context] {
	var contexts []*api.Context
	for _, context := range config.Contexts {
		contexts = append(contexts, context)
	}
	result := models.GenerictResult[[]*api.Context]{Data: contexts}
	return result
}

func GetConfig() *api.Config {
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
