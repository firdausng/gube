<script lang="ts">
    export let data:any;

    function onClick(){
        alert(data)
    }
</script>

<div>
    <iconify-icon on:click={onClick} icon="mingcute:menu-line" class="px-1 text-xl cursor-pointer text-app-dark dark:text-app-lightest text-center w-full"></iconify-icon>
</div>
