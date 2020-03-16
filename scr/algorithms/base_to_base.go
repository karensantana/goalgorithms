package algorithms

import (
	"fmt"
	"math"
	"strconv"
)

//DecimalToBase
func BaseToBase(num string, oldBase int, newBase int) string {
	baseTenNum := 0

	//Convert number to base 10 as a starting point
	if oldBase == 10 {
		n, err := strconv.Atoi(num)
		if err != nil {
			return "An error has occured!"
		} else {
			baseTenNum = n
		}
	} else {
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
		for i := len(num) - 1; i >= 0; i-- {

			baseTenNum += m[string(num[i])] * int(math.Pow(float64(oldBase), float64(exponetial)))
			exponetial++
		}
	}
	fmt.Println("Converted dec: ", baseTenNum)
	//Convert to newBase
	convertedNum := ""
	if baseTenNum == 0 {
		return "0"
	} else if baseTenNum < newBase {
		return strconv.Itoa(baseTenNum)
	} else {
		for keepDividing := true; keepDividing; keepDividing = (baseTenNum > 0) {

			switch baseTenNum % newBase {
			case 10:
				convertedNum = "A" + convertedNum
			case 11:
				convertedNum = "B" + convertedNum
			case 12:
				convertedNum = "C" + convertedNum
			case 13:
				convertedNum = "D" + convertedNum
			case 14:
				convertedNum = "E" + convertedNum
			case 15:
				convertedNum = "F" + convertedNum
			default:
				convertedNum = strconv.Itoa(baseTenNum%newBase) + convertedNum
			}

			baseTenNum = baseTenNum / newBase
		}
	}
	return convertedNum

}
