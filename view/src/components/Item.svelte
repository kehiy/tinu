<script>
    import Card from "./Card.svelte";
</script>

<Card>
    <p>Tinu: https://localhost:3000/{"40c24a8c6d354d4f"}</p>
    <p>url: {"url"}</p>
    <p>Clicked: {"clicked"}</p>
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