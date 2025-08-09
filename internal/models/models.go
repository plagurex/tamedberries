package models

import (
	"html/template"
)

type Page struct {
	Title   string
	Content template.HTML
}

type Category struct {
	Id   int
	Name string
}

type Product struct {
	Id          int
	Name        string
	Description string
	CategoryId  int `db:"category_id"`
	Img         string
}
