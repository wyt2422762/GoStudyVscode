package maximumproduct

import (
	"strings"
	"strconv"
)

/*
[4,2,3]
[4,2,1]
[3,4,2,3]
[4,2,3,5]
*/

// 给你一个长度为 n 的整数数组，请你判断在 最多 改变 1 个元素的情况下，该数组能否变成一个非递减数列。
// 我们是这样定义一个非递减数列的： 对于数组中所有的 i (0 <= i <= n-2)，总满足 nums[i] <= nums[i + 1]。

// 示例 1:
// 输入: nums = [4,2,3]
// 输出: true
// 解释: 你可以通过把第一个4变成1来使得它成为一个非递减数列。

// 示例 2:
// 输入: nums = [4,2,1]
// 输出: false
// 解释: 你不能在只改变一个元素的情况下将其变为非递减数列。
func checkPossibility(nums []int) bool {
	le := len(nums)
	if le <= 2 {
		return true
	}

	nms := make(map[string]int) //记录递减的次数
	//maps := make(map[int]int) //记录最小值的次数

	for i := 1; i < le-1; i++ {
		if nums[i] > nums[i+1] || nums[i-1] > nums[i] { //出现了递减
			if nums[i] > nums[i+1] {
				if _, ok := nms[strconv.Itoa(i)+"-"+strconv.Itoa(i+1)]; !ok {
					nms[strconv.Itoa(i)+"-"+strconv.Itoa(i+1)] = nums[i] - nums[i+1]
				}
			}
			if nums[i-1] > nums[i] {
				if _, ok := nms[strconv.Itoa(i-1)+"-"+strconv.Itoa(i)]; !ok {
					nms[strconv.Itoa(i-1)+"-"+strconv.Itoa(i)] = nums[i-1] - nums[i]
				}
			}
			if len(nms) > 1 {
				return false
			}
		}
	}

	if len(nms) > 1 {
		return false
	} else if len(nms) == 1 {
		for key := range nms {
			start, _ := strconv.Atoi(strings.Split(key, "-")[0])
			end, _ := strconv.Atoi(strings.Split(key, "-")[1])

			if start == 0 || end == le-1 {
				return true
			}
			if nums[end] < nums[start-1] && nums[end+1] < nums[start] {
				return false
			}
			return true
		}
		return false
	} else {
		return true
	}
}
