package main

// 最长子序列
// https://leetcode-cn.com/problems/longest-common-subsequence/
// 该题的暴力解法位使用双递归求出来两个字符串的所有子序列，进行双循环一一比对，取出最长的相同子序列
// 使用动态规划的方法，可以根据去除尾部元素的子序列来递推出来
// 如果两个子序列的尾部元素相同那么说明最后一个元素肯定是属于最长子序的，
// 可以把这两个序列的尾部元素都去掉，形成的两个序列的最长子胥加一。
// 可以把两个子序列的递归树构思出来，如果最后一个元素相等，那么最长子序一定在最下层的叶子节点，
// 一定是上层最长的相同子序列加一得到的
// 如果两个字符串尾部元素不等，那么相同的最长子序一定不在两个递归树的最后一层，
// 所以要取某一个字符串去掉尾部元素和另一个字符串产生的最长子序的最大值
func longestCommonSubsequence(text1 string, text2 string) int {
	// 状态容器
	m := len(text1)
	n := len(text2)
	dp := make([][]int, n+1)
	for k := 0; k < m; k++ {
		dp[k] = make([]int, m+1)
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if text1[j-1] == text2[i-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = gMax(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[n][m]

}

func gMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 打家劫舍
// https://leetcode-cn.com/problems/house-robber/
func rob(nums []int) int {
	if nums == nil {
		return 0
	}

	dp := make([][]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0
	dp[0][1] = nums[0]

	for j := 1; j < len(nums); j++ {
		dp[j][0] = gMax(dp[j-1][1], dp[j-1][0])
		dp[j][1] = dp[j-1][0] + nums[j]
	}

	return gMax(dp[len(nums)-1][0], dp[len(nums)-1][1])
}

func rob2(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))

	dp[0] = nums[0]
	dp[1] = gMax(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = gMax(dp[i-2], dp[i-1]+nums[i])
	}

	return dp[len(nums)-1]
}

func main() {

}
