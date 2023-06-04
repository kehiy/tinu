<script>
    import { Modals, closeModal, openModal } from "svelte-modals"
    import Modal from "./Modal.svelte"

    async function updateTinu(data) {
        const json = {
         url:data.url
        }
        await fetch("http://localhost:3000/tinu", {
            method: "POST",
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify(json)
        }).then(response => {
            console.log(response)
            closeModal()
        })
    }

    function handleOpen() {
        console.log("opend");
        openModal(Modal, {
            title: "Create New Tinu Link",
            send: updateTinu,
        })
    }
</script> 


<button on:click={ handleOpen() }>New</button>


<style>
    button {
        background-color: green;
        color: white;
        font-weight: bold;
        border: none;
        padding: .75rem;
        border-radius: 4px;
    }
</style>