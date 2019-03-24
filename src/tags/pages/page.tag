page
    span

    script.
        this.on('mount', function() {
            // TODO load opts.page.
            var converter = new showdown.Converter()
            converter.setFlavor('github')
            var md = '# Hello world!'
            this.root.innerHTML = converter.makeHtml(md)
        })
