package main

import (
	"fmt"
	"TRAINING_04112024/utils"
)

func main() {

	result := MagicSum(5)
	fmt.Println("Nilai Akhir MagicSum", result)
	m := 4 // Example input
	result1 := MagicPow(m)
	fmt.Println("Nilai Akhir MagicPow", result1)
	l := 6 // Example input
	result2 := MagicOdd(l)
	fmt.Println("Nilai Akhir MagicOdd", result2)
	k := 3 // Example input
	result3 := MagicGrade(k)
	fmt.Println("Nilai Akhir MagicGrade", k, result3)
	o := 1 // Example input
	result4 := MagicName(o)
	fmt.Println("Nilai Akhir MagicName", result4)
	p := 3 // Example input
	result5 := MagicTria(p)
	fmt.Println("Nilai Akhir MagicTria", result5)
	r := 5 // Example input
	fmt.Printf("Before change: %d\n", r)
	MagicChange(&r) // Pass the address of n
	fmt.Printf("After change: %d\n", r)
}
