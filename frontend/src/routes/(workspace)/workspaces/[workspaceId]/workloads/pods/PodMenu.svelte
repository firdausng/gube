<script lang="ts">
    import {type AppData, appDataStore, type Tab, type TabItem} from "$lib/store/app-data-store";
    import { Button, Dropdown, DropdownItem, ToolbarButton, DropdownDivider, Modal } from 'flowbite-svelte';
    import { DeletePod } from "$lib/wailsjs/go/services/PodService"
    import type {Action} from "./PodSectionType";

    export let podName:string;
    export let namespace: string;
    $: deletePodModal = false;

    let options: Action[] = [
        {type:'tab', name:'Log', icon:"ri-align-justify"},
        {type:'tab', name:'Terminal', icon:"ri-terminal-box-line"},
        {type:'tab', name:'edit', icon:"ri-edit-line"},
        {type:'modal', name:'delete', icon:"ri-delete-bin-6-line"},
    ]

    let appData: AppData;
    appDataStore.subscribe(data =>{
        appData = data;
    })
    $: id = `${namespace}-${podName}-modal`;

    function onActionSelected(action: Action){
        appDataStore.update(d =>{
            if(d.activeWorkspace.activeContext){
                switch (action.type){
                    case 'tab':{
                        const tabItem: TabItem = {
                            component: action.name,
                            resourceName: podName,
                            resourceNamespace: namespace,
                            name: `${action.name}: ${podName}`
                        }

                        if( !d.activeWorkspace.activeContext.tabData){
                            d.activeWorkspace.activeContext.tabData = {
                                activeTab: tabItem,
                                tabs: []
                            }
                        }
                        d.activeWorkspace.activeContext.tabData.activeTab = tabItem

                        let foundItem = d.activeWorkspace.activeContext.tabData.tabs.find(item => item.name === tabItem.name);

                        if(foundItem) {
                            foundItem = {...foundItem, ...tabItem};
                        } else {
                            d.activeWorkspace.activeContext.tabData.tabs = [...d.activeWorkspace.activeContext.tabData.tabs, tabItem];
                        }
                        break;
                    }
                    case 'modal':{
                        if(action.name === "delete"){
                            deletePodModal = true;
                        }
                        break;
                    }
                    default:
                        break;
                }
            }
            return d;
        })
    }

    async function deletePod(namespace:string, podName:string){
        if(appData.activeWorkspace.activeContext){
            await DeletePod(appData.activeWorkspace.name, appData.activeWorkspace.activeContext?.name, namespace, podName);
        }
    }
</script>

<i class="ri-more-2-line dots-menu {id}"></i>
<Dropdown triggeredBy=".{id}">
    {#each options as value (value.name)}
        <DropdownItem on:click={()=>onActionSelected(value)}>
            <i class="mr-2 font-semibold {value.icon}"></i>
            {value.name}
        </DropdownItem>
    {/each}
</Dropdown>

<Modal {id} title="Delete Pod" bind:open={deletePodModal} autoclose>
    <p class="text-base leading-relaxed text-gray-500 dark:text-gray-400">Are you sure to delete pod {podName}</p>
    <svelte:fragment slot="footer">
        <Button on:click={()=> deletePod(namespace, podName)}>OK</Button>
        <Button color="alternative">Cancel</Button>
    </svelte:fragment>
</Modal>