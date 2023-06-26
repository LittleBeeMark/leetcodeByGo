package link_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// 141
//给你一个链表的头节点 head ，判断链表中是否有环。
//
//如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。注意：pos 不作为参数进行传递 。仅仅是为了标识链表的实际情况。
//
//如果链表中存在环 ，则返回 true 。 否则，返回 false 。
// 思路: 快指针和慢指针同时向前，快指针走一步慢指针走两步，快指针和慢指针的最终结局只有三个。
// 其中如果链表无环，那么快指针会先到尾部，快指针到尾部有两种可能 1.直接到尾部，2.到了向前一步就是尾部的地方 这两个结局返回false。
// 如果链表有环那么只有快指针等于慢指针的结局，这个结局返回true
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast, low := head.Next, head
	for fast != nil && fast.Next != nil && fast != low {
		fast = fast.Next.Next
		low = low.Next
	}

	return fast == low
}

// 142
//给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
//
//如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
//
//不允许修改 链表。
//
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast, low := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		low = low.Next
		if fast == low {
			low = head
			for low != fast {
				low = low.Next
				fast = fast.Next
			}
			return low
		}
	}

	return nil
}

// 202
// 编写一个算法来判断一个数 n 是不是快乐数。
//
//「快乐数」 定义为：
//
// 对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
// 然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
// 如果这个过程 结果为 1，那么这个数就是快乐数。
// 如果 n 是 快乐数 就返回 true ；不是，则返回 false 。
//
func getNextNum(num int) int {
	var reNum int
	for num != 0 {
		reNum += (num % 10) * (num % 10)
		num = num / 10
	}

	return reNum
}

func isHappy(n int) bool {
	if n == 1 {
		return true
	}

	fast, low := getNextNum(n), n
	for fast != 1 && fast != low {
		fast = getNextNum(getNextNum(fast))
		low = getNextNum(low)
	}

	return fast == 1
}

// 206. 反转链表
//双指针
// pre 指向反转头的头节点
// cur 指向未反转的头节点
// next 缓存未反转头节点的下一个节点用来将cur移向下一个
// 思路：做一个新的头节点,将老的头节点指向新的，然后继续遍历依次反转直到nil为止，
// 因为每次要把当前的节点先指向新链表，所以每次都要缓存一下next用来将当前节点移动向下一个来实现遍历
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre, next *ListNode
	cur := head
	for cur != nil {
		// 记录未反转头的next
		next = cur.Next
		// 反转当前节点指针
		cur.Next = pre
		// 反转头移位
		pre = cur
		// 未反转头移位
		cur = next
	}

	return pre
}

// 递归
// 先递，直到得到链表尾部的nil节点的前节点 再归，再开始反转处理
// 下面的链接的动画非常易懂
//https://leetcode.cn/problems/reverse-linked-list/solution/fan-zhuan-lian-biao-shuang-zhi-zhen-di-gui-yao-mo-/
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 第一次归，得到新头
	newHead := reverseList2(head.Next)
	// 第一次到这head是尾部节点的前节点，将尾部节点的Next进行反转指向尾部节点的前节点
	head.Next.Next = head
	// 将尾部前节点的指向至空
	head.Next = nil
	// 此时就完成了归的第一个子问题，得到了第一个反转子链
	return newHead
}

// 反转链表头n个节点
func reverseListN(head *ListNode, n int) *ListNode {
	if n == 1 {
		return head
	}

	// 第一次归，得到新头
	newHead := reverseListN(head.Next, n-1)
	// 第一次到这head是尾部节点的前节点，将尾部节点的Next进行反转指向尾部节点的前节点
	head.Next.Next = head
	// 将尾部前节点的指向至空
	head.Next = nil
	// 此时就完成了归的第一个子问题，得到了第一个反转子链
	return newHead
}

//92. 反转链表 II
//给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 找虚拟头
	// 为什么要做一个虚拟头，如果需要反转的left为开头第一个元素即为1此时下面部分的反转是不成立的，
	// 我们需要单独写一个逻辑，为了逻辑简单我们必须引入链表头
	vitural := &ListNode{
		Next: head,
	}

	// 走到需要反转节点的前节点（找到前节点）
	pre := vitural
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	var next *ListNode
	cur := pre.Next
	// 前置节点及当前节点（只用来定位下一个next）不变化，而后置节点不断地翻到前置节点后
	for i := 0; i < right-left; i++ {
		next = cur.Next
		cur.Next = next.Next
		// next 指向pre.Next
		next.Next = pre.Next
		pre.Next = next
	}

	return vitural.Next
}

// 判断链表有环与否
func hasCyclet(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast := head.Next
	low := head
	for fast != nil || fast.Next != nil || fast != low {
		fast = fast.Next.Next
		low = low.Next
	}

	return low == fast
}

func detectCyclet(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, low := head, head
	for fast != nil || fast.Next != nil {
		fast = fast.Next.Next
		low = low.Next
		if low == fast {
			low = head
			for low != head {
				low = low.Next
				fast = fast.Next
			}

			return low
		}
	}

	return nil

}

func getNextNumt(num int) int {
	var next int
	for num != 0 {
		next = (num % 10) * (num % 10)
		num = num / 10
	}

	return next
}

func isHappyt(n int) bool {
	if n == 1 {
		return true
	}
	fast, low := getNextNumt(n), n
	for fast != 1 && fast != low {
		fast = getNextNumt(getNextNumt(fast))
		low = getNextNumt(low)
	}

	return fast == 1
}

func reverseListt(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre, next *ListNode
	cur := head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

func reverseList2t(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList2t(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func ReverseBetween2(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return head
	}

	pre := head
	for i := 0; i < left-1; i++ {
		//走到left前一位
		pre = pre.Next
	}

	var next *ListNode
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next = cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next

	}

	return head
}
