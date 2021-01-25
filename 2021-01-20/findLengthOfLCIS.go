package maximumproduct

import (
	"fmt"
)

//最长且 连续递增的子序列的长度
func findLengthOfLCIS(nums []int) int {
	fmt.Println("findLengthOfLCIS")
	if len(nums) == 0 || nums == nil{
		return 0
	}
	var res int = 1
	var i, j int = 0, 0
	for ;i<len(nums)-1;{
		j = i +1
		startFlag := i
		for j < len(nums){
			if nums[j] > nums[i]{
				i = j
				j++
			} else {
				i = j
				break
			}
		}
		res = max(res, j-startFlag)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
