<script lang="ts">
    export let data:any;
</script>

<div class="flex  w-[200px] gap-1">
    {#each data as d (d.name)}
        <span class="bg-green-200 dark:bg-green-800 truncate rounded border px-1">{d.name}</span>
    {/each}
</div>

