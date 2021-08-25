package render

import (
	"bytes"
	"fmt"
	"go_blog/internal/config"
	"go_blog/internal/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

func SetConfigForTemplates(a *config.AppConfig) {
	app = a
}

var functions = template.FuncMap{}

//GoRenderTemplate renders a template
func GoRenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, data *models.TemplateData) {
	var templateCache map[string]*template.Template
	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	//Getting template from template cache
	t, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("Couldn't get template from template cache")
	}
	buf := new(bytes.Buffer)
	_ = t.Execute(buf, data)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error in writing template to the browser: ", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		//Get file base name
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
