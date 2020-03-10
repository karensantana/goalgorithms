package algorithms

//ReverseString: takes an string and returns it reversed
func ReverseString(str string) string {
	if len(str) == 0 {
		return ""
	} else if len(str) == 1 {
		return string(str[0])
	}
	return string(str[len(str)-1]) + ReverseString(str[0:(len(str)-1)])
}
