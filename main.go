package main

import (
	"net/http"

	"github.com/danitrod/go-web-server/db"
	"github.com/danitrod/go-web-server/routes"
)

func main() {
	db.SetupSQLite()
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
