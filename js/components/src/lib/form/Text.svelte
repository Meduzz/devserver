<script>
import { Input, Textarea } from '@chimps/genericui'
import Spacer from './Spacer.svelte'
import { getContext } from 'svelte'

export let kind = 'text'
export let label
export let name

let form = getContext('store')
</script>

<Spacer {label}>
    {#if kind == 'textarea'}
    <Textarea {name} bind:value={$form[name]} class="py-1 border-b border-slate-300 text-primary w-full caret-primary focus:outline-none focus:border-primary" {...$$restProps} />
    {:else}
    <Input type="{kind}" {name} bind:value={$form[name]} class="py-1 border-b border-slate-300 text-primary w-full block caret-primary focus:outline-none focus:border-primary" {...$$restProps} />
    {/if}
</Spacer>