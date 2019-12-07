package main

import "fmt"

func main() {
	arr := [...]int{21, 34, 5, 23, 67}
	n := len(arr)
	fmt.Println("--------没排序前--------\n", arr)
	for i := 0; i < n; i++ {
		fmt.Println("--------第", i+1, "次冒泡--------")
		for j := i; j < n; j++ {
			if arr[i] > arr[j] {
				t := arr[i]
				arr[i] = arr[j]
				arr[j] = t
			}
			fmt.Println(arr)
		}
	}
	fmt.Println("--------最终结果--------\n", arr)
}
