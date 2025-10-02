package products

import (
	"go-crud/database"
)

type Product struct {
	Id, Quantity      int
	Name, Description string
	Price             float64
}

func GetAll() []Product {
	/*
		Abre a conexão com o Postgres e fecha no fim da função após executá-la com o comando "defer"
	*/
	db := database.Connect()

	/*
		Executa a query "SELECT * FROM products"
	*/
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		panic(err.Error())
	}

	product := Product{}
	products := []Product{}

	for rows.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = rows.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity

		products = append(products, product)
	}

	defer db.Close()
	return products
}

func Insert(name, description string, price float64, quantity int) {
	db := database.Connect()

	dbInsert, err := db.Prepare("INSERT INTO products(name, description, price, quantity) VALUES($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	dbInsert.Exec(name, description, price, quantity)
	defer db.Close()
}
