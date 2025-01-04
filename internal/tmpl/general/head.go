package tmpl

import "html/template"

type HeadParams struct {
	Title       string
	Description string
	Stylesheets []string
}

func Head(params HeadParams) template.HTML {
	links := ""
	for _, link := range params.Stylesheets {
		links += `<link rel="stylesheet" href="` + link + `">` + "\n"
	}

	return template.HTML(`
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta name="description" content="` + params.Description + `">

		<link rel="icon" type="image/x-icon" href="/favicon.ico">
		<link rel="stylesheet" href="/styling/main.css" />
		` + links + `
        
		<title>` + params.Title + `</title>
	`)
}
