<script>
	import ContextSidenav from "$lib/components/context-sidenav.svelte";
	import BottomNav from "$lib/components/bottom-nav.svelte";
	// import { resizer } from '$lib/actions/resizer';
	// import { layoutResizer } from '$lib/actions/layout-resizer';
	import { appSettingStore } from '$lib/store/app-setting-store';
	import { Pane, Splitpanes } from 'svelte-splitpanes';

</script>

<div class="flex gap-2 h-full">
	{#if !$appSettingStore.leftSidebarHidden}
		<div class="shrink">
			<ContextSidenav />
		</div>
	{/if}

	<Splitpanes class="" horizontal={true}>
		<Pane class="bg-app-lightest dark:bg-app-darkest">
			<slot />
		</Pane>
		<Pane minSize={1}>
			<div class="w-full h-full overflow-auto">
				<BottomNav />
			</div>
		</Pane>
	</Splitpanes>
<!--	<div use:layoutResizer={{flexType:'col'}} class="flex flex-col">-->
<!--		<div class="grow overflow-auto">-->
<!--			<slot />-->
<!--		</div>-->
<!--		<div class="layout-divider"></div>-->
<!--		<div class="w-full">-->
<!--&lt;!&ndash;			<div use:resizer class="h-[3px] bg-primary-500 cursor-row-resize"></div>&ndash;&gt;-->
<!--			<div class="w-full h-full overflow-auto">-->
<!--				<BottomNav />-->
<!--			</div>-->
<!--		</div>-->
<!--	</div>-->
</div>




