
riot.tag2('add', '<h3>Add page</h3> <hr> <form> <div class="row"> <div class="col"> <div class="form-group"> <label>Name:</label> <input class="form-control" type="text" ref="name" placeholder="My new markdown page"> </div> </div> <div class="col"> <div class="form-group"> <label>Type:</label> <select class="form-control" ref="type"> <option value="markdown">Markdown (Github)</option> </select> </div> </div> </div> <div class="row"> <div class="col"> <div class="form-group"> <label>Url:</label> <input class="form-control" type="text" ref="url" placeholder="my-new-markdown-page" title="This is the url your page will be bound to, ex: /p/my-new-markdown-page."> </div> </div> </div> <div class="row"> <div class="col"> <div class="form-group"> <label>Content:</label> <textarea class="form-control" ref="content" rows="15"></textarea> </div> </div> </div> <div class="row"> <div class="col"><a class="btn btn-outline-primary" tabindex="4">Save</a></div> </div> </form>', '', '', function(opts) {
});

riot.tag2('list', '<div class="list-group"> <div class="list-group-item"> <div class="input-group"> <input class="form-control" type="text"> <div class="input-group-amend"><a class="btn btn-outline-primary" href="#/add">+</a></div> </div> </div> <div class="list-group-item list-group-item-action d-flex justify-content-between align-items-center">Page1<a class="btn btn-link" href="#">Remove</a></div> <div class="list-group-item list-group-item-action d-flex justify-content-between align-items-center">Page2<a class="btn btn-link" href="#">Remove</a></div> </div>', '', '', function(opts) {
});

riot.tag2('page', '<span></span>', '', '', function(opts) {
    this.on('mount', function() {

        var converter = new showdown.Converter()
        converter.setFlavor('github')
        var md = '# Hello world!'
        this.root.innerHTML = converter.makeHtml(md)
    })
});