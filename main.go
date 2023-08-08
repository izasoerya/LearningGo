package main

import (
	"fmt"

	soalprakerja "github.com/izasoerya/Prakerja-Go/soal_prakerja"
)

func main() {
	fmt.Print("The Value is : ")
	fmt.Println(soalprakerja.DetectPrimeNumber(0))
	fmt.Println(soalprakerja.Detect7Multiplication(42))
	fmt.Println(soalprakerja.CalculateAreaTrapeziod(2,4,3))
}
