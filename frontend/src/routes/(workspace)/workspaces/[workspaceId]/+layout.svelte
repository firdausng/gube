<script lang="ts">
	import ContextSidenav from "$lib/components/context-sidenav.svelte";
	import BottomNav from "$lib/components/bottom-nav.svelte";
	import { appSettingStore } from '$lib/store/app-setting-store';
	import { PaneGroup, Pane, PaneResizer, type PaneAPI } from "paneforge";
	import {type AppData, appDataStore} from '$lib/store/app-data-store';

	let bottomSize = 25;
	appDataStore.subscribe(data =>{
		if(data.activeWorkspace.activeContext){
			if(data.activeWorkspace.activeContext.tabData.activeTab){
				bottomSize = 50;
			}
		}

	})


</script>

<div class="flex gap-1 h-full bg-app-light dark:bg-app-dark">
	{#if !$appSettingStore.leftSidebarHidden}
		<ContextSidenav />
	{/if}

	<PaneGroup direction="vertical">
		<Pane class="bg-app-lightest dark:bg-app-darkest">
			<slot />
		</Pane>
		<PaneResizer class="relative flex items-center justify-center w-full bg-app-light/50 dark:bg-app-dark/50 shadow-inner shadow-app-light dark:shadow-app-dark border-y-2 border-primary-500/50">
			<div class="z-10 flex items-center justify-center text-xs">
				<i class="ri-more-line"></i>
			</div>
		</PaneResizer>
		<Pane bind:defaultSize={bottomSize}>
			<div class="w-full h-full overflow-auto bg-app-lightest dark:bg-app-darkest">
				<BottomNav />
			</div>
		</Pane>
	</PaneGroup>
</div>

