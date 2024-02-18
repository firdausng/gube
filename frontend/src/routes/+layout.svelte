<script lang="ts">
	import "../app.css";
	import ThemeToggle from "$lib/components/theme.svelte";
    import { appDataStore } from '$lib/store/app-data-store';
    import {GetActiveWorkspace, GetWorkspaceList} from "$lib/wailsjs/go/backend/App";
    import {goto} from "$app/navigation";
    import {onMount} from "svelte";

    onMount(async ()=>{
        const listPromise = GetWorkspaceList();
        const activeWorkspaceDataPromise = GetActiveWorkspace()
        const [list, activeWorkspaceData] = await Promise.all([listPromise, activeWorkspaceDataPromise])
        console.log(list, activeWorkspaceData)

        appDataStore.update(d =>{
            const activeWorkspace = {
                id: activeWorkspaceData.data.ID,
                name: activeWorkspaceData.data.name,
                // activeResource: null,
                contexts: [],
                activeContext: null
            }

            return {
                activeWorkspace,
                workspaces: list
            }
        })

        if(list.data.length > 1){
            await goto(`/workspaces`)
        }else{
            if(activeWorkspaceData.data){
                await goto(`/workspaces/${activeWorkspaceData.data.ID}`)
            }
        }
    })

</script>

<ThemeToggle />
<slot />

