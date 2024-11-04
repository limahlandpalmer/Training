package utils

import (
	"fmt"
	"math"
	"strconv"
)

func MagicSum(n int) int {
	return n + n
}
func MagicPow(m int) float64 {
	return math.Pow(float64(m), float64(m))
}

func MagicOdd(l int) bool {
	return l%2 != 0
}

func MagicGrade(k int) string {
	switch k {
	case 0:
		return "Zero"
	case 1:
		return "Bad"
	case 2:
		return "Ok"
	case 3:
		return "Nice"
	case 4:
		return "Awesome"
	case 5:
		return "Exceptional"
	default:
		return "Unknown"
	}
}

func MagicName(o int) []string {
	name := "Adi" // Replace with your name
	names := make([]string, o)

	for i := 0; i < o; i++ {
		names[i] = name
	}

	return names
}

func MagicTria(p int) string {
	sum := 0
	terms := make([]string, p)

	for i := 1; i <= p; i++ {
		sum += i
		terms[i-1] = strconv.Itoa(i) // Convert integer to string
	}

	return fmt.Sprintf("%s = %d", fmt.Sprintf("%s", terms), sum)
}
func MagicChange(r *int) {
	*r = *r * 2
}
