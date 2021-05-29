package maximumproduct

import (
	"fmt"
	"strconv"
)

// 罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。

// 字符          数值
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000

// 例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

// 通常情况下，罗马数字中小的数字在大的数字的右边。
// 但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。
// 同样地，数字 9 表示为 IX。
// 这个特殊的规则只适用于以下六种情况：
// I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
// X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
// C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。

// 1 <= num <= 3999

// 示例 1:
// 输入: num = 3
// 输出: "III"

// 示例 2:
// 输入: num = 4
// 输出: "IV"

// 示例 3:
// 输入: num = 9
// 输出: "IX"

// 示例 4:
// 输入: num = 58
// 输出: "LVIII"
// 解释: L = 50, V = 5, III = 3.

// 示例 5:
// 输入: num = 1994
// 输出: "MCMXCIV"
// 解释: M = 1000, CM = 900, XC = 90, IV = 4

func intToRoman(num int) (res string) {
	change := make(map[int]string, 7)
	change[1] = "I"
	change[5] = "V"
	change[10] = "X"
	change[50] = "L"
	change[100] = "C"
	change[500] = "D"
	change[1000] = "M"

	str := fmt.Sprint(num)
	byt := []byte(str)
	le := len(byt)

	res = ""

	for i, v := range byt {
		// fmt.Printf("%d: %s\n", i, string(v))
		zeros := le - 1 - i //用来补0
		s := string(v)

		str_int, _ := strconv.Atoi(s)

		switch str_int {
		case 1, 5:
			res += change[str_int*getTens(zeros)]
		case 2, 3:
			for ii := 0; ii < str_int; ii++ {
				res += change[1*getTens(zeros)]
			}
		case 6, 7, 8:
			res += change[5*getTens(zeros)]
			for ii := 0; ii < str_int-5; ii++ {
				res += change[1*getTens(zeros)]
			}
		case 4:
			res += change[1*getTens(zeros)]
			res += change[5*getTens(zeros)]
		case 9:
			res += change[1*getTens(zeros)]
			res += change[10*getTens(zeros)]
		default:
		}
	}
	return
}

func getTens(zeros int) (res int) {
	res = 1
	for i := 0; i < zeros; i++ {
		res = res * 10
	}
	return
}

