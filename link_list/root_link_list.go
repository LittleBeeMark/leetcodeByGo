package link_list

// 使用（哨兵结点）根结点可简化链表的插入删除操作
// 使用哨兵的例子
// 查询某个值在数组中的下标
func GetIndexOfArray(list []int, key int) int {
	if list == nil {
		return -1
	}

	for i := 0; i < len(list); i++ {
		if list[i] == key {
			return i
		}
	}

	return -1

}

func GetIndexOfArray2(list []int, key int) int {
	// 缓存最后一个值
	n := len(list)
	tmp := list[n-1]

	// 将 key 付给最后一个值作为哨兵
	list[n-1] = key
	if tmp == key {
		return n - 1
	}

	i := 0
	for {
		if list[i] == key {
			break
		}
		i++
	}

	// 恢复原数组
	list[n-1] = tmp

	if i == n-1 {
		return -1
	}

	return i
}
