package backend

import (
	"context"
	"gube/backend/db"
	"gube/backend/models"
	"sync"
)

// App struct
type App struct {
	Ctx context.Context
	//service    *services.AppService
	cacheMutex sync.Mutex
	eventCache map[string]bool
	Workspace  models.Workspace
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (app *App) Startup(ctx context.Context) {
	app.Ctx = ctx
	db.InitDB()
	workspace, err := models.SetDefaultWorkspace()
	if err != nil {
		panic(err.Error())
	}
	app.Workspace = workspace
	//app.service = services.NewServiceContainer()
}

//func (app *App) GetActiveWorkspace() models.GenerictResult[models.Workspace] {
//	return services.GetActiveWorkspace()
//}
//
//func (app *App) GetWorkspaceList() models.GenerictResult[[]models.Workspace] {
//	return services.GetAllWorkspace()
//}

//func (app *App) GetWorkspaceList() models.GenerictResult[[]models.Workspace] {
//	return services.GetAllWorkspace()
//}

//func (app *App) GetNamespaces(contextName string) {
//	clientConfig := clientcmd.NewNonInteractiveClientConfig(*app.config, contextName, nil, nil)
//	restConfig, _ := clientConfig.ClientConfig()
//
//	clientset, err := kubernetes.NewForConfig(restConfig)
//	if err != nil {
//		runtime.LogError(app.Ctx, err.Error())
//		result := models.GenerictResult[[]corev1.Pod]{ErrorMessage: err.Error()}
//		log.Printf("Error getting pod list %s\n", result)
//	}
//	namespaceList, err := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
//}
//
//func (app *App) GetContextList() models.GenerictResult[[]*api.Context] {
//	return app.service.Context.GetContextList()
//}
//
//func (app *App) RunCommand(command string) models.GenerictResult[string] {
//	//cmd := exec.Command("powershell", "/c", "Get-ChildItem")
//	cmd := exec.Command("powershell", "/c", command)
//	out, err := cmd.Output()
//	if err != nil {
//		result := models.GenerictResult[string]{ErrorMessage: err.Error()}
//		return result
//	}
//	result := models.GenerictResult[string]{Data: string(out)}
//	return result
//}
//
//func (app *App) GetPodList(contextName string, namespaceName string) models.GenerictResult[[]corev1.Pod] {
//	return app.service.Pod.GetPodList(contextName, namespaceName)
//}
//
//func (app *App) StreamPods(contextName string, namespaceName string) {
//	cacheKey := fmt.Sprintf("EmitPodList-%s-%s-%s", app.Workspace.ID, contextName, namespaceName)
//
//	app.cacheMutex.Lock()
//	if app.eventCache[cacheKey] {
//		app.cacheMutex.Unlock()
//		return
//	}
//	app.eventCache[cacheKey] = true
//	app.cacheMutex.Unlock()
//
//	client, err := app.service.Context.GetContextClient(contextName)
//	if err != nil {
//		runtime.LogError(app.Ctx, err.Error())
//		result := models.GenerictResult[[]corev1.Pod]{ErrorMessage: err.Error()}
//		log.Printf("Error getting pod list %s\n", result)
//	}
//	podWatcher, err := client.CoreV1().Pods(namespaceName).Watch(context.Background(), metav1.ListOptions{})
//
//	// create a channel to stop the watch when needed
//	stop := make(chan struct{})
//	// handle watch events in a separate go routine
//	go func() {
//		for {
//			select {
//			case event := <-podWatcher.ResultChan():
//				pod, ok := event.Object.(*corev1.Pod)
//				if !ok {
//					// not a pod - continue
//					continue
//				}
//				// handle the pod instance as required
//				runtime.EventsEmit(app.Ctx, cacheKey, pod)
//			case <-stop:
//				runtime.EventsOff(app.Ctx, cacheKey)
//				podWatcher.Stop()
//				return
//			}
//		}
//	}()
//}
//
//func (app *App) StreamLog(contextName string, namespaceName string, podName string) {
//	podLogs, err := app.service.Pod.StreamPodLogs(contextName, namespaceName, podName)
//	if err != nil {
//		result := models.GenerictResult[string]{ErrorMessage: "error2: namespace:" + namespaceName + "pd:" + podName + err.Error()}
//		log.Printf("Error getting deployment logs %s\n", result)
//	}
//	scanner := bufio.NewScanner(podLogs)
//	for scanner.Scan() {
//		// Read line by line
//		line := scanner.Text()
//		log.Printf("line %s\n", line)
//		runtime.EventsEmit(app.Ctx, "PodLog", line)
//	}
//}
