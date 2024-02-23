<script lang="ts">
    export let data:any;
</script>

<div class="flex w-[200px] gap-1">
    {#if data && data.value}
        {#each data.value as d (d.uid)}
            <div class="px-1.5 border border-app-darkest dark:border-app-lightest">
                <p>{d.kind}</p>
            </div>
        {/each}
    {/if}

</div>
