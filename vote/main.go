package main

import "fmt"

// ## 主要元素
//
//数组中占比超过一半的元素称之为主要元素。给你一个 整数 数组，找出其中的主要元素。若没有，返回 -1 。请设计时间复杂度为 O(N) 、空间复杂度为 O(1) 的解决方案
// 摩尔投票法
// 	理解：这个看似很难理解的摩尔投票法，可以理解成乱军厮杀，每个数组中的数字都代表一方势力，哪一方势力可以力压群雄，那么那一方就胜出
//  比如： [1,1,2,2,2,2,3,2] 可以把1，2，3 理解为蜀国，魏国和吴国。一个国家要胜出那么这个国家的数字一定比另外两个加起来还多，
//  从1开始，每次出现相同就说明蜀国势力加一，这时默认蜀国是最强国，魏国吴国会来怼他，出现2或3都怼掉一个1，也就是势力减一，减到0说明蜀国的势力已经怼掉完了，
//  此时下次出现的数字2，就是目前默认最强国魏国，蜀国吴国来攻击他，同理加减。厮杀到最后剩下的一定是最强国，比其他几个加起来还多，
//  但要注意的是，如果这个数组本身就没有最强势力（主要元素），那么最后出现的数字将成为伪最强国，此时要把最强候选国的数量计算一下看看他的两倍是不是大于总数量
//  如果题目本身就告诉你一定有一个主要元素那就不需要了
//
func Vote(nums []int) int {
	//
	//cnt := 1
	//var candidate int
	//for _, num := range nums {
	//	if cnt == 0 {
	//		candidate = num
	//	}
	//
	//	if candidate == num {
	//		cnt++
	//	} else {
	//		cnt--
	//	}
	//}
	//
	//var count int
	//for _, num := range nums {
	//	if num == candidate {
	//		count++
	//	}
	//}
	//
	//if 2*count > len(nums) {
	//	return candidate
	//}
	//return -1
	var candidate int
	vote := 1
	for _, n := range nums {
		if vote == 0 {
			candidate = n
		}

		if candidate == n {
			vote++
		} else {
			vote--
		}
	}

	var count int
	for _, num := range nums {
		if num == candidate {
			count++
		}
	}

	if 2*count > len(nums) {
		return candidate
	}

	return -1
}

func main() {
	testNums := [][]int{
		{1, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5, 6},
		{1, 2, 4, 5, 7},
		{1, 1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 2, 2},
	}

	for i, nums := range testNums {
		fmt.Printf("index: %d, most num: %d \n", i, Vote(nums))
	}
}
