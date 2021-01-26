package maximumproduct

import (
	"fmt"
)

//给你一个由一些多米诺骨牌组成的列表 dominoes。
//如果其中某一张多米诺骨牌可以通过旋转 0 度或 180 度得到另一张多米诺骨牌，我们就认为这两张牌是等价的。
//形式上，dominoes[i] = [a, b] 和 dominoes[j] = [c, d] 等价的前提是 a==c 且 b==d，或是 a==d 且 b==c。
//在 0 <= i < j < dominoes.length 的前提下，找出满足 dominoes[i] 和 dominoes[j] 等价的骨牌对 (i, j) 的数量。
func numEquivDominoPairs(dominoes [][]int) int {
	fmt.Println("numEquivDominoPairs")
	if dominoes == nil || len(dominoes) <= 1 {
		return 0
	}
	uMap := make(map[int]int) //用来接收等价的数量，key为数字按由小到大排列，value为等价的数量
	for i := 0; i < len(dominoes); i++ {
		temp := uMap[buildKey(dominoes[i])] //获取对应key的数据，不存在默认为0
		uMap[buildKey(dominoes[i])] = temp +1
	}
	res := 0
	for _, val := range uMap{ //遍历map
		res += (val * (val - 1)) / 2 //根据数量计算有多少对
	}
	return res
}

//组装map用到的key
func buildKey(nums []int) int{
	if nums[0] > nums[1] {
		return nums[0] * 10 + nums[1]
	}
	return nums[1] * 10 + nums[0]
}
