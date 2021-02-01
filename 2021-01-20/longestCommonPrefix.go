package maximumproduct

import (
	"sort"
)

type cusStrs []string //自定义一个类型为字符串数组类型的类型，用来自定义排序

//获取长度
func (cs cusStrs) Len() int {
	return len(cs)
}

//交换数据
func (cs cusStrs) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

//自定义比较方法，这里按字符串的长度比较
func (cs cusStrs) Less(i, j int) bool {
	return len(cs[i]) < len(cs[j])
}

//最长公共前缀
func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) <= 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	sort.Sort(cusStrs(strs)) //根据自定义排序数组
	str0 := strs[0]
	if len(str0) == 0 {
		return ""
	}

	var resIndex int = -1
	//根据长度最短的字符串来判断最长公共前缀
	outFlag:
	for i := range str0 {
		sl := strs[1:]
		for _, cl := range sl{
			if(str0[i] != cl[i]){
				break outFlag
			}
		}
		resIndex = i
	}
	r := []rune(str0)
	return string(r[0 : resIndex+1])
}


