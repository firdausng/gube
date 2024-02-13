<script>
	import 'iconify-icon'
	import Theme from './theme.svelte';
</script>

<div class="flex flex-col gap-1 shadow-sm shadow-app-dark/50 dark:shadow-app-light/50 pt-1">
	<section class="border-r-2 border-app-darkest dark:border-app-lightest p-1">
		<Theme>
			<iconify-icon icon="fluent:dark-theme-20-filled" class="px-1 text-xl cursor-pointer text-app-dark dark:text-app-lightest text-center w-full"></iconify-icon>
		</Theme>
	</section>

	<section class="border-r-2 border-app-darkest dark:border-app-lightest p-1">
		<div class="px-1 border-2 border-app-dark dark:border-app-light">
			<p class="text-sm">SMD</p>
		</div>
	</section>

</div>