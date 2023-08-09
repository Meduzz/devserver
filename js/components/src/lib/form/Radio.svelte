<script>
import { RadioGroup } from '@chimps/genericui'
import { getContext } from 'svelte'

export let label
export let items
export let name

let form = getContext('store')
</script>

<div>
    <span class="block">{label}:</span>
    <RadioGroup {items} {name} bind:value={$form[name]} />
</div>