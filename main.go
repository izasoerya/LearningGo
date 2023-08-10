package main

import (
	"fmt"

	soalprakerja "github.com/izasoerya/Prakerja-Go/soal_prakerja"
)

func main() {
	//* Sesi 2
	// fmt.Println(soalprakerja.DetectPrimeNumber(0))
	// fmt.Println(soalprakerja.Detect7Multiplication(42))
	// fmt.Println(soalprakerja.CalculateAreaTrapeziod(2, 4, 3))

	//* Sesi 3
	fmt.Println(soalprakerja.ArrayMerge(
		[]string{"king", "devil", "akuma"},
		[]string{"eddie", "steve", "geese"}))
	fmt.Println(soalprakerja.ArrayMerge(
		[]string{"sergei", "jin"},
		[]string{"jin", "steve", "brayn"}))
	fmt.Println(soalprakerja.ArrayMerge(
		[]string{"alisha", "yoshi"},
		[]string{"jin", "yoshi", "alisha", "law"}))
	fmt.Println(soalprakerja.ArrayMerge(
		[]string{},
		[]string{"jin", "sergei"}))
	fmt.Println(soalprakerja.ArrayMerge(
		[]string{"hworang"},
		[]string{}))
	fmt.Println(soalprakerja.Mapping(
		[]string{"asd", "pdm", "asd", "pln", "pdm", "pki", "pdm"}))
	fmt.Println(soalprakerja.MunculSekali("1234123"))
	fmt.Println(soalprakerja.MunculSekali("76523752"))
	fmt.Println(soalprakerja.MunculSekali("12345"))
	fmt.Println(soalprakerja.MunculSekali("1122334455"))
	fmt.Println(soalprakerja.MunculSekali("0872504"))
}
