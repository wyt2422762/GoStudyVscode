package main

// 给出一个字符串 s（仅含有小写英文字母和括号）。
// 请你按照从括号内到外的顺序，逐层反转每对匹配括号中的字符串，并返回最终的结果。
// 注意，您的结果中 不应 包含任何括号。

// 示例 1：
// 输入：s = "(abcd)"
// 输出："dcba"

// 示例 2：
// 输入：s = "(u(love)i)"
// 输出："iloveu"

// 示例 3：
// 输入：s = "(ed(et(oc))el)"
// 输出："leetcode"

// 示例 4：
// 输入：s = "a(bcdefghijkl(mno)p)q"
// 输出："apmnolkjihgfedcbq"


func reverseParentheses(s string) string {
	stack := [][]byte{} //用二位数组模拟栈
    str := []byte{}

	for i, c := range s {
		if c == '(' {
			//遇到左括号，将字符压入栈中
			stack = append(stack, str)
			str = []byte{}
		} else if c == ')' {
			//遇到右括号，将字符串翻转，并将栈顶元素放到首位
			//翻转字符串
			for j, n := 0, len(str); j < n/2; j++ {
                str[j], str[n-1-j] = str[n-1-j], str[j]
            }
			//栈顶元素放到首位
			str = append(stack[len(stack)-1], str...)
			//用切片模拟出栈操作
			stack = stack[:len(stack)-1]
		} else {
			//字符数组拼接
			str = append(str, s[i])
		}
	}

	return string(str)
}
