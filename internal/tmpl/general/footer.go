package tmpl

import "html/template"

func Footer() template.HTML {
	return template.HTML(`
	<footer class="flex">
      <div class="container">
        <a
          class="link"
          target="_blank"
          href="https://github.com/ropfoo/gothulhu"
          >Github</a
        >
      </div>
    </footer>
	`)
}
