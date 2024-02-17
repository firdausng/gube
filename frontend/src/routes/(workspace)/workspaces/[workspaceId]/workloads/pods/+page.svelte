<script lang="ts">
	import {GetPodList} from "$lib/wailsjs/go/backend/App"
	import {type Readable, writable} from 'svelte/store'
	import {
		createSvelteTable,
		flexRender,
		getCoreRowModel, type Row, type Table,
	} from '@tanstack/svelte-table'
	import type { ColumnDef, TableOptions } from '@tanstack/svelte-table'
	import {onMount} from "svelte";
	import {type AppData, appDataStore} from '$lib/store/app-data-store';

	type Pod = {
		name: string
		namespace: string
		phase: string
	}

	let appData: AppData;
	appDataStore.subscribe(data =>{
		appData = data;
	})

	const defaultColumns: ColumnDef<Pod>[] = [
		{
			accessorKey: 'name',
			header: () => 'Name',
			cell: info => info.getValue(),
			footer: info => info.column.id,
		},
		{
			accessorKey: 'namespace',
			header: () => 'Namespace',
			cell: info => info.getValue(),
			footer: info => info.column.id,
		},
		{
			accessorKey: 'phase',
			header: () => 'Phase',
			cell: info => info.getValue(),
			footer: info => info.column.id,
		},
	]

	let table:Readable<Table<Pod>>

	let podListPromise = getPodList()
	async function getPodList(){
		if(appData){
			const list = await GetPodList(appData.activeWorkspace.activeContext.name,"")
			console.log(list)
			let podList = list.data.map((l:any) => {
				const pod: Pod = {
					name: l.metadata.name,
					namespace: l.metadata.namespace,
					phase: l.status.phase,
				}
				return pod
			})
			console.log(podList)

			const options = writable<TableOptions<Pod>>({
				data: podList,
				columns: defaultColumns,
				getCoreRowModel: getCoreRowModel(),
			})

			table = createSvelteTable(options)
		}

	}

	function handlePodClicked(row: Row<Pod>) {
		console.log(row.original)
	}



</script>

{#await podListPromise}
	<p>...waiting</p>
{:then data}
	<div class="flex flex-col w-full">
		<div class="flex flex-col">
			<div class="">
				<div class="inline-block min-w-full">
					<div class="inline-block min-w-full">
						<table class="min-w-full text-left text-sm">
							<thead class="border-b font-medium">
							{#each $table.getHeaderGroups() as headerGroup}
								<tr>
									{#each headerGroup.headers as header}
										<th class="px-2 py-2">
											{#if !header.isPlaceholder}
												<svelte:component
														this={flexRender(
															header.column.columnDef.header,
															header.getContext()
                  										)}
												/>
											{/if}
										</th>
									{/each}
								</tr>
							{/each}
							</thead>
							<tbody>
							{#each $table.getRowModel().rows as row}
								<tr on:click={()=> handlePodClicked(row)} class="text-app-dark/80 dark:text-app-light/80
								even:bg-app-light odd:bg-app-lightest dark:even:bg-app-dark dark:odd:bg-app-darkest
								hover:text-app-darkest dark:hover:text-app-lightest
								cursor-pointer">
									{#each row.getVisibleCells() as cell}
										<td class="whitespace-nowrap px-2 py-2 font-sm">
											<svelte:component
													this={flexRender(cell.column.columnDef.cell, cell.getContext())}
											/>
										</td>
									{/each}
								</tr>
							{/each}
							</tbody>
						</table>
					</div>
				</div>
			</div>
		</div>

		<div class="h-4" />
	</div>
{:catch error}
	<p style="color: red">{error.message}</p>
{/await}

