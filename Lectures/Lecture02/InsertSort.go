package Lecture02

//Insert Sort Algorithms
func InsertSort(inputs []int) {
	for index := 1; index < len(inputs); index++ {
		temp := inputs[index]
		leftindex := index - 1
		for ; leftindex >= 0 && inputs[leftindex] > temp; leftindex-- {
			inputs[leftindex+1] = inputs[leftindex]
		}
		inputs[leftindex+1] = temp
	}
}
