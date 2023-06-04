<script>
    import Item from "./Item.svelte"
    import { onMount } from "svelte"

    let tinus = []

    onMount(async () => {
        const res = await fetch("http://localhost:3000/tinu");
        tinus = await res.json();
    })
</script>

{#each tinus as tinu}
    <Item />
{/each}