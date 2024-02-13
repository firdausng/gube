<script lang="ts">
	import type {NavigationItem} from "$lib/types";
	import 'iconify-icon'
	let navigationList: NavigationItem[] = [
		{
			folder: 'Workloads',
			name: 'Pods',
			path: 'workloads/pods'
		},
		{
			folder: 'Workloads',
			name: 'Deployments',
			path: 'workloads/deployments'
		},
		{
			folder: 'Config',
			name: 'ConfigMaps',
			path: 'config/config-maps'
		},
		{
			folder: 'Config',
			name: 'Secrets',
			path: 'config/secrets'
		},
		{
			folder: 'Network',
			name: 'Services',
			path: 'network/services'
		},
		{
			folder: 'Network',
			name: 'Endpoints',
			path: 'network/endpoints'
		},
		{
			folder: 'Network',
			name: 'Ingresses',
			path: 'network/ingresses'
		}
	]

	let folderSettings = [
		{name:'Workloads', icon: "clarity:rack-server-line", open: true},
		{name:'Config', icon: "system-uicons:document-list", open: true},
		{name:'Network', icon: "bi:hdd-network", open: true},
	]

	const groupedByFolder: Record<string, NavigationItem[]> = navigationList.reduce<Record<string, NavigationItem[]>>(
		(result, item) => {
			(result[item.folder] = result[item.folder] || []).push(item);
			return result;
		},
		{}
	);

	function toggleOpenFolder(name:string) {
		folderSettings = folderSettings.map(f => ({...f, open: (f.name === name) ? !f.open: f.open}));
		console.log(folderSettings);
	}
</script>

<div class="flex flex-col gap-1 h-full w-full max-w-48 overflow-y-auto text-sm bg-app-light dark:bg-app-dark pr-2 pt-2 shadow-sm shadow-app-dark dark:shadow-app-light">
	{#each folderSettings as folder (folder.name)}
		<div>
			<button class="flex gap-1" onclick={()=>toggleOpenFolder(folder.name)}>
				<iconify-icon icon={folder.icon} class="px-1 text-xl cursor-pointer text-app-dark dark:text-app-lightest text-center w-full"></iconify-icon>
				<span class="break-words">{folder.name}</span>
			</button>
			{#if folder.open}
				<div class="flex flex-col overflow-y-auto pl-8">
					{#each groupedByFolder[folder.name] as navigation (navigation)}
						<a class="py-1 break-words" href="{`/workspaces/bac/${navigation.path}`}">{navigation.name}</a>
					{/each}
				</div>
			{/if}
		</div>
	{/each}
</div>

