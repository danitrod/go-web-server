package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/danitrod/go-web-server/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.GetProducts()
	templates.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Price float conversion err:", err)
		}

		qtdInt, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Quantity int conversion err:", err)
		}

		models.InsertProduct(name, description, priceFloat, qtdInt)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Id int conversion err:", err)
	}
	models.DeleteProduct(idInt)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Id int conversion err:", err)
	}
	product := models.GetProduct(idInt)

	templates.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Id int conversion err:", err)
		}
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Price float conversion err:", err)
		}

		qtdInt, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Quantity int conversion err:", err)
		}

		models.UpdateProduct(idInt, name, description, priceFloat, qtdInt)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
