<script lang="ts">
    import {appDataStore, type Tab, type TabItem} from "$lib/store/app-data-store";
    import { Button, Dropdown, DropdownItem, ToolbarButton, DropdownDivider, Modal } from 'flowbite-svelte';
    // import { fly } from "svelte/transition";
    // import { preloadData, pushState, goto } from '$app/navigation';
    // import { page } from '$app/stores';
    import type {Action} from "./PodSectionType";

    export let podName:string;
    let deletePodModal = false;

    let options: Action[] = [
        {type:'tab', name:'Log', icon:"ri-align-justify"},
        {type:'tab', name:'Terminal', icon:"ri-terminal-box-line"},
        {type:'tab', name:'edit', icon:"ri-edit-line"},
        {type:'modal', name:'delete', icon:"ri-delete-bin-6-line"},
    ]

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
</script>

<i class="ri-more-2-line dots-menu"></i>
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
<!--    <DropdownItem slot="footer">Sign out</DropdownItem>-->
</Dropdown>

<Modal title="Delete Pod" bind:open={deletePodModal} autoclose>
    <p class="text-base leading-relaxed text-gray-500 dark:text-gray-400">Are you sure to delete pod {podName}</p>
    <svelte:fragment slot="footer">
        <Button on:click={() => alert('Handle "success"')}>OK</Button>
        <Button color="alternative">Cancel</Button>
    </svelte:fragment>
</Modal>

