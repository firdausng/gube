<script>
	import ContextSidenav from "$lib/components/context-sidenav.svelte";
	import BottomNav from "$lib/components/bottom-nav.svelte";
	import { appSettingStore } from '$lib/store/app-setting-store';
	import { Pane, Splitpanes } from 'svelte-splitpanes';

</script>

<div class="flex gap-2 h-full">
	{#if !$appSettingStore.leftSidebarHidden}
		<ContextSidenav />
	{/if}

	<Splitpanes class="" horizontal={true}>
		<Pane class="bg-app-lightest dark:bg-app-darkest">
			<slot />
		</Pane>
		<Pane minSize={1} size={1}>
			<div class="w-full h-full overflow-auto">
				<BottomNav />
			</div>
		</Pane>
	</Splitpanes>
</div>




