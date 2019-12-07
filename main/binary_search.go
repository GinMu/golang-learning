package main

import "fmt"

// BinarySearch search value from sorted data
func BinarySearch(arr *[]int, leftIndex int, rightIndex int, findVal int) {
	if leftIndex > rightIndex {
		fmt.Println("could not find")
		return
	}
	middleIndex := (leftIndex + rightIndex) / 2
	middleVal := (*arr)[middleIndex]
	if middleVal > findVal {
		BinarySearch(arr, leftIndex, middleIndex-1, findVal)
	} else if middleVal < findVal {
		BinarySearch(arr, middleIndex+1, rightIndex, findVal)
	} else {
		fmt.Printf("had found, index is: %v\n", middleIndex)
	}
}

func main() {
	arr := []int{1, 2, 5, 7, 9, 10, 34, 35, 100}
	BinarySearch(&arr, 0, len(arr)-1, 100)
	fmt.Println("main arr=", arr)
}
