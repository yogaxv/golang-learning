package main

import (
	"fmt"
	"pertama/calculation"
)

func main() {
	fmt.Println("Hello World")

	kalimat := TestAja()
	fmt.Println(kalimat)

	numberOne := 5
	numberTwo := 5

	resultAdd := calculation.Add(numberOne,numberTwo)
	fmt.Println(fmt.Sprintf("hasil penjumlahan dari %d dan %d adalah %d ", numberOne, numberTwo, resultAdd))

	resultMultiply := calculation.Multiply(numberOne, numberTwo)
	fmt.Printf("hasil perkalian dari %d dan %d adalah %d", numberOne, numberTwo, resultMultiply)
}
