package maximumproduct

func removeStones(stones [][]int) int {

	

	res := 0
	for i := 0; i < len(stones)-1; i++ {
		stone1 := stones[i]
		for j := 0; j < len(stones); j++ {
			if i == j {
				continue
			}
			stone2 := stones[j]
			if stone1[0] == stone2[0] || stone1[1] == stone2[1] {
				res++
				break
			}
		}
	}
	return res
}
