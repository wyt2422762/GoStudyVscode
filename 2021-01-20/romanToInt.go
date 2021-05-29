package maximumproduct

// 罗马数字中小的数字在大的数字的右边。若输入的字符串满足该情况，那么可以将每个字符视作一个单独的值，累加每个字符对应的数值即可。
// 例如 XXVII 可视作 X + X + V + I + I = 10 + 10 + 5 + 1 + 1 = 27 

// 若存在小的数字在大的数字的左边的情况，根据规则需要减去小的数字。
// 对于这种情况，我们也可以将每个字符视作一个单独的值，若一个数字右侧的数字比它大，则将该数字的符号取反。
// 例如 XIV 可视作 X - I + V = 10 - 1 + 5 = 14

func romanToInt(s string) int {
	res := 0

	data := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
	}

	byt := []byte(s)
	le := len(byt)

	for i := 0; i < le; i++{
		str := string(byt[i])
		value := data[str]
        if i < le-1 && value < data[string(byt[i+1])] {
            res -= value
        } else {
            res += value
        }
	}

	return res
}
