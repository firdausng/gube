<script>
	import { onMount } from 'svelte';
	import { EventsOn } from '$lib/wailsjs/runtime';
	import { appSettingStore } from '$lib/store/app-setting-store';

	let theme = ''

	appSettingStore.subscribe(val => {
		theme = val.theme;
	})
	onMount(()=>{
		// EventsOn("change-theme", data => {
		// 	setTheme()
		// })

		if (theme === "dark") {
			document.documentElement.classList.add('dark')
		}else{
			document.documentElement.classList.remove('dark')
		}
	})

	function setTheme() {
		console.log(theme)
		if (theme === "dark") {
			appSettingStore.update((val)=>{
				val.theme = 'light'
				return val;
			})
			document.documentElement.classList.remove('dark')
		}else{
			appSettingStore.update((val)=>{
				val.theme = 'dark'
				return val;
			})
			document.documentElement.classList.add('dark')
		}
	}
</script>

<div on:click={setTheme} class="flex items-center">
	<slot />
</div>
