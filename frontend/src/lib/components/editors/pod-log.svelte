<script lang="ts">
    import { codeMirrorEditor } from './editor';
    import {onDestroy, onMount} from 'svelte';
    import {type AppData, appDataStore} from "$lib/store/app-data-store";
    import { StreamPodLog } from "$lib/wailsjs/go/services/PodService"
    import {EventsOn} from "$lib/wailsjs/runtime";
    import {appEditorDataStore} from "$lib/store/app-editor-data";
    import {bottomTabDataStore} from "$lib/store/app-tab-data";


    let tabId:string
    let workspaceName = '';
    let contextName = '';
    let appData: AppData;
    appDataStore.subscribe(d =>{
        if(appData.activeWorkspace.activeContext){
            appData = data;
            workspaceName = appData.activeWorkspace.name;
            contextName = appData.activeWorkspace.activeContext!.name;
            tabId = `${d.activeWorkspace.name}-${d.activeWorkspace.activeContext?.name}`
        }

    })

    let namespaceName = '';
    let podName = '';
    let id = ``;
    const appDataStoreSubscriber = bottomTabDataStore.subscribe(d =>{
        if(d[tabId] && d[tabId].activeTab){
            namespaceName = d[tabId].activeTab!.resourceNamespace;
            podName = d[tabId].activeTab!.resourceName;
            id = `${tabId}-${namespaceName}-${podName}`;
        }
    })

    let data = {
        value: ''
    }

    appEditorDataStore.subscribe(d =>{
        if(d[id]){
            for (let i = 0; i < d[id].line.length; i++) {
                data.value += d[id].line[i] + "\n"
            }
        }
    })


    onMount(async()=>{
        if(id.length > 0){
            StreamPodLog(workspaceName, contextName, namespaceName, podName, "");
            let eventName = `EmitPodLog-${id}`;
            console.log('get log', eventName)

            EventsOn(eventName, (emitPodLog)=>{
                appEditorDataStore.update(d =>{
                    if(!d[id]){
                        d[id] = {
                            workspaceName,
                            contextName,
                            namespaceName,
                            resourceName: podName,
                            settings: {},
                            line:[emitPodLog.line]
                        }
                    }else {
                        d[id].line.push(emitPodLog.line)
                    }

                    console.log(d[id].line + "\n")
                    return d;
                })
            })
        }

    })

    onDestroy(()=>{
        appDataStoreSubscriber()
    })
</script>

<div class="">
    <div use:codeMirrorEditor={data}></div>
</div>
