<script lang="ts">
    export let data:any;
</script>

<div class="flex w-[100px]">
    {#each data as d (d.name)}
        <span class="text-green-700 dark:text-green-400 truncate">{d.name}</span>
    {/each}
</div>

