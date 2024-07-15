package models

import (
	"study.com/golang-web/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Amount      int
}

func FindAllProducts() []Product {
	db := db.Connect()

	result, err := db.Query("select * from product order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for result.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = result.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Amount = amount

		products = append(products, p)
	}

	defer db.Close()

	return products
}

func InsertProduct(name, description string, price float64, amount int) {
	db := db.Connect()

	stmt, err := db.Prepare("insert into product(name, description, price, amount) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	stmt.Exec(name, description, price, amount)

	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.Connect()

	stmt, err := db.Prepare("delete from product where id = $1")
	if err != nil {
		panic(err.Error())
	}

	stmt.Exec(id)

	defer db.Close()
}

func FindProductById(id string) Product {
	db := db.Connect()

	result, err := db.Query("select * from product where id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	p := Product{}

	for result.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = result.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Amount = amount
	}

	defer db.Close()

	return p
}

func UpdateProduct(id, name, description string, price float64, amount int) {
	db := db.Connect()

	stmt, err := db.Prepare("update product set name = $1, description = $2, price = $3, amount = $4 where id = $5")
	if err != nil {
		panic(err.Error())
	}

	stmt.Exec(name, description, price, amount, id)

	defer db.Close()
}
