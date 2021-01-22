package maximumproduct

func maximumproduct(coordinates [][]int) bool {
	if len(coordinates) <= 2 {
		return true
	}
	for i := 1; i < len(coordinates)-1; i++ {
		x := (coordinates[i][0] - coordinates[i-1][0]) * (coordinates[i+1][1] - coordinates[i][1])
		y := (coordinates[i+1][0] - coordinates[i][0]) * (coordinates[i][1] - coordinates[i-1][1])
		if x != y {
			return false
		}
	}
	return true
}
