<script lang="ts">
	import CodemirrorEditor from "$lib/components/editors/codemirror-editor.svelte";
	import XtermTerminal from "$lib/components/terminal/xterm-terminal.svelte";
	import PodLogEditor from "$lib/components/editors/pod-log.svelte";
	import type { SvelteComponent } from 'svelte';
	import {appDataStore} from "$lib/store/app-data-store";
	import {bottomTabDataStore, type TabItem} from "$lib/store/app-tab-data";


	let tabId:string
	appDataStore.subscribe(d =>{
		tabId = `${d.activeWorkspace.name}-${d.activeWorkspace.activeContext?.name}`
	})

	let activeTab: TabItem|null|undefined;
	let tabData: TabItem[] = []

	bottomTabDataStore.subscribe(d => {
		if(d[tabId]){
			activeTab = d[tabId].activeTab
			tabData = d[tabId].tabs

			if(activeTab){
				changeTab(activeTab!)
			}
		}
	})

	let activeComponent: typeof SvelteComponent<any>

	function changeTab(tab: TabItem) {
		activeTab = tab
		if(activeTab.component.toLowerCase() === "log"){
			activeComponent = PodLogEditor;
		}
		if(activeTab.component.toLowerCase() === "terminal"){
			activeComponent = XtermTerminal;
		}
	}
</script>

<footer class="w-full">
	<div class="bg-app-lightest dark:bg-app-darkest">
		<div class="flex justify-items-start items-stretch w-full">
			<ul class="-mb-px flex gap-4 text-sm font-medium bg-app-light dark:bg-app-dark">
				{#each tabData as tab (tab.name)}
					<li class="flex-1">
						<button
							on:click={()=> changeTab(tab)}
							class="relative flex items-center justify-center gap-2 px-1 py-1 hover:text-app-dark dark:hover:text-app-light text-xs
						{activeTab?.name === tab.name ?
						'text-app-darkest dark:text-app-lightest after:absolute after:left-0 after:bottom-0 after:h-0.5 after:w-full after:bg-app-light dark:after:bg-app-dark':
						'text-app-darkest/50 dark:text-app-lightest/50'
						}"
						>{tab.name}</button>
					</li>
				{/each}
			</ul>
		</div>
		<div>
			{#each tabData as tab (tab.name)}
				{#if activeTab?.name === tab.name}
					<svelte:component this={activeComponent} />
				{/if}
			{/each}
		</div>
	</div>
</footer>