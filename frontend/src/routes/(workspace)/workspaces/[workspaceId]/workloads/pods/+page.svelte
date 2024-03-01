<script lang="ts">
	import {GetPodList,StreamPods} from "$lib/wailsjs/go/services/PodService"
	import {writable} from 'svelte/store'
	import {type AppData, appDataStore} from '$lib/store/app-data-store';
	import ContainerRow from './ContainerRow.svelte'
	import OwnerReference from './OwnerReference.svelte'
	import PodMenu from './PodMenu.svelte'
	import {onDestroy, onMount} from "svelte";
	import {EventsOn, EventsOff} from "$lib/wailsjs/runtime";
	import {createRender, createTable, DataBodyRow, Render, Subscribe} from "svelte-headless-table";
	import {addTableFilter} from "svelte-headless-table/plugins";
	import type {Pod} from "$lib/types";

	let appData: AppData;
	appDataStore.subscribe(data =>{
		appData = data;
	})

	let eventName = '';
	let namespaceFilter = ''; // let it blank so that we just get all
	let podList: Pod[] = [];

	const data = writable<Pod[]>([]);
	const table = createTable(data, {
		tableFilter: addTableFilter(),
	});

	const columns = table.createColumns([
		table.column({
			header: 'Name',
			accessor: 'name',
		}),
		table.column({
			header: 'Namespace',
			accessor: 'namespace',
		}),
		table.column({
			header: 'Phase',
			accessor: 'phase',
		}),
		table.column({
			header: 'Containers',
			accessor: 'containers',
			cell: (data, { pluginStates }) =>{
				// const { isSelected, isSomeSubRowsSelected } = pluginStates.select.getRowState(row);
				return createRender(ContainerRow, {
					data
				});
			},
		}),
		table.column({
			header: 'Owner',
			accessor: 'owner',
			cell: (data, { pluginStates }) =>{
				// const { isSelected, isSomeSubRowsSelected } = pluginStates.select.getRowState(row);
				return createRender(OwnerReference, {
					data
				});
			},
		}),
		table.column({
			header: 'Menu',
			accessor: 'menu',
			cell: (data, { pluginStates }) =>{
				// console.log('menu', data.value)
				return createRender(PodMenu, {
					podName: data.value.podName,
					namespace: data.value.namespace
				});
			},
		}),
	]);

	const { visibleColumns, headerRows, rows, tableAttrs, tableBodyAttrs, pluginStates } = table.createViewModel(columns);
	const { filterValue } = pluginStates.tableFilter;

	onMount(async ()=>{
		await getPodList();
		if(appData.activeWorkspace.activeContext){
			// console.log('EventsOn', eventName)
			EventsOn(eventName, (emitPodEvent)=>{
				// console.log('emit log', emitPodEvent)
				const emitPodData = emitPodEvent.Object
				// console.log('emit log', emitPodData)
				let pod: Pod= setupPod(emitPodData)
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
							podList =[pod, ...podList];
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

				data.update(opt =>{
					opt = [...podList]
					return opt;
				})
			})
		}
	})

	async function getPodList(){
		if(appData.activeWorkspace.activeContext){
			const response = await GetPodList(appData.activeWorkspace.name, appData.activeWorkspace.activeContext.name,namespaceFilter)
			// console.log(response.data)
			podList = response.data.map((l:any) => setupPod(l))
			// console.log(podList)
			eventName = `EmitPodList-${appData.activeWorkspace.id}-${appData.activeWorkspace.activeContext.name}-${namespaceFilter}`
			StreamPods(`${appData.activeWorkspace.id}`, appData.activeWorkspace.activeContext.name,namespaceFilter)

			data.update(opt =>{
				opt = podList
				return opt;
			})
		}
	}

	function handlePodClicked(row: DataBodyRow<Pod>) {
		// console.log(row)
	}

	function setupPod(data:any){
		let pod: Pod = {
			name: data.metadata.name,
			namespace: data.metadata.namespace,
			phase: data.status.phase,
			containers: data.status.containerStatuses,
			owner: data.metadata.ownerReferences,
			menu: {
				podName: data.metadata.name,
				namespace: data.metadata.namespace
			}
		}
		return pod
	}

	onDestroy(()=>{
		EventsOff(eventName)
	})
</script>

{#if $data.length > 0}
	<div class="flex flex-col w-full h-full overflow-auto">
		<div class="flex flex-col">
			<div class="">
				<div class="inline-block min-w-full">
					<div class="inline-block min-w-full">
						<table class=" min-w-full text-left text-sm" {...$tableAttrs}>
							<thead class="sticky top-0  border-b font-medium w-full bg-app-lightest dark:bg-app-darkest shadow-md shadow-app-light dark:shadow-app-dark">
							{#each $headerRows as headerRow (headerRow.id)}
								<Subscribe rowAttrs={headerRow.attrs()} let:rowAttrs>
									<tr {...rowAttrs}>
										{#each headerRow.cells as cell (cell.id)}
											<Subscribe attrs={cell.attrs()} let:attrs props={cell.props()} let:props>
												<th class="px-2 py-2" {...attrs}>
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
									<tr {...rowAttrs} on:click={()=> handlePodClicked(row)}
										class="text-app-dark/80 dark:text-app-light/80
										even:bg-app-light odd:bg-app-lightest dark:even:bg-app-dark dark:odd:bg-app-darkest
										hover:text-app-darkest dark:hover:text-app-lightest
										cursor-pointer">
										{#each row.cells as cell (cell.id)}
											<Subscribe attrs={cell.attrs()} let:attrs props={cell.props()} let:props>
												<td {...attrs} class="whitespace-nowrap px-2 py-1 font-sm">
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
