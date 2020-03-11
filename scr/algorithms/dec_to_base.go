package algorithms

import (
	"fmt"
	"strconv"
)

//DecimalToBase
func DecimalToBase(num int, base int) string {
	convertedNum := ""
	if num == 0 {
		return "0"
	} else if num < base {
		return string(num)
	} else {
		for keepDividing := true; keepDividing; keepDividing = (num > 0) {

			switch num % base {
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
				convertedNum = strconv.Itoa(num%base) + convertedNum
			}

			fmt.Println(convertedNum)
			num = num / base
		}
	}
	return convertedNum
}
