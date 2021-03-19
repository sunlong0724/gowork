package main

import (
	"fmt"
	"strings"
)

/*
请实现一个算法，确定一个字符串的所有字符【是否全都不同】。这里我们要求【不允许使用额外的存储结构】。 给定一个string，请返回一个bool值,true代表所有字符全都不同，false代表存在相同的字符。 保证字符串中的字符为【ASCII字符】。字符串的长度小于等于【3000】。
*/
func isUniqueString(s string) bool {
	if len(s) > 256 {
		return false
	}
	for i := 0; i < len(s); i++{
		for j := 0; j < len(s); j++{
			if i == j {
				continue
			}
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}
/*
请实现一个算法，在不使用【额外数据结构和储存空间】的情况下，翻转一个给定的字符串(可以使用单个过程变量)。
给定一个string，请返回一个string，为翻转后的字符串。保证字符串的长度小于等于5000。
 */
func reverString(s string)(string ,bool){
	str := []rune(s)
	l := len(str)

	for i := 0; i < l/2; i++{
		str[i],str[l-i-1] = str[l-i-1],str[i]
	}
	return string(str),true
}
/*
判断两个给定的字符串排序后是否一致
问题描述
给定两个字符串，请编写程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。 这里规定【大小写为不同字符】，且考虑字符串重点空格。给定一个string s1和一个string s2，请返回一个bool，代表两串是否重新排列后可相同。 保证两串的长度都小于等于5000。
 */
func isRegroup(s1,s2 string) bool{
	for _,v := range s1{
		if strings.Count(s1, string(v)) != strings.Count(s2,string(v)){
			return false
		}
	}
	return true
}
/*
符串替换问题
问题描述
请编写一个方法，将字符串中的空格全部替换为“%20”。 假定该字符串有足够的空间存放新增的字符，并且知道字符串的真实长度(小于等于1000)，同时保证字符串由【大小写的英文字母组成】。 给定一个string为原始的串，返回替换后的string。
 */
func replaceBlank(s string) (string, bool) {
	return strings.Replace(s," ", "%20", -1), true
}

/*
机器人坐标问题
问题描述
有一个机器人，给一串指令，L左转 R右转，F前进一步，B后退一步，问最后机器人的坐标，最开始，机器人位于 0 0，方向为正Y。 可以输入重复指令n ： 比如 R2(LF) 这个等于指令 RLFLF。 问最后机器人的坐标是多少？
 */
func move(l,r,f,b int){

}
func main(){
	fmt.Println( isUniqueString("abcdef"))
	fmt.Println( reverString("abcdefghijklmn我爱你中国"))
	fmt.Println(replaceBlank("1 2 3 4  "))
}
