package maximumproduct

// 给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

// 示例：
// 输入："Let's take LeetCode contest"
// 输出："s'teL ekat edoCteeL tsetnoc"

// 提示：
// 在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。

func reverseWords1(s string) string {

    len := len(s)

    runes := []rune(s)

    left, right :=0, 0

    for right < len{
        for right < len && runes[right] != ' '{
            right++
        }

        l, r := left, right-1

        for l < r {
            runes[l], runes[r] = runes[r], runes[l]
            l++
            r--
        }

        left, right = right+1, right+1
    }
    
	ss := string(runes[0 : len])

    return ss
}