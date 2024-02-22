package services

import (
	"gube/backend/models"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/release"
)

func GetHelmReleaseList() models.GenerictResult[[]*release.Release] {
	actionConfig := new(action.Configuration)
	listAction := action.NewList(actionConfig)
	listAction.All = true

	releases, err := listAction.Run()
	if err != nil {
		panic(err)
	}
	result := models.GenerictResult[[]*release.Release]{Data: releases}
	return result
}

//func GetHelmRelease(releaseNamespace string) {
//	actionConfig := new(action.Configuration)
//
//	if err := actionConfig.Init(kube.GetConfig(*kubeconfig, "", releaseNamespace), releaseNamespace, os.Getenv("HELM_DRIVER"), func(format string, v ...interface{}) {
//		fmt.Sprintf(format, v)
//	}); err != nil {
//		panic(err)
//	}
//}
