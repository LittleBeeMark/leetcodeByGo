package main

import "fmt"

// 北，东，
// 模拟机器人行走
// https://leetcode-cn.com/problems/walking-robot-simulation/
func robotSim(commands []int, obstacles [][]int) int {
	// 定义由方向对应的x , y 坐标变化
	// 北 +y 东 +x  南 -y 西 -x
	directX := []int{0, 1, 0, -1}
	directY := []int{1, 0, -1, 0}
	// 定义当前方向
	di := 0
	// 定义当前位置
	currentX, currentY, result := 0, 0, 0

	obstacles_m := make(map[[2]int]bool)
	for _, v := range obstacles {
		if len(v) == 2 {
			obstacles_m[[2]int{v[0], v[1]}] = true
		}
	}

	for i := 0; i < len(commands); i++ {
		cd := commands[i]
		if cd == -1 {
			di = (di + 1) % 4
		} else if cd == -2 {
			di = (di + 3) % 4
		} else {
			for j := 1; j <= cd; j++ {
				currentX += directX[di]
				currentY += directY[di]
				if _, ok := obstacles_m[[2]int{currentX, currentY}]; ok {
					currentX -= directX[di]
					currentY -= directY[di]
					break
				}
				result = gMax(result, currentY*currentY+currentX*currentX)
			}
		}
	}

	return result
}

func gMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(robotSim([]int{4, -1, 4, -2, 4}, [][]int{{2, 4}}))
}
