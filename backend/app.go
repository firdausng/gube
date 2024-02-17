package backend

import (
	"bufio"
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gube/backend/db"
	"gube/backend/models"
	"gube/backend/services"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"log"
	"os/exec"
	"sync"
)

// App struct
type App struct {
	Ctx        context.Context
	config     *api.Config
	cacheMutex sync.Mutex
	eventCache map[string]bool
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
	err := models.SetDefaultWorkspace()
	if err != nil {
		panic(err.Error())
	}
	app.config = services.GetConfig()
}

func (app *App) GetActiveWorkspace() models.GenerictResult[models.Workspace] {
	return services.GetActiveWorkspace()
}

func (app *App) GetContextList() models.GenerictResult[[]*api.Context] {
	return services.GetContextList(app.config)
}

func (app *App) RunCommand(command string) models.GenerictResult[string] {
	//cmd := exec.Command("powershell", "/c", "Get-ChildItem")
	cmd := exec.Command("powershell", "/c", command)
	out, err := cmd.Output()
	if err != nil {
		result := models.GenerictResult[string]{ErrorMessage: err.Error()}
		return result
	}
	result := models.GenerictResult[string]{Data: string(out)}
	return result
}

func (app *App) GetPodList(contextName string, namespaceName string) models.GenerictResult[[]corev1.Pod] {
	if len(namespaceName) == 0 {
		namespaceName = "default"
	}
	clientConfig := clientcmd.NewNonInteractiveClientConfig(*app.config, contextName, nil, nil)
	restConfig, _ := clientConfig.ClientConfig()

	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		runtime.LogError(app.Ctx, err.Error())
		//logger.Logger.Error(err.Error())
		//log.Printf("This is a %s message", err.Error())
		result := models.GenerictResult[[]corev1.Pod]{ErrorMessage: err.Error()}
		return result
	}
	pods, err := clientset.CoreV1().Pods(namespaceName).List(context.TODO(), metav1.ListOptions{})
	result := models.GenerictResult[[]corev1.Pod]{Data: pods.Items}
	return result
}

func (app *App) StreamPods(contextName string, namespaceName string) {
	cacheKey := fmt.Sprintf("EmitPodList-%s-%s", contextName, namespaceName)

	app.cacheMutex.Lock()
	if app.eventCache[cacheKey] {
		app.cacheMutex.Unlock()
		return
	}
	app.eventCache[cacheKey] = true
	app.cacheMutex.Unlock()

	clientConfig := clientcmd.NewNonInteractiveClientConfig(*app.config, contextName, nil, nil)
	restConfig, _ := clientConfig.ClientConfig()

	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		runtime.LogError(app.Ctx, err.Error())
		//logger.Logger.Error(err.Error())
		//log.Printf("This is a %s message", err.Error())
		result := models.GenerictResult[[]corev1.Pod]{ErrorMessage: err.Error()}
		log.Printf("Error getting pod list %s\n", result)
	}
	podWatcher, err := clientset.CoreV1().Pods(namespaceName).Watch(context.Background(), metav1.ListOptions{})

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
				runtime.EventsEmit(app.Ctx, "", pod)
			case <-stop:
				podWatcher.Stop()
				return
			}
		}
	}()
}

func (app *App) StreamLog(contextName string, namespaceName string, podName string) {
	podLogs, err := services.StreamPodLogs(app.config, contextName, namespaceName, podName)
	if err != nil {
		result := models.GenerictResult[string]{ErrorMessage: "error2: namespace:" + namespaceName + "pd:" + podName + err.Error()}
		log.Printf("Error getting deployment logs %s\n", result)
	}
	scanner := bufio.NewScanner(podLogs)
	for scanner.Scan() {
		// Read line by line
		line := scanner.Text()
		log.Printf("line %s\n", line)
		runtime.EventsEmit(app.Ctx, "PodLog", line)
	}
}
