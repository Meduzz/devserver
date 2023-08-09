<script>
    import {Form, Checkbox, Radio, Select, Text} from '../form'
    import {Page, Navbar, List, Columns, ColumnShrink, ColumnGrow, Routable, Column, Markdown} from '../generic'
    import HttpLoader from './HttpLoader.svelte'
    import NotFound from './NotFound.svelte'
    import formMeta from '../form/meta.json'
    import genericMeta from '../generic/meta.json'
    import loaderMeta from './meta.json'

    export let page

    const components = {
        "form": Form,
        "checkboxes": Checkbox,
        "radioboxes": Radio,
        "select": Select,
        "text": Text,
        "navbar": Navbar,
        "page": Page,
        "list": List,
        "columns": Columns,
        "columnShrink": ColumnShrink,
        "columnGrow": ColumnGrow,
        "httpLoader": HttpLoader,
        "routable": Routable,
        "column": Column,
        "markdown": Markdown,
    }

    let allMeta = formMeta.components.concat(genericMeta.components, loaderMeta.components)

    let currentMeta = allMeta.find(it => it.name == page.component) || {slot:false}
    let currentComponent = components[page.component] || NotFound
    let currentSlotted = currentMeta.slot

    console.log("Loading component:", page.component)
</script>

<svelte:component this={currentComponent} {...page.settings} {...$$restProps}>
{#if currentSlotted}
    {#each page.children as child}
    <svelte:self page={child}></svelte:self>
    {/each}
{/if}
</svelte:component>