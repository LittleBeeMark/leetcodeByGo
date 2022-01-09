package binary_search

import "sort"

// 二分查找成立条件
// 1.目标函数具有单调性（单调递增或是单调递减）
// 2.存在上下界
// 3.能过通过索引访问

//// go 语言的二分查找模板
//func BinaryModel(array []int, target int) int {
//	left, right := 0, len(array)-1
//	for left <= right {
//		mid := (left + right) / 2
//		if array[mid] == target {
//			return mid
//		}
//		if array[mid] > target {
//			right = mid - 1
//		} else {
//			left = mid + 1
//		}
//	}
//	return -1
//}

// 1552. 两球间的磁力
// https://leetcode-cn.com/problems/magnetic-force-between-two-balls/
// 使用二分查找的思想不断地尝试区间中点是否合适并根据判断缩减区间
// 需要注意的是因为要取的是最大值，所以一旦区间中点为合适应把它直接付给 min 而不能加一也就是必须是min = mid 非 min = mid + 1
// 因为相等(count==m)的时刻是必取的，相等即证明最小距离其实就是mid, 此时代码把 min = mid ,若min=mid+1 	则正解被跳过，
// 后续查找将一致都不通过，从而右边界不断向左压缩到min,	也无法找到正解
// 第二点要注意的是 终结循环条件一定是 	min < max 	如果写了等号，那么一旦找到唯一正解 mid, 	右边界将会不断压缩直到min == max
// 此时check是通过的又再次把min = mid ,形成死循环
func maxDistance(position []int, m int) int {
	sort.Ints(position)
	min, max := 1, position[len(position)-1]-position[0]
	var mid int
	for min < max {
		mid = min + (max-min+1)/2
		if checkLen(position, m, mid) {
			min = mid
		} else {
			max = mid - 1
		}
	}

	return min

}

func checkLen(position []int, m int, mid int) bool {
	pre, count := position[0], 1
	for i := 1; i < len(position); i++ {
		if position[i]-pre >= mid {
			count++
			pre = position[i]
		}
	}

	return count >= m

}
