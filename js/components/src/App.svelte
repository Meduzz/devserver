<script>
  import Components from './lib/loaders/Components.svelte'
  import { writable } from 'svelte/store'

  let form = writable({
    text1:'Hello world!',
    select1: 'item2'
  })
  let page = {
    component: "page",
    settings: {},
    children: [
      {
        component: 'navbar',
        settings:{
          items: [
            {
              label: 'Home',
              href: '/'
            },{
              label: 'Databases',
              href: '/dbs'
            },{
              label: 'Pages',
              href: '/pages'
            }
          ],
          brand: 'CMS'
        }
      },
      {
        component: 'routable',
        settings: {
          path: '/'
        },
        children: [
          {
            component: "form",
            settings: {
              method: 'POST',
              action: '/',
              title: 'Test!',
              format:'form',
              back:"/",
              redirect: '/',
              store: form
            },
            children: [
              {
                component:'text',
                settings: {
                  kind:'text',
                  label: 'Some text',
                  name: 'text1',
                  placeholder: 'Supplies...'
                }
              },{
                component: 'select',
                settings: {
                  label: 'Some select',
                  name: 'select1',
                  items: [
                    {
                      label: 'Item1',
                      value: 'item1'
                    },{
                      label: 'Item2',
                      value: 'item2'
                    }
                  ]
                }
              }
            ]
          }
        ]
      },
      {
        component: 'routable',
        settings: {
          path: '/dbs/*'
        },
        children: [
          {
            component: 'list',
            settings: {
              items: [
                {
                  label: 'Dbs'
                }
              ]
            }
          }
        ]
      },
      {
        component: 'routable',
        settings: {
          path: '/pages/*'
        },
        children: [
          {
            component: 'list',
            settings: {
              items: [
                {
                  label: 'Pages'
                }
              ]
            }
          }
        ]
      },
    ]
  }
</script>

<Components {page} />
