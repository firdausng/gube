<script>
	import "../app.css";
	import ThemeToggle from "$lib/components/theme.svelte";

    import { appDataStore } from '$lib/store/app-data-store';
    import {GetActiveWorkspace} from "$lib/wailsjs/go/backend/App.js";
    import {goto} from "$app/navigation";
    import {onMount} from "svelte";

    onMount(async ()=>{
        const activeWorkspace = await GetActiveWorkspace()
        appDataStore.update(d =>{
            return {
                activeWorkspace: {
                    id: activeWorkspace.data.ID,
                    name: activeWorkspace.data.name,
                    contexts: []
                }
            }
        })
        if(activeWorkspace.data){
            await goto(`/workspaces/${activeWorkspace.data.ID}`)
            return activeWorkspace.data;
        }
        throw new Error(activeWorkspace);
    })

</script>

<ThemeToggle />
<slot />

