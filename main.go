package main

import (
	"fmt"

	"github.com/izasoerya/Prakerja-Go/controller"
)

func main() {
	fmt.Print("The Value is : ")
	fmt.Println(controller.Calculate(12, 15))
	fmt.Println("Conflict A again now")
}
