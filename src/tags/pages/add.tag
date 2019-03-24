add
    h3 Add page
    hr
    form
        .row
            .col
                .form-group
                    label Name:
                    input.form-control(type="text", ref="name", placeholder="My new markdown page")
            .col
                .form-group
                    label Type:
                    select.form-control(ref="type")
                        option(value="markdown") Markdown (Github)
        .row
            .col
                .form-group
                    label Url:
                    input.form-control(type="text", ref="url", placeholder="my-new-markdown-page", title="This is the url your page will be bound to, ex: /p/my-new-markdown-page.")
        .row
            .col
                .form-group
                    label Content:
                    textarea.form-control(ref="content", rows="15")
        .row
            .col
                a.btn.btn-outline-primary(tabindex="4") Save