<script>
    import Card from "./Card.svelte";
    export let tinu
</script>

<Card>
    <p>Tinu: https://localhost:3000/{tinu.id}</p>
    <p>url: {tinu.url}</p>
    <p>Clicked: {tinu.clicked}</p>
    <button class="update">Update</button>
    <button class="delete">Delete</button>
</Card>

<style>
    button {
        color: white;
        font-weight: bolder;
        border: none;
        padding: .75rem;
        border-radius: 4px;
    }
    .update{
        background-color: yellowgreen;
    }
    .delete{
        background-color: red;
    }
</style>