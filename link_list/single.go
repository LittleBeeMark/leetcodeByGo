package link_list

// 单链表代码
type Object interface{}

type Node struct {
	Data Object
	next *Node
}

type ListSingle struct {
	size uint64
	head *Node
	tail *Node
}

func (list *ListSingle) Init() {
	(*list).size = 0
	(*list).head = nil
	(*list).tail = nil
}

// 向链表追加节点
func (list *ListSingle) Append(node *Node) bool {
	if node == nil {
		return false
	}

	(*node).next = nil // 新加节点在末尾，没有next
	if (*list).size == 0 {
		(*list).head = node
	} else {
		oldTail := (*list).tail // 取尾结点
		(*oldTail).next = node  // 尾结点的next指向新加节点
	}

	(*list).tail = node // 新节点是尾结点
	(*list).size++
	return true
}

// 向第i个节点处插入节点
func (list *ListSingle) Insert(i uint64, node *Node) bool {
	if node == nil || i > (*list).size || (*list).size == 0 {
		return false
	}

	if i == 0 {
		(*node).next = (*list).head
		(*list).head = node
	} else {
		preNode := (*list).head
		for j := uint64(1); j < i; j++ {
			preNode = (*preNode).next
		}

		(*node).next = (*preNode).next // 新节点指向旧节点原来所指的next
		(*preNode).next = node         // 原节点的next指向新节点
	}
	(*list).size++

	return true
}

// 移除指定位置的节点
func (list *ListSingle) Remove(i uint64) bool {
	if i >= (*list).size {
		return false
	}

	if i == 0 {
		preHead := (*list).head     // 取出旧的链表头
		(*list).head = preHead.next // 旧链表头的next变为新的头

		// 如果仅有一个节点，则头尾节点清空
		if (*list).size == 1 {
			(*list).head = nil
			(*list).tail = nil
		}
	} else {
		preNode := (*list).head
		for j := uint64(1); j < i; j++ {
			preNode = (*preNode).next
		}

		node := (*preNode).next     // 找到当前要删除的节点
		(*preNode).next = node.next // 把当前要删除节点的next赋给其父节点的next,完成后代转移

		// 若删除的尾部，尾部指针需要调整
		if i == ((*list).size - 1) {
			(*list).tail = preNode
		}
	}

	(*list).size--

	return true
}

// 移除所有节点
func (list *ListSingle) RemoveAll() bool {
	(*list).Init()
	return true
}

// 获取指定位置的节点
func (list *ListSingle) Get(i uint64) *Node {
	if i >= (*list).size {
		return nil
	}

	node := (*list).head
	for j := uint64(0); j < i; j++ {
		node = (*node).next
	}

	return node
}

// 搜索某个数据的节点位置
func (list *ListSingle) IndexOf(data Object) int64 {
	pos := int64(-1)
	node := (*list).head
	if node.Data == data {
		return 0
	}

	for j := uint64(1); j < (*list).size; j++ {
		if node != nil {
			node = (*node).next
			if node != nil && node.Data == data {
				pos = int64(j)
				break
			}
		}
	}
	return pos
}

// 取得链表长度
func (list *ListSingle) GetSize() uint64 {
	return (*list).size
}

// 取得链表头
func (list *ListSingle) GetHead() *Node {
	return (*list).head
}

// 取得链表尾
func (list *ListSingle) GetTail() *Node {
	return (*list).tail
}
