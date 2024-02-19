<script lang="ts">
    export let data:any;
</script>

<div class="flex flex-col w-[200px] gap-1">
    {#each data as d (d.uid)}
        <div class="">
            <p>{d.kind}</p>
        </div>
    {/each}
</div>
