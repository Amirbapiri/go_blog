package handlers

import (
	"go_blog/internal/config"
	"go_blog/internal/models"
	"go_blog/internal/render"
	"net/http"
)

//Repo is a pointer to the repository that used by the handlers
var Repo *Repository

//Repository is the repository type to hold a pointer to app configuration
type Repository struct {
	App *config.AppConfig
}

func CreateNewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//SetRepository sets the repository for the handlers
func SetRepository(r *Repository) {
	Repo = r
}

//Home is the handler for the home page
func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	data["author"] = "Amir"

	render.GoRenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{
		StringMap: data,
	})
}

func (repo *Repository) Register(w http.ResponseWriter, r *http.Request) {
	render.GoRenderTemplate(w, r, "register.page.tmpl", &models.TemplateData{})
}
