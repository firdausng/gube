<script lang="ts">
	import 'iconify-icon'
	import Theme from '$lib/components/theme.svelte';
	import NamespaceFilter from '$lib/components/namespace-filter.svelte';
	import { appSettingStore } from '$lib/store/app-setting-store';
	import { appDataStore, type Workspace } from '$lib/store/app-data-store';


	let workspace: Workspace;
	appDataStore.subscribe(data => {
		if(data.activeWorkspace){
			workspace = data.activeWorkspace
		}
	})

	function setActiveContext(id:string) {
		workspace.contexts = workspace.contexts.map(context => ({...context, active: context.id === id}));
		workspace = workspace;
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

		{#if workspace}
			{#each workspace.contexts as context (context.id)}
				<section class="flex items-center border-b-2 {context.active ? ' border-app-darkest dark:border-app-lightest': 'border-app-lightest/50 dark:border-app-darkest/50'} p-1 cursor-pointer ease-in duration-200">
					<button class="text-sm px-1" on:click={()=>setActiveContext(context.id)}>{context.name}</button>
				</section>
			{/each}

			<section class="dark:border-app-lightest p-1 flex items-center">
				<p class="capitalize">{workspace.activeContext?.activeResource ? workspace.activeContext.activeResource.name : ''}</p>
			</section>

		{/if}


	</div>

	<div class="shrink">
		<NamespaceFilter />
	</div>
</div>