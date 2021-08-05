package main

import (
	"fmt"
	"unicode"
)

/**
增加一个新字符串，遍历原来的字符串：
如果非空格则将原来的字符直接拼接到新字符串中；
如果遇到空格则将%20拼接到新字符串中

重点：熟练掌握go的字符串替换，拼接
*/
func replaceSpace(s string) string {
	var newString []rune
	// []rune(【字符串】)，将【字符串】转化为unicode编码的切片
	replace := []rune("%20")
	// v的类型是rune，即int32，unicode编码
	for _, v := range s {
		// 使用 unicode.IsSpace() 判断是否是空格
		if unicode.IsSpace(v) {
			// 拼接多个数组，需加上三个句点...
			newString = append(newString, replace...)
		} else {
			newString = append(newString, v)
		}
	}
	return string(newString)
}

func main() {
	//s := "中 国 We are happy."
	s := "We are happy."
	//go标准库自带的字符串替换函数 string.ReplaceAll()，但题目要求你写算法，不要用内置的函数
	//strings.ReplaceAll(s, " ", "%20")
	newString := replaceSpace(s)
	fmt.Println(newString)
}
