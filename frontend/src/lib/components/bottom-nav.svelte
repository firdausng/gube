<script lang="ts">
	import CodemirrorEditor from "$lib/components/editors/codemirror-editor.svelte";
	import XtermTerminal from "$lib/components/terminal/xterm-terminal.svelte";
	import type { SvelteComponent } from 'svelte';

	type TabItem = {
		/* eslint-disable  @typescript-eslint/no-explicit-any */
		component: typeof SvelteComponent<any>,
		name: string
	}
	type Tab = {
		active: TabItem,
		tabs: TabItem[]
	}

	let tabData: Tab = {
		active: {name: "logs", component: CodemirrorEditor},
		tabs: [
			{name: "logs", component: CodemirrorEditor},
			{name: "terminal", component: XtermTerminal},
		]
	}

	function changeTab(tab: TabItem) {
		tabData.active = tab
		tabData = {...tabData};
	}
</script>

<footer class="w-full">
	<div class="bg-app-light dark:bg-app-dark border-t-2 border-app-darkest dark:border-app-lightest">
		<div class="border-b border-app-darkest dark:border-app-lightest flex justify-items-start items-stretch w-full">
			<ul class="-mb-px flex gap-4 text-sm font-medium bg-app-light dark:bg-app-dark">
				{#each tabData.tabs as tab (tab.name)}
					<li class="flex-1">
						<button
							on:click={()=> changeTab(tab)}
							class="relative flex items-center justify-center gap-2 px-1 py-1 hover:text-app-dark dark:hover:text-app-light text-xs
						{tabData.active.name === tab.name ?
						'text-app-darkest dark:text-app-lightest after:absolute after:left-0 after:bottom-0 after:h-0.5 after:w-full after:bg-app-light dark:after:bg-app-dark':
						'text-app-darkest/50 dark:text-app-lightest/50'
						}"
						>
							{tab.name}</button
						>
					</li>
				{/each}
			</ul>
		</div>
		<div>
			<div
				class=""
			>
				<svelte:component this={tabData.active.component} />
			</div>
		</div>
	</div>
</footer>

