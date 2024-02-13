<script>
	import 'iconify-icon'
	import Theme from '$lib/components/theme.svelte';
	import NamespaceFilter from '$lib/components/namespace-filter.svelte';
	import { appSettingStore } from '$lib/store/app-setting-store';

	let workspaces = [
		{id: 1, name:"SMD", active:true},
		{id: 2, name:"TBD", active:false},
	]

	function setActiveWorkspace(id) {
		workspaces = workspaces.map(workspace => ({...workspace, active: workspace.id === id}));
	}
</script>

<div class="flex">
	<div class="grow flex gap-2">
		<section class="dark:border-app-lightest p-1 flex items-center">
			<button on:click={()=> $appSettingStore.leftSidebarHidden = !$appSettingStore.leftSidebarHidden} class="flex items-center">
				<iconify-icon icon="mingcute:menu-line" class="px-1 text-xl cursor-pointer text-app-dark dark:text-app-lightest text-center w-full"></iconify-icon>
			</button>

		</section>

		<section class="dark:border-app-lightest p-1 flex items-center">
			<Theme>
				<iconify-icon icon="fluent:dark-theme-20-filled" class="px-1 text-xl cursor-pointer text-app-dark dark:text-app-lightest text-center w-full"></iconify-icon>
			</Theme>
		</section>

		<section class="dark:border-app-lightest p-1 flex items-center">
			<a href="/workspaces" class="flex items-center">
				<iconify-icon icon="carbon:home" class="px-1 text-xl cursor-pointer text-app-dark dark:text-app-lightest text-center w-full"></iconify-icon>
			</a>

		</section>

		{#each workspaces as workspace (workspace.id)}
			<section class="flex items-center border-b-2 {workspace.active ? ' border-app-darkest dark:border-app-lightest': 'border-app-lightest/50 dark:border-app-darkest/50'} p-1 cursor-pointer ease-in duration-200">
				<button class="text-sm px-1" on:click={()=>setActiveWorkspace(workspace.id)}>{workspace.name}</button>
			</section>
		{/each}
	</div>

	<div class="shrink">
		<NamespaceFilter />
	</div>
</div>