<script lang="ts">
	import { writable } from 'svelte/store';
	import {createTable, Subscribe, Render, DataBodyRow} from 'svelte-headless-table';
	import { addTableFilter } from 'svelte-headless-table/plugins';
	import {GetContextList} from "$lib/wailsjs/go/services/ContextService";
	import {onMount} from "svelte";
	import { goto } from '$app/navigation';
	import {type AppData, appDataStore, type WorkspaceContext} from '$lib/store/app-data-store';

	type K8sContext = {
		cluster: string
		user: string
	}

	let appData: AppData;
	appDataStore.subscribe(p =>{
		appData = p;
	})

	const data = writable<K8sContext[]>([]);

	const table = createTable(data, {
		tableFilter: addTableFilter(),
	});

	const columns = table.createColumns([
		table.column({
			header: 'Cluster',
			accessor: 'cluster',
		}),
		table.column({
			header: 'User',
			accessor: 'user',
		}),
	]);

	const { visibleColumns, headerRows, rows, tableAttrs, tableBodyAttrs, pluginStates } = table.createViewModel(columns);
	const { filterValue } = pluginStates.tableFilter;

	onMount(async ()=>{
		const list = await GetContextList()
		console.log(list)
		data.set(list.data)
	})

	// let contextListPromise = getContextList()
	// async function getContextList(){
	// 	const list = await GetContextList()
	//
	// 	// let options = writable<TableOptions<K8sContext>>({
	// 	// 	data: list.data,
	// 	// 	columns: defaultColumns,
	// 	// 	getCoreRowModel: getCoreRowModel(),
	// 	// })
	// 	// table = createSvelteTable(options)
	// 	// if(list.data){
	// 	// 	return list.data;
	// 	// }
	// }

	async function handleContextClicked(row: DataBodyRow<K8sContext>) {
		/*
                * 1. check if context already exist in active workspace context list
                * 2. if not exist,
                * 		- create the context
                * 		- add the context to the list
                * 3. update active context to the selected context
                * 4. set the resources, navigate to the resource
                * */

		console.log(row)
		let workspaceId = appData.activeWorkspace.id;
		let resourceName= appData.activeWorkspace.activeContext?.activeResource.name ?
				appData.activeWorkspace.activeContext?.activeResource.name:
				"pods";

		appDataStore.update(d =>{
			const activeContext: WorkspaceContext = {
				id: row.original.cluster,
				name: row.original.cluster,
				active: true,
				activeResource: {
					name: resourceName
				},
				tabData: {
					activeTab: null,
					tabs: []
				}
			}

			let foundItem = d.activeWorkspace.contexts.find(item => item.id === activeContext.id);

			if(foundItem) {
				foundItem = {...foundItem, ...activeContext};
			} else {
				d.activeWorkspace.contexts = [...d.activeWorkspace.contexts, activeContext];
			}

			d.activeWorkspace = {
				...d.activeWorkspace,
				activeContext
			}
			return d;
		})

		console.log("appData", appData)
		await goto(`/workspaces/${workspaceId}/workloads/${resourceName}?user=${row.original.user}&cluster=${row.original.cluster}`)
	}
</script>

{#if $data.length > 0}
	<div class="flex flex-col w-full">
		<div class="flex flex-col">
			<div class="">
				<div class="inline-block min-w-full">
					<div class="inline-block min-w-full">
						<table class="min-w-full text-left text-sm" {...$tableAttrs}>
							<thead class="border-b font-medium">
							{#each $headerRows as headerRow (headerRow.id)}
								<Subscribe rowAttrs={headerRow.attrs()} let:rowAttrs>
									<tr {...rowAttrs}>
										{#each headerRow.cells as cell (cell.id)}
											<Subscribe attrs={cell.attrs()} let:attrs props={cell.props()} let:props>
												<th {...attrs}>
													<Render of={cell.render()} />
												</th>
											</Subscribe>
										{/each}
									</tr>
								</Subscribe>
							{/each}
							</thead>
							<tbody {...$tableBodyAttrs}>
							{#each $rows as row (row.id)}
								<Subscribe rowAttrs={row.attrs()} let:rowAttrs>
									<tr {...rowAttrs}
										on:click={()=> handleContextClicked(row)}
										class="text-app-dark/80 dark:text-app-light/80
											even:bg-app-light odd:bg-app-lightest dark:even:bg-app-dark dark:odd:bg-app-darkest
											hover:text-app-darkest dark:hover:text-app-lightest
											cursor-pointer">
										{#each row.cells as cell (cell.id)}
											<Subscribe attrs={cell.attrs()} let:attrs props={cell.props()} let:props>
												<td {...attrs} class:matches={props.tableFilter.matches} class="whitespace-nowrap px-2 py-2 font-sm">
													<Render of={cell.render()} />
												</td>
											</Subscribe>
										{/each}
									</tr>
								</Subscribe>
							{/each}
							</tbody>
						</table>
					</div>
				</div>
			</div>
		</div>

		<div class="h-4" />
	</div>
{/if}
