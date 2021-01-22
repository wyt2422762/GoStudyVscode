package main

import "fmt"

// 给定一个非负整数数组 A， A 中一半整数是奇数，一半整数是偶数。
// 对数组进行排序，以便当 A[i] 为奇数时，i 也是奇数；当 A[i] 为偶数时， i 也是偶数。
// 你可以返回任何满足上述条件的数组作为答案。

// 示例：

// 输入：[4,2,5,7]
// 输出：[4,5,2,7]
// 解释：[4,7,2,5]，[2,5,4,7]，[2,7,4,5] 也会被接受。
func sortArrayByParityII(A []int) []int {
	//fmt.Println("按奇偶排序数组 II")

	length1 := len(A)

	for i := 0; i < length1; i++ {
		if (A[i]-i)%2 != 0 { //不符合条件
			//fmt.Println("奇偶特性不同")
			for j := (i + 1); j < length1; j++ {
				if (A[j]-i)%2 == 0 { //不符合条件
					//fmt.Println("交换")
					A[i], A[j] = A[j], A[i]
					break
				}
			}
		}
	}
	return A
}

func sortArrayByParityII2(A []int) ([]int, []string) {
	//fmt.Println("按奇偶排序数组 II")

	length1 := len(A)

	for i, j := 0, 1; i < length1; i += 2 { //i代表偶数index, j代表奇数index
		if A[i]%2 != 0 { //不是偶数
			for A[j]%2 != 0 {
				j += 2
			}
			A[i], A[j] = A[j], A[i]
		}
	}
	return A, nil
}

func main() {
	A := []int{4,1,1,0,1,0}
	A, _ = sortArrayByParityII2(A)
	fmt.Printf("%d", A)
}
