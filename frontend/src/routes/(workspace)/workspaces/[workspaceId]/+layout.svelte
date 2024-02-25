<script lang="ts">
	import ContextSidenav from "$lib/components/context-sidenav.svelte";
	import BottomNav from "$lib/components/bottom-nav.svelte";
	import { appSettingStore } from '$lib/store/app-setting-store';
	// import { Pane, Splitpanes } from 'svelte-splitpanes';

	import { PaneGroup, Pane, PaneResizer, type PaneAPI } from "paneforge";

	let paneOne: PaneAPI;
	let collapsed = false;

</script>

<div class="flex gap-1 h-full bg-app-light dark:bg-app-dark">
	{#if !$appSettingStore.leftSidebarHidden}
		<ContextSidenav />
	{/if}

<!--	<Splitpanes class="" horizontal={true}>-->
<!--		<Pane class="bg-app-lightest dark:bg-app-darkest">-->
<!--			<slot />-->
<!--		</Pane>-->
<!--		<Pane minSize={1} size={1}>-->
<!--			<div class="w-full h-full overflow-auto">-->
<!--				<BottomNav />-->
<!--			</div>-->
<!--		</Pane>-->
<!--	</Splitpanes>-->
	<PaneGroup direction="vertical">
		<Pane class="bg-app-lightest dark:bg-app-darkest">
			<slot />
		</Pane>
		<PaneResizer class="relative flex items-center justify-center w-full bg-app-light/50 dark:bg-app-dark/50 shadow-inner shadow-app-light dark:shadow-app-dark border-y-2 border-primary-500/50">
			<div class="z-10 flex items-center justify-center text-xs">
<!--				<iconify-icon icon="solar:slider-horizontal-bold-duotone" class="px-1 text-xs text-app-dark dark:text-app-lightest text-center w-full"></iconify-icon>-->
				<i class="ri-more-line"></i>
			</div>
		</PaneResizer>
		<Pane defaultSize={0}>
			<div class="w-full h-full overflow-auto bg-app-lightest dark:bg-app-darkest">
				<BottomNav />
			</div>
		</Pane>
	</PaneGroup>
</div>

