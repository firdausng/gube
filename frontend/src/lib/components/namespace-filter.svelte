<script lang="ts">
	// import { createSelect, melt } from '@melt-ui/svelte';
	// import { Check, ChevronDown } from '$icons/index.js';

	const options = {
		sweet: [
			'Caramel',
			'Chocolate',
			'Strawberry',
			'Cookies & Cream',
			'Cookies & 1',
			'Cookies & 2',
			'Cookies & 3',
		],
	};

	// const {
	// 	elements: { trigger, menu, option, group, groupLabel, label },
	// 	states: { selectedLabel, open },
	// 	helpers: { isSelected },
	// } = createSelect({
	// 	forceVisible: true,
	// 	positioning: {
	// 		placement: 'bottom',
	// 		fitViewport: true,
	// 		sameWidth: true,
	// 	},
	// 	multiple: true,
	// });
</script>

<div class="flex gap-1 p-2">
<!--	<button-->
<!--		class="flex min-w-[220px] items-center justify-between rounded-lg bg-app-lightest dark:bg-app-darkest px-3 py-1 shadow transition-opacity hover:opacity-90 max-w-64"-->
<!--		{...$trigger} use:trigger-->
<!--		aria-label="Food"-->
<!--	>-->
<!--		{$selectedLabel || 'Namespaces'}-->
<!--&lt;!&ndash;		<ChevronDown class="size-5" />&ndash;&gt;-->
<!--		<iconify-icon icon="heroicons:arrow-down" class="text-md cursor-pointer text-app-dark dark:text-app-lightest"></iconify-icon>-->
<!--	</button>-->
<!--	{#if $open}-->
<!--		<div-->
<!--			class=" z-10 flex max-h-[200px] flex-col-->
<!--    overflow-y-auto rounded-lg bg-app-lightest dark:bg-app-darkest p-1 shadow-sm shadow-app-dark dark:shadow-app-light-->
<!--     focus:!ring-0"-->
<!--			{...$menu} use:menu-->
<!--		>-->
<!--			{#each Object.entries(options) as [key, arr]}-->
<!--				<div {...$group(key)} use:group>-->
<!--					<div-->
<!--						class="py-1 pl-4 pr-4 font-semibold capitalize"-->
<!--						{...$groupLabel(key)} use:groupLabel-->
<!--					>-->
<!--						{key}-->
<!--					</div>-->
<!--					{#each arr as item}-->
<!--						<div-->
<!--							class="relative cursor-pointer rounded-lg py-1 pl-8 pr-4-->
<!--              focus:z-10 focus:text-magnum-700-->
<!--            data-[highlighted]:bg-magnum-50 data-[selected]:bg-magnum-100-->
<!--            data-[highlighted]:text-magnum-900 data-[selected]:text-magnum-900"-->
<!--							{...$option({ value: item, label: item })} use:option-->
<!--						>-->
<!--							<div class="flex gap-2">-->
<!--								<p>{item}</p>-->
<!--								<div class="check {$isSelected(item) ? 'block' : 'hidden'}">-->
<!--									&lt;!&ndash;								<Check class="size-4" />&ndash;&gt;-->
<!--									<iconify-icon icon="mingcute:check-fill" class="text-md cursor-pointer text-app-dark dark:text-app-lightest"></iconify-icon>-->

<!--								</div>-->
<!--							</div>-->

<!--						</div>-->
<!--					{/each}-->
<!--				</div>-->
<!--			{/each}-->
<!--		</div>-->
<!--	{/if}-->
</div>

<!--<style lang="postcss">-->
<!--    .check {-->
<!--        position: absolute;-->
<!--        left: theme(spacing.2);-->
<!--        top: 50%;-->
<!--        z-index: theme(zIndex.20);-->
<!--        translate: 0 calc(-50% + 1px);-->
<!--        color: theme(colors.magnum.500);-->
<!--    }-->
<!--</style>-->