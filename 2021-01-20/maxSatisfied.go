package maximumproduct

// 今天，书店老板有一家店打算试营业 customers.length 分钟。每分钟都有一些顾客（customers[i]）会进入书店，所有这些顾客都会在那一分钟结束后离开。
// 在某些时候，书店老板会生气。 如果书店老板在第 i 分钟生气，那么 grumpy[i] = 1，否则 grumpy[i] = 0。 当书店老板生气时，那一分钟的顾客就会不满意，不生气则他们是满意的。
// 书店老板知道一个秘密技巧，能抑制自己的情绪，可以让自己连续 X 分钟不生气，但却只能使用一次。
// 请你返回这一天营业下来，最多有多少客户能够感到满意的数量。

/* 示例：
输入：customers = [1,0,1,2,1,1,7,5], grumpy = [0,1,0,1,0,1,0,1], X = 3
输出：16 */

func maxSatisfied(customers []int, grumpy []int, X int) int {
	res := 0

	le := len(grumpy)
	count := getCount(customers)
	baseSatisfiedCount := getBaseSatisfiedCount(customers, grumpy)

	if le <= X {
		res = count
	} else {
		//[0, 1, 2, 3, 4]
		maxIncrease := 0 //最大新增满意客户数
		//计算新增满意客户数
		for start := 0; start < le-X+1; start++ {
			increase := 0 //新增满意客户数
			for j := 0; j < X; j++ {
				if grumpy[start+j] == 1 {
					increase += customers[start+j]
				}
			}
			maxIncrease = max(maxIncrease, increase)
		}
		res = baseSatisfiedCount + maxIncrease
	}

	return res
}

//求顾客总数
func getCount(customers []int) int {
	res := 0
	for _, v := range customers {
		res += v
	}
	return res
}

//求不改动的情况下的满意的客户数
func getBaseSatisfiedCount(customers []int, grumpy []int) int {
	res := 0
	for key, v := range grumpy {
		if v == 0 {
			res += customers[key]
		}
	}
	return res
}
