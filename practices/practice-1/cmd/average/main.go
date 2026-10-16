package main

import "fmt"

func main() {
	minutes1 := 20
	minutes2 := 25
	minutes3 := 35

	sum := minutes1 + minutes2 + minutes3
	avg := float64(sum) / 3

	fmt.Printf("Сумма: %d\n", sum)
	fmt.Printf("Ср. знач: %.2f", avg)
}
