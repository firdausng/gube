<script>
	import {GetActiveWorkspace} from "$lib/wailsjs/go/backend/App"
	import { goto } from '$app/navigation';

	let workspaceListPromise = getWorkspaceList()

	async function getWorkspaceList(){
		const list = await GetActiveWorkspace()
		if(list.Data){
			await goto(`/workspaces/${list.Data.ID}`)
			return list.Data;
		}
		throw new Error(list);
	}

	// let promise = getRandomNumber();
	//
	// async function getRandomNumber() {
	// 	const res = await fetch('http://localhost:3001')
	// 	const text = await res.json();
	//
	// 	if (res.ok) {
	// 		return text;
	// 	} else {
	// 		throw new Error(text);
	// 	}
	// }

</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>

<div>
	<a href="/workspaces">workspace</a>
</div>

<div>
	<a href="/workspaces/abc">workspace abc</a>
</div>

{#await workspaceListPromise}
	<p>...waiting</p>
{:then data}
	<p>The message is {JSON.stringify(data)}</p>
{:catch error}
	<p style="color: red">{error.message}</p>
{/await}