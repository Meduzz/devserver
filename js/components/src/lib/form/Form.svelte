<script>
import { Panel, Button } from '@chimps/genericui'
import { setContext } from 'svelte'
import { writable } from 'svelte/store'

export let method
export let action
export let title
export let format
export let back
export let redirect
export let store = writable({})

function submit() {
    let req = {
        method,
        headers: {}
    }

    switch (format) {
        case 'form':
            let form = new FormData()
            Object.keys($store).forEach(it => {
                form.append(it, $store[it])
            })
            req.body = form
            break
        case 'json':
            req.body = JSON.stringify($store)
            req.headers['Content-Type'] = 'application/json'
            break
        default:
            throw `unknown format: ${format}`
    }

    fetch(action, req)
        .then(res => {
            if (!res.ok) {
                throw res.statusText
            }

            store.set({})
            location.assign(redirect)
        })
}

setContext('store', store)
</script>

<svelte:head>
    {#if title}
        <title>{title}</title>
    {/if}
</svelte:head>

<Panel class="text-slate-700 lg:w-2/3 xl:w-1/2 lg:mx-auto">
    <form method={method} action={action} on:submit|preventDefault={submit}>
        <div>
            <slot name="header">
                <h3 class="pb-2 underline text-primary text-xl font-bold">{title}</h3>
            </slot>
            <hr />
        </div>
        <div>
            <slot></slot>
        </div>
        <div class="mt-2">
            <hr />
            <div class="pt-2">
                <slot name="footer">
                    <Button on:click={submit} label="Submit" class="bg-primary px-3 py-1 text-white" />
                    {#if back}
                    <a href="{back}" class="text-primary underline ml-2">Back</a>
                    {/if}
                </slot>
            </div>
        </div>
    </form>
</Panel>
