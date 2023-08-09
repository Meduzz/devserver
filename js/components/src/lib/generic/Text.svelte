<script>
    export let text
    export let html = false
</script>

<div>
    {#if html}
        {@html text}
    {:else}
        {text}
    {/if}
</div>