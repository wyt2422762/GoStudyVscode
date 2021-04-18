package maximumproduct

import (
	"fmt"
)

// 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

//定义两个指针fast和slow分别为快指针和慢指针，快指针表示遍历数组到达的下标位置，慢指针表示下一个不同元素要填入的下标位置，初始时两个指针都指向下标 1。
// 假设数组 nums 的长度为 nn。将快指针fast依次遍历从1到n−1的每个位置，对于每个位置，如果nums[fast] =nums[fast−1]，
// 说明 nums[fast] 和之前的元素都不同，因此将 nums[fast] 的值复制到 nums[slow]，然后将 slow 的值加 1，即指向下一个位置。
// 遍历结束之后，从nums[0]到nums[slow−1]的每个元素都不相同且包含原数组中的每个不同的元素，因此新的长度即为slow，返回 slow 即可。

func removeDuplicates(nums []int) int {
	lenOld := len(nums)
	fmt.Printf("原始长度为%d\n", lenOld)

	if 0 == lenOld {
		return 0
	}

	slow := 1
	for fast := 1; fast < lenOld; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}
