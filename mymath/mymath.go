package mymath

import (
	"fmt"
	"math"
)

// Sqrt - функция для вычисления квадратного корня из числа
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

// Pow - функция для возведения числа в степень
func Pow(x float64, n float64) float64 {
	return math.Pow(x, n)
}

// Sin - функция для вычисления синуса угла в радианах
func Sin(x float64) float64 {
	return math.Sin(x)
}

func main() {

	fmt.Println("Квадратный корень из 4:", mymath.Sqrt(4))       // 2
	fmt.Println("2 в степени 3:", mymath.Pow(2, 3))              // 8
	fmt.Println("Синус 0.5 (угол в радианах):", mymath.Sin(0.5)) // 0.479425538604203
}
