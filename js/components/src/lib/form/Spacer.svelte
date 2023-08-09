<script>
    export let label
</script>

<div class="py-2">
    {#if label}
        <span class="block font-bold">{label}:</span>
    {/if}
    <slot></slot>
</div>