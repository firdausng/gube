<script lang="ts">
    import {appDataStore, type Tab, type TabItem} from "$lib/store/app-data-store";
    import {Button, DropdownMenu} from "bits-ui";
    import { fly } from "svelte/transition";

    export let data:string;
    type Action = {
        type:'tab'|'modal',
        name:string,
        icon: string
    }

    let options: Action[] = [
        {type:'tab', name:'Log', icon:"ri-align-justify"},
        {type:'tab', name:'Terminal', icon:"ri-terminal-box-line"},
        {type:'modal', name:'delete', icon:"ri-delete-bin-6-line"},
    ]

    function onActionSelected(action: Action){
        appDataStore.update(d =>{
            if(d.activeWorkspace.activeContext){
                switch (action.type){
                    case 'tab':{
                        const tabItem: TabItem = {
                            component: action.name,
                            resourceName: data,
                            name: `${action.type}: ${action}`
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
                        console.log('delete pod')
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
<DropdownMenu.Root>
    <DropdownMenu.Trigger class="focus-visible inline-flex h-8 w-8 items-center justify-center rounded-full text-sm font-medium text-foreground shadow-btn active:scale-98">
        <i class="ri-more-2-line"></i>
    </DropdownMenu.Trigger>

    <DropdownMenu.Content
            class="w-full max-w-[120px] rounded-lg border-2 border-app-lightest dark:border-app-darkest bg-app-light dark:bg-app-dark px-1 shadow-popover cursor-pointer"
            transition={fly}
            sideOffset={8}
    >
        <DropdownMenu.Label />
        {#each options as value (value.name)}
            <DropdownMenu.Item
                    class="flex select-none items-center rounded-button py-1 pl-3 pr-1.5 text-sm font-medium !ring-0 !ring-transparent data-[highlighted]:bg-muted"
            >
                <button class="flex items-center capitalize" on:click={()=>onActionSelected(value)}>
                    <!--                <UserCircle class="mr-2 size-5 text-foreground-alt" />-->
                    <i class="mr-2 size-5 text-foreground-alt {value.icon}"></i>
                    {value.name}
                </button>
            </DropdownMenu.Item>
        {/each}
    </DropdownMenu.Content>
</DropdownMenu.Root>

