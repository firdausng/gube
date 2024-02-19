package services

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gube/backend/models"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"sync"
)

type PodService struct {
	ctx              context.Context
	contextService   *ContextService
	workspaceService *WorkspaceService
	cacheMutex       sync.Mutex
	eventCache       map[string]bool
}

func NewPodService() *PodService {
	return &PodService{}
}

func (service *PodService) SetContext(ctx context.Context, contextService *ContextService, workspaceService *WorkspaceService) {
	service.contextService = contextService
	service.workspaceService = workspaceService
	service.ctx = ctx
}

func (podService *PodService) GetPodList(workspaceId, contextName, namespaceName string) models.GenerictResult[[]corev1.Pod] {
	if len(namespaceName) == 0 {
		namespaceName = "default"
	}
	client, err := podService.contextService.GetContextClient(contextName)
	if err != nil {
		result := models.GenerictResult[[]corev1.Pod]{ErrorMessage: err.Error()}
		log.Printf("Error getting pod list %s\n", result)
	}
	pods, err := client.CoreV1().Pods(namespaceName).List(context.TODO(), metav1.ListOptions{})
	result := models.GenerictResult[[]corev1.Pod]{Data: pods.Items}
	return result
}

func (podService *PodService) StreamPods(workspaceId, contextName, namespaceName string) {
	cacheKey := fmt.Sprintf("EmitPodList-%s-%s-%s", workspaceId, contextName, namespaceName)

	podService.cacheMutex.Lock()
	if podService.eventCache[cacheKey] {
		podService.cacheMutex.Unlock()
		return
	}
	podService.eventCache[cacheKey] = true
	podService.cacheMutex.Unlock()

	client, err := podService.contextService.GetContextClient(contextName)
	if err != nil {
		runtime.LogError(podService.ctx, err.Error())
		result := models.GenerictResult[[]corev1.Pod]{ErrorMessage: err.Error()}
		log.Printf("Error getting pod list %s\n", result)
	}
	podWatcher, err := client.CoreV1().Pods(namespaceName).Watch(context.Background(), metav1.ListOptions{})

	// create a channel to stop the watch when needed
	stop := make(chan struct{})
	// handle watch events in a separate go routine
	go func() {
		for {
			select {
			case event := <-podWatcher.ResultChan():
				pod, ok := event.Object.(*corev1.Pod)
				if !ok {
					// not a pod - continue
					continue
				}
				// handle the pod instance as required
				runtime.EventsEmit(podService.ctx, cacheKey, pod)
			case <-stop:
				runtime.EventsOff(podService.ctx, cacheKey)
				podWatcher.Stop()
				return
			}
		}
	}()
}
