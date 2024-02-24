<script lang="ts">
    import {appDataStore, type Tab, type TabItem} from "$lib/store/app-data-store";
    export let data:string;
    type Action = {
        type:'tab'|'modal',
        name:string
    }

    let selected: Action
    let options = [
        {type:'tab', name:'Log'},
        {type:'tab', name:'Terminal'},
        {type:'modal', name:'delete'},
    ]

    $: if(selected){
        appDataStore.update(d =>{
            if(d.activeWorkspace.activeContext){
                switch (selected.type){
                    case 'tab':{
                        const tabItem: TabItem = {
                            component: selected.name,
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
<select class="rounded px-1 border-2 text-app-darkest focus:outline-none appearance-none" bind:value={selected}>
    {#each options as value}<option {value}>{value.name}</option>{/each}
</select>

