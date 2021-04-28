package main

import "fmt"

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

// 示例:
// 输入: [0,1,0,3,12]
// 输出: [1,3,12,0,0]

// 说明:
// 必须在原数组上操作，不能拷贝额外的数组。
// 尽量减少操作次数。

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
}

// 1. 双指针，slow从0开始 fast从1开始
// 2. 判断slow对应的值是否为0， 不为0，slow++，为0则执行步骤3
// 3. 判断fast对应的值是否为0，不为0，交换slow, fast对应的值，为0则不交换
// 4. fast++
func moveZeroes(nums []int) {
	len := len(nums)

	if len <= 1 {
		return
	}

	slow, fast := 0, 1 //双指针(快慢)

	for ; fast < len; fast++ {
		//首先判断nums[slow]是不是0
		if nums[slow] != 0 {
			slow++
		} else if nums[fast] != 0 {
			//交换数据
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}

	fmt.Println(nums)
}
