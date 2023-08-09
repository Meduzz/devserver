<script>
    import {Panel, ListY} from '@chimps/genericui'
    
    export let items = []
</script>

<Panel class="p-0">
    <ListY {items} let:item class="divide-y">
        {#if item.href}
            <a href="{item.href}" class="hover:bg-slate-200 p-2 text-primary underline">
                {item.label}
            </a>
        {:else}
            <span class="p-2 hover:bg-slate-200">{item.label}</span>
        {/if}
    </ListY>
</Panel>