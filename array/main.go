package main

import (
	"fmt"
	"math"
	"sort"
)

// 数组，链表， 跳表

// 数组 开辟了一段连续的内存地址，每一个地址都可以通过下标直接访问，访问，前追加，后追加时间复杂度为O(1),
// 因为地址连续所以增加，删除时间复杂度为O(n)

// 链表 通过后指针前指针把一段数据串联在一起，访问时间复杂度为O(n), 其他操作为O(1)

// 跳表 只能针对有序数据，插入删除搜索都为O(log(n))时间复杂度  这里跳表采用了升维，空间换时间的思想加速有序数据的查找

//  给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//示例:
//
//输入: [0,1,0,3,12]
//输出: [1,3,12,0,0]
//说明:
//
//必须在原数组上操作，不能拷贝额外的数组。
//尽量减少操作次数。

// 理解： 从0下标开始遇到非零元素累计下标，遇到零元素后：继续向后遍历并不断把非零元素向累计到的非零下标放置，把对应位置改写为零
// https://leetcode-cn.com/problems/move-zeroes/
func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if j != i {
				nums[j] = nums[i]
				nums[i] = 0
			}
			j++
		}
	}
}

// 盛水最多的容器
// 理解
// https://leetcode-cn.com/problems/container-with-most-water/
// 	暴力解法
func maxArea(height []int) int {
	var h, max, r int
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {

			if height[i] <= height[j] {
				h = height[i]
			} else {
				h = height[j]
			}

			r = h * (j - i)
			if max < r {
				max = r
			}
		}
	}
	return max
}

func gmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func gmin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 盛水最多的容器
// 理解 : 控制变量长为最长，看有没有更高的高(左右夹逼）
// https://leetcode-cn.com/problems/container-with-most-water/
func maxArea2(height []int) int {
	var max int
	left, right := 0, len(height)-1
	for right > left {
		max = gmax(gmin(height[right], height[right])*(right-left), max)
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return max

	//var h, max, r int
	//i := 0
	//j := len(height) - 1
	//for i < j {
	//	if height[i] <= height[j] {
	//		h = height[i]
	//		i++
	//	} else {
	//		h = height[j]
	//		j--
	//	}
	//
	//	// 比对应长缩小了一个，所以要加一
	//	r = h * (j - i + 1)
	//	if max < r {
	//		max = r
	//	}
	//}
	//return max
}

// 动态规划
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func climbStairs2(n int) int {
	if n <= 3 {
		return n
	}

	f1, f2, f3 := 1, 2, 3
	for i := 3; i < n; i++ {
		f2 = f3
		f1 = f2
		f3 = f1 + f2
	}

	return f3
}

// 无重复字符的最长子串
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
// 滑动窗口解法，其实这个滑动窗口的解法，很多解释都不清不楚，包括画图法之类的，
// 其实就是以每个位置都为起始位，不断的向后获取不重复的终止位，最后获取最大的那个起始位对应的不重复数值
// 这样的时间复杂度就接近n^2
// 优化方式就是找到重复字母的下标，并以下标后一个为起始位，思考为什么可以这样？
// 因为每次遇到重复的字符其实当前的重复位置下标的前一个非重复位置到当前起始位置的这个窗口，
// 是当前起始位置可到被重复字符下标的下一个字符的下标位置的最大长度位置，
// 因为此时从当前起始为到被重复字符位置的最大右界都是固定的(重复位置下标的前一个非重复位置)，此时前面的位置都是可以不要的
// 这样就可以达到时间复杂度为n
// 遇到的问题： 无法界定每次挪动最大长度，遇到重复的时候，应当记录左窗口下标变化后的左下标到当前右下标，
// 不重复的时候应该记录当前的右下标（左下标不变）
func lengthOfLongestSubstring(s string) int {
	dpMap := make(map[byte]int, 0)
	sum, start, end := 0, 0, 0
	for end < len(s) {

		if v, ok := dpMap[s[end]]; ok {
			start = gMax(start, v+1)
		}

		dpMap[s[end]] = end
		sum = gMax(end-start+1, sum)
		end++
	}

	return sum
}

func gMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 两数之和
// https://leetcode-cn.com/problems/two-sum/
// 暴力解法
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// 哈希表法
func twoSum2(nums []int, target int) []int {
	targetMap := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		if v, ok := targetMap[target-nums[i]]; ok {
			return []int{v, i}
		}
		targetMap[nums[i]] = i
	}
	return []int{}
}

// 三数之和
//  https://leetcode-cn.com/problems/3sum/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := [][]int{}

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		start := i + 1
		end := n - 1
		for start < end {
			if start > i+1 && nums[start] == nums[start-1] {
				start++
				continue
			}

			if nums[start]+nums[end] == -nums[i] {
				res = append(res, []int{nums[i], nums[start], nums[end]})
				start++
				end--
				continue
			}

			if nums[start]+nums[end] > -nums[i] {
				end--
			} else {
				start++
			}

		}
	}

	return res

}

// 两数相加
//

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	var move, current int

	for l1 != nil || l2 != nil {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := v1 + v2 + move
		current, move = sum%10, sum/10
		if head == nil {
			head = &ListNode{
				Val: current,
			}
			tail = head
		} else {
			tail.Next = &ListNode{
				Val: current,
			}
			tail = tail.Next

		}

	}

	if move > 0 {
		tail.Next = &ListNode{
			Val: move,
		}
	}
	return
}

// 5. 最长回文子串
// https://leetcode-cn.com/problems/longest-palindromic-substring/
func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	start, maxLen, maxStr := 0, 0, string(s[0])

	for start < len(s)-1 {
		for end := len(s) - 1; start < end; end-- {
			if IsTrue(s[start : end+1]) {
				if maxLen < end-start+1 {
					maxLen = end - start + 1
					maxStr = s[start : end+1]
				}

				break
			}
			end--
		}

		start++
	}

	return maxStr
}

func IsTrue(s string) bool {
	head, tail := 0, len(s)-1
	for head < tail {
		if s[head] != s[tail] {
			return false
		}

		head++
		tail--
	}

	return true
}

func reverse(x int) int {

	var res int
	for x != 0 {
		if res > math.MaxInt32/10 || res < math.MinInt32/10 {
			return 0
		}

		digit := x % 10
		x = x / 10
		res = res*10 + digit
	}

	return res
}

//给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
//
//你返回所有和为 0 且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。
func threeSum2(nums []int) [][]int {
	for i := 0; i < len(nums); i++ {
		numHash := make(map[int]int, 2)
		numHash[i] = nums[i]
		for j := i + 1; j < len(nums); j++ {

		}

	}

}
func main() {
	//fmt.Println(lengthOfLongestSubstring("abba"))
	//	fmt.Println(bytes.Index([]byte("chicken"), []byte("ken")))
	//fmt.Println(bytes.Index([]byte("chicken"), []byte("dmr")))
	//fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(reverse(123))
}
