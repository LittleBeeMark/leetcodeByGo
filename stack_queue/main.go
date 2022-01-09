package main

import (
	"fmt"
	"math"
	"strings"
)

// stack 栈
// 特性先入后出 无序
// 删除，添加 时间复杂度O(1)
// 查询时间复杂度 O(n)

// queue 队列
// 特性先入先出 无序
// 删除，添加 时间复杂度O(1)
// 查询时间复杂度 O(n)

// dequeue 双端队列
// 可以从前放入前拿出，后放入后拿出
// 删除，添加 时间复杂度O(1)
// 查询时间复杂度 O(n)
//

// priority queue 优先队列
// 取出操作 log（n)
// 	放入操作 O(n)
// 具有优先级，是一个相对有序的数列
// 底层实现的数据结构可以多种多样，如heap binary-tree

// 何时使用栈，从内到外，从内到外，逐渐扩散 最外层和最外层为一对，最内层和最内层为一对 具有一致性，叫做最近相关性
// 先来后到用 queue
// 有效括号
// https://leetcode-cn.com/problems/valid-parentheses/
// 暴力解法
// 每次进行是否有括号的判断，如果可以匹配，则去掉进行下一次匹配，这样写其实时间复杂度达到了O(n^3)
func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	for strings.Contains(s, "()") || strings.Contains(s, "[]") || strings.Contains(s, "{}") {
		if strings.Contains(s, "()") {
			s = strings.ReplaceAll(s, "()", "")
		}
		if strings.Contains(s, "[]") {
			s = strings.ReplaceAll(s, "[]", "")
		}
		if strings.Contains(s, "{}") {
			s = strings.ReplaceAll(s, "{}", "")
		}

	}

	return len(s) == 0
}

// 栈的方法
// 栈的方法如何理解以及这个例子为何可以使用栈呢
// 首先有效括号指的是左右括号对应且左右相对位置要一致，比如第一位为zuox
func StackIsValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	match := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	stack := []byte{}

	for i := 0; i < len(s); i++ {
		if v, ok := match[s[i]]; ok {
			stack = append(stack, v)
			continue
		}

		if len(stack) < 0 || stack[len(stack)-1] != s[i] {
			fmt.Println(len(stack))
			fmt.Println(match[stack[len(stack)-1]] != s[i])
			fmt.Println("stack ", string(stack[len(stack)-1]), string(s[i]))
			return false
		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}

// 最小栈
// https://leetcode-cn.com/problems/min-stack/
type MinStack struct {
	Stack []int
	Min   []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Stack: []int{},
		Min:   []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.Stack = append(this.Stack, val)
	if val < this.Min[len(this.Min)-1] {
		this.Min = append(this.Min, val)
	} else {
		this.Min = append(this.Min, this.Min[len(this.Min)-1])
	}
}

func (this *MinStack) Pop() {
	this.Min = this.Min[:len(this.Min)-1]
	this.Stack = this.Stack[:len(this.Stack)-1]
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]

}

func (this *MinStack) GetMin() int {
	return this.Min[len(this.Min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

//https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
//84. 柱状图中最大的矩形
// 暴力法 这是纯粹的暴力枚举方法，提交以后超时了，这个暴力枚举是人脑最简单思维方式，时间复杂度达到O(n^4)
func largestRectangleArea(heights []int) int {
	var sMax = gMaxest(heights)
	for i := 0; i <= len(heights)-2; i++ {
		for j := i + 1; j <= len(heights)-1; j++ {
			sMax = gMax((j-i+1)*getMinest(heights[i:j+1]), sMax)
		}
	}

	return sMax
}

func getMinest(a []int) int {
	var min = math.MaxInt64
	for i := 0; i < len(a); i++ {
		if a[i] < min {
			min = a[i]
		}
	}
	return min
}

func gMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func gMaxest(a []int) int {
	var max int
	for i := 0; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max

}

// 暴力解法2 寻找最大边界
// 这个暴力解法其实是上一个的优化，换了一种遍历每个值的最大边界的解法，时间复杂度为O(n^2)
// 但提交还是超时，别人python 都可以
func largestRectangleArea2(heights []int) int {
	var res int
	for i := 0; i < len(heights); i++ {
		currentH := heights[i]
		left := i
		for left > 0 {
			if heights[left-1] < currentH {
				break
			}
			left--
		}

		right := i
		for right < len(heights)-1 {
			if heights[right+1] < currentH {
				break
			}

			right++
		}

		res = gMax(currentH*(right-left+1), res)
	}

	return res
}

func main() {
	//fmt.Println("answer:", StackIsValid("(([)])"))
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
