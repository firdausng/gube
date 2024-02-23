<script lang="ts">
    import {appDataStore, type Tab, type TabItem} from "$lib/store/app-data-store";

    export let data:string;
    let selected = '';
    let options = [
        '',
        'Log',
        'Terminal',
    ]

    $: if(selected){
        appDataStore.update(d =>{
            if(d.activeWorkspace.activeContext){
                const tabItem: TabItem = {
                    component: selected,
                    resourceName: data,
                    name: `${selected}: ${data}`
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
            }
            return d;
        })
    }


</script>
<select class="rounded px-1 border-2 text-app-darkest focus:outline-none appearance-none" bind:value={selected}>
    {#each options as value}<option {value}>{value}</option>{/each}
</select>

