<script>
import { CheckboxGroup } from '@chimps/genericui'
import Spacer from './Spacer.svelte'
import { getContext } from 'svelte'

export let items
export let label
export let name

let form = getContext('store')
</script>

<Spacer {label}>
    <CheckboxGroup bind:values={$form[name]} {items} {name} />
</Spacer>