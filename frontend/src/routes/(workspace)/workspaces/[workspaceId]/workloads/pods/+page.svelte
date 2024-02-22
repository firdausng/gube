<script lang="ts">
	import {GetPodList,
		StreamPods
	} from "$lib/wailsjs/go/services/PodService"
	import {type Readable, writable} from 'svelte/store'
	import {
		createSvelteTable,
		flexRender,
		getCoreRowModel, renderComponent, type Row, type Table,
	} from '@tanstack/svelte-table'
	import type { ColumnDef, TableOptions } from '@tanstack/svelte-table'
	import {type AppData, appDataStore} from '$lib/store/app-data-store';
	import ContainerRow from './ContainerRow.svelte'
	import OwnerReference from './OwnerReference.svelte'
	import PodMenu from './PodMenu.svelte'
	import {onMount} from "svelte";
	import {EventsOff, EventsOn} from "$lib/wailsjs/runtime";

	type Pod = {
		name: string
		namespace: string
		phase: string
		containers: any
		owner: any
		menu: any
	}

	let appData: AppData;
	appDataStore.subscribe(data =>{
		appData = data;
	})

	let eventName = '';
	let namespaceFilter = ''; // let it blank so that we just get all

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
		{
			accessorKey: 'containers',
			header: () => 'Containers',
			cell: info => renderComponent(ContainerRow, {data: info.getValue() as any}),
			footer: info => info.column.id,
		},
		{
			accessorKey: 'owner',
			header: () => 'Owner',
			cell: info => renderComponent(OwnerReference, {data: info.getValue() as any}),
			footer: info => info.column.id,
		},
		{
			accessorKey: 'menu',
			header: () => '',
			cell: info => renderComponent(PodMenu, {data: info.getValue() as any}),
			footer: info => info.column.id,
		},
	]

	let podList: Pod[] = [];
	let table:Readable<Table<Pod>>
	let options = writable<TableOptions<Pod>>({
		data: podList,
		columns: defaultColumns,
		getCoreRowModel: getCoreRowModel(),
	})
	table = createSvelteTable(options)

	async function getPodList(){
		if(appData.activeWorkspace.activeContext){
			const response = await GetPodList(appData.activeWorkspace.name, appData.activeWorkspace.activeContext.name,namespaceFilter)
			console.log(response.data)
			podList = response.data.map((l:any) => {
				const pod: Pod = {
					name: l.metadata.name,
					namespace: l.metadata.namespace,
					phase: l.status.phase,
					containers: l.status.containerStatuses,
					owner: l.metadata.ownerReferences,
					menu: l.metadata.name
				}
				return pod
			})
			console.log(podList)
			eventName = `EmitPodList-${appData.activeWorkspace.id}-${appData.activeWorkspace.activeContext.name}-${namespaceFilter}`
			StreamPods(`${appData.activeWorkspace.id}`, appData.activeWorkspace.activeContext.name,namespaceFilter)

			options.update(opt =>{
				opt.data = podList
				return opt;
			})
		}
	}
	let podListPromise = getPodList()

	$: if(appData.activeWorkspace.activeContext){

		console.log('EventsOn', eventName)
		EventsOn(eventName, (emitPodEvent)=>{
			console.log('emit log', emitPodEvent)
			const emitPodData = emitPodEvent.Object
			// console.log('emit log', emitPodData)
			let pod: Pod = {
				name: emitPodData.metadata.name,
				namespace: emitPodData.metadata.namespace,
				phase: emitPodData.status.phase,
				containers: emitPodData.status.containerStatuses,
				owner: emitPodData.metadata.ownerReferences,
				menu: emitPodData.metadata.name
			}
			// console.log('emit log', pod)


			switch (emitPodEvent.Type){
				case 'MODIFIED': {
					console.log('MODIFIED', pod.phase)
					podList = podList.map(p => {
						if(p.name === pod.name) {
							return pod
						}
						return p
					})
					break;
				}
				case 'ADDED': {
					console.log('ADDED', pod.phase)
					if(!podList.some(obj => obj.name === pod.name)) {
						podList.push(pod);
					}
					break;
				}
				case 'DELETED': {
					console.log('DELETED', pod.phase)
					podList = podList.filter(obj => obj.name !== pod.name)
					break;
				}
				default: {

				}
			}
			podList = [...podList];

			options.update(opt =>{
				// console.log('update log', podList)
				opt.data = [...podList]
				return opt;
			})
		})


	}

	function handlePodClicked(row: Row<Pod>) {
		console.log(row.renderValue('name'))
	}
</script>

{#await podListPromise}
	<p>...waiting</p>
{:then _}
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
										<td class="whitespace-nowrap px-2 py-1 font-sm">
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

