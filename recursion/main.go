package main

import (
	"fmt"
	"strings"
)

// 	递归万能模板
//const (
//	MaxLevel int = 128
//)
//
//func Recursion(level int, param1, param2 interface{}) interface{} {
//	// 终止条件，最终子节点的值返出
//	var res string
//	if level >= MaxLevel {
//		return res
//	}
//
//	// 逻辑处理
//	process(level, param1,param2)
//	level++
//
//	// dril down 下探
//	return Recursion(level	,param1,param2)

//  // 清理数据（不必要）
//}

// Z 字变换
// https://leetcode-cn.com/problems/zigzag-conversion/
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	resNums := make([]string, numRows)
	flag := -1
	i := 0
	for _, sb := range s {
		if i == 0 || i == len(resNums)-1 {
			flag = -flag
		}

		resNums[i] += string(sb)
		i += flag
	}

	return strings.Join(resNums, "")
}

// 周期法
func convert2(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	// 2n-2 为一个周期
	resultRow := make([]string, numRows)

	period := 2*numRows - 2
	for i, v := range s {
		p := i % period
		if p >= numRows {
			p = numRows - 1 - (p - numRows + 1)
		}

		resultRow[p] += string(v)
	}

	return strings.Join(resultRow, "")
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	numMap := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	var res []string
	recur(0, digits, "", numMap, &res)
	return res
}

func recur(index int, digits, target string, numMap map[string]string, res *[]string) {
	if index > len(digits)-1 {
		*res = append(*res, target)
		return
	}

	if v, ok := numMap[string(digits[index])]; ok {
		for _, num := range v {
			recur(index+1, digits, target+string(num), numMap, res)
		}
	} else {
		recur(index+1, digits, target, numMap, res)
	}

}

func main() {
	fmt.Println("res: ", letterCombinations("23"))
}
