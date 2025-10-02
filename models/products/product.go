package products

import (
	"go-crud/database"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
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
