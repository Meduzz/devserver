
riot.tag2('menu', '<div class="list-group"> <virtual each="{items}"><a class="list-group-item list-group-item-action" href="{href}">{name}</a></virtual> </div>', '', '', function(opts) {
this.items = [{href:'/p/', name:'Pages'}]
});