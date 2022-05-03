package tpl

import (
	"embed"
	"fmt"
	"text/template"
)

var (
	//go:embed templates
	res   embed.FS
	files = map[string]string{
		"invoice-en-mono": "templates/invoice-en-mono.gtpl",
		"invoice-he-mono": "templates/invoice-he-mono.gtpl",
	}
)

func LoadTemplates() (map[string]*template.Template, error) {
	templates := map[string]*template.Template{}

	for k, v := range files {
		tmpl, err := template.ParseFS(res, v)
		if err != nil {
			return nil, err
		}

		templates[k] = tmpl
	}

	return templates, nil
}

func LoadTemplate(tmplName string) (*template.Template, error) {
	var (
		tmpl  *template.Template
		err   error
		found bool = false
	)

	for k, v := range files {
		if k == tmplName {
			tmpl, err = template.ParseFS(res, v)
			if err != nil {
				return nil, err
			}
			found = true
		}
	}

	if !found {
		return nil, fmt.Errorf("template='%s' not found", tmplName)
	}

	return tmpl, nil
}
