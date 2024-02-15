<script lang="ts">
	import {GetContextList} from "$lib/wailsjs/go/backend/App"
	import {type Readable, writable} from 'svelte/store'
	import {
		createSvelteTable,
		flexRender,
		getCoreRowModel, type Table, type Row
	} from '@tanstack/svelte-table'
	import type { ColumnDef, TableOptions } from '@tanstack/svelte-table'

	type K8sContext = {
		cluster: string
		user: string
	}

	const defaultColumns: ColumnDef<K8sContext>[] = [
		{
			accessorFn: row => row.cluster,
			id: 'cluster',
			cell: info => info.getValue(),
			header: () => 'Cluster',
			footer: info => info.column.id,
		},
		{
			accessorFn: row => row.user,
			id: 'user',
			cell: info => info.getValue(),
			header: () => 'User',
			footer: info => info.column.id,
		},
	]

	let table:Readable<Table<K8sContext>>

	let contextListPromise = getContextList()
	async function getContextList(){
		const list = await GetContextList()
		let options = writable<TableOptions<K8sContext>>({
			data: list.Data,
			columns: defaultColumns,
			getCoreRowModel: getCoreRowModel(),
		})
		table = createSvelteTable(options)
		if(list.Data){
			return list.Data;
		}
		throw new Error(list);
	}

	function handleContextClicked(row: Row<K8sContext>) {
		console.log(row)
	}
</script>

{#await contextListPromise}
	<p>...waiting</p>
{:then data}
	<div class="flex flex-col w-full">
		<h2>Workspaces</h2>

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
								<tr on:click={()=> handleContextClicked(row)} class="text-app-dark/80 dark:text-app-light/80
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