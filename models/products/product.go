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

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity

		products = append(products, product)
	}

	defer db.Close()
	return products
}

func GetById(id string) Product {
	db := database.Connect()

	rows, err := db.Query("SELECT * FROM products WHERE id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	product := Product{}
	for rows.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = rows.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity
	}

	defer db.Close()
	return product
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

func Delete(id string) {
	db := database.Connect()

	dbDelete, err := db.Prepare("DELETE FROM products WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}

	dbDelete.Exec(id)
	defer db.Close()
}

func Update(id string, name, description string, price float64, quantity int) {
	db := database.Connect()

	dbUpdate, err := db.Prepare("UPDATE products SET name=$1, description=$2, price=$3, quantity=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}

	dbUpdate.Exec(name, description, price, quantity, id)
	defer db.Close()
}
