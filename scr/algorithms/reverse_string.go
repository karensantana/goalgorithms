package algorithms

import "strings"

//ReverseStringRecursively: takes an string and returns it reversed
func ReverseStringRecursively(str string) string {
	if len(str) == 0 {
		return ""
	} else if len(str) == 1 {
		return string(str[0])
	}
	return string(str[len(str)-1]) + ReverseString(str[0:(len(str)-1)])
}

//ReverseString
func ReverseString(str string) string {
	var sb strings.Builder

	for i := len(str) - 1; i >= 0; i-- {
		sb.WriteByte(str[i])
	}

	return sb.String()
}
