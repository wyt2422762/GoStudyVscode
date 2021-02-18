package maximumproduct

// 给出一个由二维数组表示的矩阵，以及两个正整数r和c，分别表示想要的重构的矩阵的行数和列数。
// 重构后的矩阵需要将原始矩阵的所有元素以相同的行遍历顺序填充。
// 如果具有给定参数的reshape操作是可行且合理的，则输出新的重塑矩阵；否则，输出原始矩阵。

/* 示例 1:
输入:
nums =
[[1,2],
 [3,4]]
r = 1, c = 4
输出:
[[1,2,3,4]]
解释:
行遍历nums的结果是 [1,2,3,4]。新的矩阵是 1 * 4 矩阵, 用之前的元素值一行一行填充新矩阵。 */

/* 示例 2:
输入:
nums =
[[1,2],
 [3,4]]
r = 2, c = 4
输出:
[[1,2],
 [3,4]]
解释:
没有办法将 2 * 2 矩阵转化为 2 * 4 矩阵。 所以输出原矩阵。 */
func matrixReshape(nums [][]int, r int, c int) [][]int {
	rNums, cNums := len(nums), len(nums[0]) //原矩阵的行数, 原矩阵的列数

	if rNums*cNums != r*c {
		return nums
	}

	hang := 0
	lie := 0

	res := make([][]int, r) //初始化一个行数为r的二维数组
	for i := range res {
		res[i] = make([]int, c) //二维数组的内部是一个长度为c的数组

		for j := range res[i] {
			res[i][j] = nums[hang][lie]

			if lie == cNums-1 {
				hang++
				lie = 0
			} else {
				lie++
			}
		}
	}
	return res
}

func matrixReshape1(nums [][]int, r int, c int) (res [][]int) {
	rNums, cNums := len(nums), len(nums[0]) //原矩阵的行数, 原矩阵的列数

	if rNums*cNums != r*c {
		return nums
	}

	res = make([][]int, r) //初始化一个行数为r的二维数组

	for i := range res {
		res[i] = make([]int, c) //二维数组的内部是一个长度为c的数组
	}

	for i := 0; i < r*c; i++ {
		res[i/c][i%c] = nums[i/cNums][i%cNums]
	}

	return res
}
