menu
    .list-group
        virtual(each="{ items }")
            a.list-group-item.list-group-item-action(href="{href}") {name}
    script.
        this.items = [{href:'/p/', name:'Pages'}]