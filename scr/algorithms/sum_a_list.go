package algorithms

//SumAList a list of integers
func SumAList(list []int) int {
	if len(list) == 0 {
		return 0
	}
	return list[0] + SumAList(list[1:])
}
