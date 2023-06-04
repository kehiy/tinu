<script>
    import Card from "./Card.svelte";
    import Modal from "./Modal.svelte";
    import { Modals, closeModal, openModal  } from "svelte-modals";

    export let tinu
    let showCard = true

    async function update(data){
        const json = {
            id:data.id,
            url:data.url
        }
        console.log(json);

        await fetch("http://localhost:3000/tinu",{
            method:"PATCH",
            headers:{"Content-Type":"application/json"},
            body: JSON.stringify(json)
        }).then(res=>{
            closeModal();
        })
    }

    function handleOpen(tinu){
        openModal(Modal,{
            title:"Update Tinu Link",
            send: update,
            url:tinu.url,
            id:tinu.id
        })
    }
</script>

{#if showCard}
<Card>
    <p>Tinu: http://localhost:3000/{tinu.id}</p>
    <p>url: {tinu.url}</p>
    <p>Clicked: {tinu.clicked}</p>
    <button class="update" on:click={handleOpen(tinu)}>Update</button>
</Card>
<Modals>
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <div on:click={closeModal} slot="backdrop" class="backdrop" />
</Modals>
{/if}
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