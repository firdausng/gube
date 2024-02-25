<script lang="ts">
    import {type AppData, appDataStore, type Tab, type TabItem} from "$lib/store/app-data-store";
    import { Button, Dropdown, DropdownItem, ToolbarButton, DropdownDivider, Modal } from 'flowbite-svelte';
    import {DeletePod} from "$lib/wailsjs/go/services/PodService"
    // import { fly } from "svelte/transition";
    // import { preloadData, pushState, goto } from '$app/navigation';
    // import { page } from '$app/stores';
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

    let modalOptions: Action[] = [
        {type:'modal', name:'delete', icon:"ri-delete-bin-6-line"},
    ]

    function onActionSelected(action: Action){
        appDataStore.update(d =>{
            if(d.activeWorkspace.activeContext){
                switch (action.type){
                    case 'tab':{
                        const tabItem: TabItem = {
                            component: action.name,
                            resourceName: podName,
                            name: `${action.type}: ${podName}`
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
                            // id = `${namespace}-${podName}-modal`
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
        console.log('deleting pod: ',namespace, podName )
        if(appData.activeWorkspace.activeContext){
            await DeletePod(appData.activeWorkspace.name, appData.activeWorkspace.activeContext?.name, namespace, podName);
        }

    }
</script>

<i class="ri-more-2-line dots-menu"></i>
<!--<p>{podName} {namespace}</p>-->
<Dropdown triggeredBy=".dots-menu">
    {#each options as value (value.name)}
        <DropdownItem on:click={()=>onActionSelected(value)}>
            <i class="mr-2 font-semibold {value.icon}"></i>
            {value.name}
        </DropdownItem>
    {/each}
<!--    <DropdownItem>Dashboard</DropdownItem>-->
<!--    <DropdownItem>Settings</DropdownItem>-->
<!--    <DropdownItem>Earnings</DropdownItem>-->
<!--    <DropdownItem slot="footer">-->
<!--        <button on:click={()=>deletePod(namespace, podName)}>-->
<!--            Delete-->
<!--        </button>-->
<!--    </DropdownItem>-->
</Dropdown>

<Modal {id} title="Delete Pod" bind:open={deletePodModal} autoclose>
    <p class="text-base leading-relaxed text-gray-500 dark:text-gray-400">Are you sure to delete pod {id}</p>
    <svelte:fragment slot="footer">
        <Button on:click={()=> deletePod(podName, namespace)}>OK</Button>
        <Button color="alternative">Cancel</Button>
    </svelte:fragment>
</Modal>

