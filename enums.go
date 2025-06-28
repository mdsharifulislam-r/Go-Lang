package main

import "fmt"

type OrderStatus string

const (
	PENDING OrderStatus = "pending"
	CONFIRM             = "confirmed"
	CANCEL              = "cencel"
	DELVER              = "delver"
)

func changeOrderStatus(status OrderStatus) {

	fmt.Println(status)

}

func main() {
	changeOrderStatus(PENDING)
}
