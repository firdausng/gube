package services

import (
	"bufio"
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
	podLog           map[string][]PodLog
}

type PodLog struct {
	Context   string `json:"context"`
	Namespace string `json:"namespace"`
	Pod       string `json:"pod"`
	Line      string `json:"line"`
}

func NewPodService() *PodService {
	return &PodService{
		eventCache: map[string]bool{},
	}
}

func (podService *PodService) SetContext(ctx context.Context, contextService *ContextService, workspaceService *WorkspaceService) {
	podService.contextService = contextService
	podService.workspaceService = workspaceService
	podService.ctx = ctx
}

func (podService *PodService) GetPodList(workspaceId, contextName, namespaceName string) models.GenerictResult[[]corev1.Pod] {
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
	log.Printf("EmitPodList-%s-%s-%s", workspaceId, contextName, namespaceName)

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
				_, ok := event.Object.(*corev1.Pod)
				if !ok {
					// not a pod - continue
					continue
				}
				// handle the pod instance as required
				runtime.EventsEmit(podService.ctx, cacheKey, event)
			case <-stop:
				runtime.EventsOff(podService.ctx, cacheKey)
				podWatcher.Stop()
				return
			}
		}
	}()
}

func (podService *PodService) StreamPodLog(workspaceId, contextName, namespaceName, podName string) {
	cacheKey := fmt.Sprintf("EmitPodLog-%s-%s-%s-%s", workspaceId, contextName, namespaceName, podName)
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
		log.Printf("Error getting context client %s\n", result)
	}

	podLogOpts := corev1.PodLogOptions{
		Follow:     true,
		Timestamps: true,
	}
	req := client.CoreV1().Pods(namespaceName).GetLogs(podName, &podLogOpts)
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		result := models.GenerictResult[string]{ErrorMessage: "error2: namespace:" + namespaceName + "pd:" + podName + err.Error()}
		log.Printf("Error getting deployment logs %s\n", result)
	}
	//defer podLogs.Close()
	//
	podLogList := make([]PodLog, 0)
	scanner := bufio.NewScanner(podLogs)
	for scanner.Scan() {
		line := scanner.Text()
		log.Printf("line %s\n", line)

		podLog := PodLog{
			Context:   contextName,
			Namespace: namespaceName,
			Pod:       podName,
			Line:      line,
		}
		podService.podLog[cacheKey] = append(podLogList, podLog)

		// Emit the JSON
		runtime.EventsEmit(podService.ctx, cacheKey, podLog)
	}
}

func (podService *PodService) DeletePod(workspaceId, contextName, namespaceName, podName string) {
	client, err := podService.contextService.GetContextClient(contextName)
	if err != nil {
		runtime.LogError(podService.ctx, err.Error())
		result := models.GenerictResult[[]corev1.Pod]{ErrorMessage: err.Error()}
		log.Printf("Error getting context client %s\n", result)
	}

	podLogOpts := metav1.DeleteOptions{}
	err = client.CoreV1().Pods(namespaceName).Delete(podService.contextService.ctx, podName, podLogOpts)
	if err != nil {
		runtime.LogError(podService.ctx, err.Error())
		log.Printf("Error delete pod %s\n", err)
	}
}
