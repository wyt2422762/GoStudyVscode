package maximumproduct

//公平的糖果棒交换
func fairCandySwap(A []int, B []int) []int {
	var res []int = make([]int, 2)
	sumA := sum(A)
	sumB := sum(B)

	for _, vA := range A {
		for _, vB := range B {
			if abs(sumA-sumB) == abs(vA-vB)*2 {
				res[0] = vA
				res[1] = vB
				return res
			}
		}
	}

	return res
}

//优化版本-公平的糖果棒交换
func fairCandySwap1(A []int, B []int) []int {
	var res []int = make([]int, 2)
	sumA := sum(A)
	sumB := sum(B)

	mapA := make(map[int]int, len(A))
	for _, vA := range A{
		mapA[vA] = 1
	}

	for _, vB := range B{
		vvA := (sumA - sumB) / 2 + vB
		if _, has := mapA[vvA];has{
			res[0] = vvA
			res[1] = vB
			return res
		}
	}
	
	return res
}

//计算数组的和
func sum(nums []int) int {
	var res int
	for _, v := range nums {
		res += v
	}
	return res
}

//计算int的绝对值
func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}
