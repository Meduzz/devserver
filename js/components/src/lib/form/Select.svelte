<script>
import { Select } from '@chimps/genericui'
import Spacer from './Spacer.svelte'
import { getContext } from 'svelte'

export let label
export let items
export let name
/*
    The plan was to have MultiSelect in here as well, but it did not play
    well with other components, which there inheritly are in a form.
*/
let form = getContext('store')
</script>

<Spacer {label}>
    <Select {name} {items} bind:value={$form[name]} class="py-1 border-b border-slate-300 text-primary w-full block caret-primary focus:outline-none focus:border-primary" {...$$restProps} let:item>
        <option value="{item.value}">{item.label}</option>
    </Select>
</Spacer>