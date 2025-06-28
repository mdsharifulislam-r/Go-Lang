package main

import (
	"fmt"
)

type order struct {
	id    string
	price int
}

func (o *order) changePrice(price int) {

	o.price = price

}

func main() {
	orderItem := order{
		id:    "sharif",
		price: 3,
	}

	orderItem.changePrice(30)

	fmt.Println(orderItem.price)
}
