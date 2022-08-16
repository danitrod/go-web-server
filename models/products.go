package models

import (
	"log"

	"github.com/danitrod/go-web-server/db"
)

type Product struct {
	Name, Description string
	Price             float64
	Id, Quantity      int
}

func GetProducts() []Product {
	connection := db.ConnectToDB()
	defer connection.Close()

	selectProductsQuery, err := connection.Query("SELECT * FROM products")
	if err != nil {
		log.Fatal(err.Error())
	}

	products := []Product{}

	for selectProductsQuery.Next() {
		var id, quantity int
		var name, description string
		var price float64
		err = selectProductsQuery.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			log.Fatal(err.Error())
		}
		products = append(products, Product{Id: id, Name: name, Description: description, Price: price, Quantity: quantity})
	}

	return products
}

func InsertProduct(name, description string, price float64, quantity int) {
	connection := db.ConnectToDB()
	defer connection.Close()

	insertProductQuery, err := connection.Prepare("INSERT INTO products (name, description, price, quantity) VALUES ($1, $2, $3, $4)")
	if err != nil {
		log.Fatal(err.Error())
	}

	insertProductQuery.Exec(name, description, price, quantity)
}

func DeleteProduct(id int) {
	connection := db.ConnectToDB()
	defer connection.Close()

	deleteProductQuery, err := connection.Prepare("DELETE FROM products WHERE id = $1")
	if err != nil {
		log.Fatal(err.Error())
	}

	deleteProductQuery.Exec(id)
}
