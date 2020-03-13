package algorithms

import (
	"math"
)

//DecimalToBase
func BaseToDecimal(num string, base int) int {

	m := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"A": 10,
		"B": 11,
		"C": 12,
		"D": 13,
		"E": 14,
		"F": 15,
	}

	exponetial := 0
	result := 0

	for i := len(num) - 1; i >= 0; i-- {

		result += m[string(num[i])] * int(math.Pow(float64(base), float64(exponetial)))
		//fmt.Println(result)
		exponetial++
	}
	return result
}
