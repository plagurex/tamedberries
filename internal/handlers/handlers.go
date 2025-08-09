package handlers

import (
	"html/template"
	"net/http"
	"strings"
	"tb/internal/app"
	"tb/internal/models"
	"tb/internal/utils"
)

func HomeHandler(a *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			f := NotFound(a)
			f(w, r)
			return
		}
		// data := models.Page{
		// 	Title:   "TamedBerries - интернет магазин",
		// 	Content: template.HTML("Лютый дроч"),
		// }
		// utils.SendHtmlResponse(w, a, data)
		http.Redirect(w, r, "/search", http.StatusMovedPermanently)
	}
}

func AboutUsHandler(a *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := models.Page{
			Title:   "О нас",
			Content: template.HTML("Мы дауны ыыы"),
		}
		utils.SendHtmlResponse(w, a, data)
	}
}

func CatalogHandler(a *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := a.LoadTemplate("catalog.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		items := make([]models.Category, 0)

		err = a.DB().Select(&items, "SELECT * FROM categories")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		html, err := utils.RenderTemplateToHTML(t, items)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := models.Page{
			Title:   "Каталог",
			Content: html,
		}
		utils.SendHtmlResponse(w, a, data)
	}
}

func SearchHandler(a *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := a.LoadTemplate("search.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		items := make([]models.Product, 0)

		query := r.URL.Query()
		category := query.Get("category")
		s := "SELECT * FROM products"
		if category == "" {
			err = a.DB().Select(&items, s)
		} else {
			s += " WHERE category_id = $1"
			err = a.DB().Select(&items, s, category)
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for i, v := range items {
			if v.Img == "" {
				items[i].Img = "nophoto.png"
			}
		}

		html, err := utils.RenderTemplateToHTML(t, items)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := models.Page{
			Title:   "Поиск",
			Content: html,
		}
		utils.SendHtmlResponse(w, a, data)
	}
}

func ProductHandler(a *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := a.LoadTemplate("product.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var product models.Product

		s, _ := strings.CutPrefix(r.URL.Path, "/product/")
		err = a.DB().Get(&product, "SELECT * FROM products WHERE id = $1", s)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if product.Img == "" {
			product.Img = "nophoto.png"
		}

		html, err := utils.RenderTemplateToHTML(t, product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := models.Page{
			Title:   product.Name,
			Content: html,
		}
		utils.SendHtmlResponse(w, a, data)
	}
}

func NotFound(a *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := a.LoadTemplate("not-found.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		html, err := utils.RenderTemplateToHTML(t, r.URL.Path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := models.Page{
			Title:   "404 Не найдено",
			Content: html,
		}
		utils.SendHtmlResponse(w, a, data)
	}
}
