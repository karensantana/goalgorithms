package algorithms

//NumInList unction that takes an array and looks for one number, returns true if the number is found, otherwise false
func NumInList(list []int, num int) bool {
	//When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.
	// With _ you are skipping the index value, just using the current value of the slice
	for _, elementAt := range list {
		if elementAt == num {
			return true
		}
	}
	return false
}
