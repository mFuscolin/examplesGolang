package main

import (
	"fmt"
	"time"

	"github.com/mFuscolin/projeto-fiado/model"
)

func main() {

	customer := model.Customer{}
	customer.ID = 1
	customer.Name = "Aldo"
	customer.Documents.RG = "19859158s5"
	customer.Documents.CPF = "59489041050"

	product1 := model.Product{}
	product2 := model.Product{}
	products := []model.Product{}

	product1.ID = 1234
	product1.Name = "Litrão"
	product1.Value = 8

	product2.ID = 12345
	product2.Name = "Litrão Schin"
	product2.Value = 12

	products = append(products, product1)
	products = append(products, product2)

	debt := model.Debt{}
	debt.IDCustomer = customer.ID
	debt.Customer = customer
	debt.Products = products

	debt.DueDate = time.Now().AddDate(0, 1, 0)

	fmt.Printf("O cliente: %+v\r\n", debt)

}
