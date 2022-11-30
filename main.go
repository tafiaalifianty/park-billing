package main

import (
	"exercise-go-fundamentals/park"
	"fmt"
)

func runParkBilling() {
	price := park.CalculateParkingBill(park.Car, 47)
	fmt.Println("price", price)
}

func main() {
	runParkBilling()
}
