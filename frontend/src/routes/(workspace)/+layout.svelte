<script>
	import TopNav from "$lib/components/top-nav.svelte";
</script>

<div class="relative flex flex-col gap-1 h-screen">
	<div class="fixed z-10 w-full bg-app-lightest dark:bg-app-darkest shadow shadow-app-light dark:shadow-app-dark">
		<TopNav />
	</div>

	<div class="bg-app-lightest dark:bg-app-darkest w-full h-full p-1 pt-8">
		<slot />
	</div>
</div>
