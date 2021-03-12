package main

import "fmt"

/*
bfs 解决最短路径问题
dfs 解决可达性问题
*/
/*
查找最大连通面积
*/
func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	var maxArea int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			ma := dfs0(grid, i, j)
			if ma > maxArea {
				maxArea = ma
			}
		}
	}
	return maxArea
}
func dfs0(grid [][]int, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == 0 {
		return 0
	}
	var (
		dirs = [][2]int{
			{-1, 0},
			{1, 0},
			{0, -1},
			{0, 1},
		}
	)
	grid[i][j] = 0
	area := 1
	for _, d := range dirs {
		area += dfs0(grid, i+d[0], j+d[1])
	}
	return area
}
/*
11000
11000
00100
00011
 */
func numIslands(grid [][]int) int{
	if len(grid) == 0 {
		return 0
	}
	var nums int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++{
			if grid[i][j] != 0 {
				nums++
				dfs1(grid, i, j)
			}
		}
	}
	return nums
}
func dfs1(grid [][]int, i, j int){
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == 0{
		return
	}
	grid[i][j] = 0
	dirs := [][2]int{
		{-1,0},
		{1,0},
		{0,-1},
		{0,1},
	}
	for _,d := range dirs{
		dfs1(grid, i + d[0], j + d[1])
	}
}


func findCircleNum(M [][]int) int{
	hasVisited := make([]bool,len(M))
	var nums int
	for i := 0; i < len(M); i++{
		if !hasVisited[i]{
			nums++
			dfs3(M, i, hasVisited)
		}
	}
	return nums
}

func dfs3(M [][]int, i int, hasVisited []bool){
	hasVisited[i] = true
	for j := 0; j < len(M[i]); j++{
		if M[i][j] == 1 && !hasVisited[j] {
			dfs3(M, j, hasVisited)
		}
	}
}


func main() {
	/*
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	fmt.Println(maxAreaOfIsland(grid))
	 */
/*
	grid := [][]int{
		{1,1,0,0,0},
		{1,1,0,0,0},
		{0,0,1,0,0},
		{0,0,0,1,1},
	}
	fmt.Println(numIslands(grid))
 */
	M := [][]int{
	{1,1,0},
	{1,1,0},
	{0,0,1},
	}
	fmt.Println(findCircleNum(M))
}
