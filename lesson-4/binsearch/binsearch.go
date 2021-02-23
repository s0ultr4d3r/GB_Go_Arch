package binsearch

func binSearch(array []int, target int, lowIndex int, highIndex int) int {
	if highIndex < lowIndex {
		return -1
	}
	mid := int((lowIndex + highIndex) / 2)
	if array[mid] > target {
		return binSearch(array, target, lowIndex, mid)
	} else if array[mid] < target {
		return binSearch(array, target, mid+1, highIndex)
	} else {
		return mid
	}
}

func iterbinSearch(array []int, target int, lowIndex int, highIndex int) int {
	startIndex := lowIndex
	endIndex := highIndex

	var mid int

	for startIndex < endIndex {
		mid = int((lowIndex + highIndex) / 2)
		if array[mid] > target {
			return binSearch(array, target, lowIndex, mid)
		} else if array[mid] < target {
			return binSearch(array, target, mid+1, highIndex)
		} else {
			return mid
		}

	}
	return -1
}
