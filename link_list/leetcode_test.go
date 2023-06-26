package link_list

import (
	"fmt"
	"testing"
	"time"
)

func TestFORRange(t *testing.T) {
	tasks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < len(tasks)-1; i++ {
		val := tasks[i]
		index := i
		go func() {
			fmt.Printf("第 %d 个任务：%d \n", index, val)
		}()
	}

	time.Sleep(12 * time.Second)

}
