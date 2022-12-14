package main

import (
	"database/sql"

	db2 "github.com/Vitoriano/go-hexagonal/adapters/db"
	"github.com/Vitoriano/go-hexagonal/application"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite3")
	productDbAdapter := db2.NewProductDb(db)

	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product Example", 30)

	productService.Enable(product)

}
