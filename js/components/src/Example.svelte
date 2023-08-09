<script>
    import { Form, Text, Checkbox, Radio, Select, Navbar } from './index'
  
    let options = [
      {value:'spam', label:'spam'},
      {value:'not spam', label:'not spam'}
    ]
    let nav = [
      {
        label:'I',
        href:'/'
      },{
        label:'Can',
        href:'/'
      },{
        label:'Haz',
        href:'/'
      }
    ]
  </script>
  
  <main>
    <Navbar items={nav} brand="Asdf" />
    <Form method="post" action="/" format="form" back="/" title="Test!">
      <Text kind="text" name="spam" label="Spam" />
      <Text kind="password" name="password" label="Password" />
      <Text kind="number" name="number" label="Number" />
      <Text kind="textarea" name="spam" label="Spam" />
      <Text kind="text" label="Empty" placeholder="Empty" />
      <Text kind="other" label="Plain text password" name="password" />
      <Checkbox label="Checkboxes with spam" name="checks"items={options} />
      <Radio label="Radiobuttons with spam" name="spam" items={options} />
      <Select label="Dropdown with spam" items={options} name="spam" />
    </Form>
  </main>