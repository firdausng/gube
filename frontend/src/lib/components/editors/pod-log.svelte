<script lang="ts">
    import { codeMirrorEditor } from './editor';
    import { onMount } from 'svelte';
    import {type AppData, appDataStore, type Tab, type TabItem} from "$lib/store/app-data-store";
    import { StreamPodLog } from "$lib/wailsjs/go/services/PodService"
    import {EventsOn} from "$lib/wailsjs/runtime";

    let appData: AppData;
    appDataStore.subscribe(data =>{
        appData = data;
    })

    let data = {
        value: ''
    }

    onMount(async()=>{
        if(appData.activeWorkspace.activeContext && appData.activeWorkspace.activeContext?.tabData.activeTab){

            let workspaceId = appData.activeWorkspace.name;
            let contextName = appData.activeWorkspace.activeContext.name;
            let namespaceName = appData.activeWorkspace.activeContext.tabData.activeTab.resourceNamespace;
            let podName = appData.activeWorkspace.activeContext.tabData.activeTab.resourceName;
            StreamPodLog(workspaceId, contextName, namespaceName, podName);
            let eventName = `EmitPodLog-${workspaceId}-${contextName}-${namespaceName}-${podName}`;
            console.log('get log', eventName)

            EventsOn(eventName, (emitPodLog)=>{
                console.log(emitPodLog)
                data = {...data, value: emitPodLog.line + "\n"}
                console.log(data.value)
                // data = data;
            })
        }

    })
</script>

<div class="">
    <div use:codeMirrorEditor={data}></div>
</div>
