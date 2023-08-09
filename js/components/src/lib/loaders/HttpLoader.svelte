<script>
    export let method = 'GET'
    export let url
    export let format = 'json'
    export let body

    let content = contentType(format)
    let req = {
        method,
        headers:{
            'Content-Type': content
        }
    }

    if (method != 'GET' && body) {
        switch (format) {
            case 'json':
                req.body = JSON.stringify(body)
            case 'form':
                let form = new FormData()
                Object.keys(body).forEach(it => {
                    form.append(it, body[it])
                })
                req.body = form
            default:
        }
    }

    let promise = fetch(req, url)

    function contentType(type) {
        switch (type) {
            case 'json':
                return 'application/json'
            case 'form':
                return 'application/x-www-form-urlencoded'
            default:
                return 'text/plain'
        }
    }
</script>

{#await promise}
    <p>Loading...</p>
{:then result}
    <slot data={result}></slot>
{/await}
