<script lang="ts">
    async function fetchMessage(): Promise<string> {
        const res = await fetch("/api");
        const json: { message: string } = await res.json();
        console.log(json)
        console.log(typeof json)
        console.log(json.message)
        console.log(typeof json.message)
        return json.message;
    }
</script>

<h1>Message:</h1>
{#await fetchMessage() then message}
    <h2>{ message }</h2>
{/await}
