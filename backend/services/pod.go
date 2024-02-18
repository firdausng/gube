package services

import (
	"context"
	"gube/backend/models"
	"io"
	corev1 "k8s.io/api/core/v1"
	"log"
)

type PodService struct {
	contextService *ContextService
}

func newPodService(contextService *ContextService) *PodService {
	container := &PodService{
		contextService: contextService,
	}
	return container
}

//func StreamPodList(config *api.Config, contextName string, namespaceName string) (*v1.Pod, error) {
//	clientConfig := clientcmd.NewNonInteractiveClientConfig(*config, contextName, nil, nil)
//	restConfig, _ := clientConfig.ClientConfig()
//
//	clientset, err := kubernetes.NewForConfig(restConfig)
//	if err != nil {
//		result := models.GenerictResult[[]corev1.Pod]{ErrorMessage: err.Error()}
//		log.Printf("Error getting pod list %s\n", result)
//	}
//	pods, err := clientset.CoreV1().Pods(namespaceName).Watch(context.TODO(), metav1.ListOptions{})
//	if err != nil {
//		result := models.GenerictResult[string]{ErrorMessage: "error2: namespace:" + namespaceName + err.Error()}
//		log.Printf("Error getting pods %s\n", result)
//	}
//
//	stop := make(chan struct{})
//	for {
//		select {
//		case event := <-pods.ResultChan():
//			pod, ok := event.Object.(*corev1.Pod)
//			if !ok {
//				// not a pod - continue
//				continue
//			}
//			// handle the pod instance as required
//			fmt.Printf("Received watch event: %v\n", pod.Name)
//			//runtime.EventsEmit(app.Ctx, "", pod)
//			return pod, nil
//		case <-stop:
//			pods.Stop()
//			return
//		}
//	}
//
//
//	return pods, nil
//}

func (podService PodService) StreamPodLogs(contextName string, namespaceName string, podName string) (io.ReadCloser, error) {
	client, err := podService.contextService.GetContextClient(contextName)
	if err != nil {
		result := models.GenerictResult[string]{ErrorMessage: "error1: " + err.Error()}
		log.Printf("Error getting deployment logs %s\n", result)
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
	return podLogs, nil
}
