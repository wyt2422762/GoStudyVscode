package maximumproduct

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