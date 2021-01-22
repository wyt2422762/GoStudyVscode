package main

import (
	"fmt"
	"sort"
)

func searchRange(nums []int, target int) []int {

	i := sort.SearchInts(nums, target)
	fmt.Println(i)
	return []int{-1, -1}

}

func main() {
	i := sort.SearchInts([]int{1, 2, 3}, 4)
	fmt.Println(i)
}
