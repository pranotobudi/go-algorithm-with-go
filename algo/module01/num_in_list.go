package module01

// NumInList will return true if num is in the
// list slice, and false otherwise.
func NumInList(list []int, num int) bool {
	// fmt.Println("list: ", list, "num: ", num)
	if list == nil {
		return false
	}
	for _, numInList := range list {
		// fmt.Println("numInList: ", numInList)
		if numInList == num {
			return true
		}
	}
	return false
}
