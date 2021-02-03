package main

import "fmt"

func binarySearch(s []int, target int) int {
	var i = 0
	var j = len(s) - 1

	for {
		if i >= j {
			return -1
		}
		mid := (i + j) / 2
		if s[mid] == target {
			return mid
		}
		if s[mid] > target {
			i = i
			j = mid
			continue
		}
		if s[mid] < target {
			i = mid + 1
			j = j
			continue
		}
	}
	return -1
}
func main() {
	s := []int{1, 2, 4, 5, 6, 7, 8}
	i := binarySearch(s, 0)
	fmt.Println(i)
}
