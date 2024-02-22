<script lang="ts">
    export let data:any;
</script>

{#if data}
<div class="flex {data.length > 3 ? 'flex-col': 'flex-row'} w-[300px] gap-1">
        {#each data as d (d.name)}
            <div class="bg-green-200 dark:bg-green-800 truncate rounded border px-1">
                <p>{d.name}</p>
                <p>restart: {d.restartCount}</p>
            </div>
        {/each}
</div>
{/if}

