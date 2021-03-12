package main

import "fmt"

/*
[[1,1,0,1],
 [1,0,1,0],
 [1,1,1,1],
 [1,0,1,1]]
0 表示可以经过某个位置，求解从左上角到右下角的最短路径长度。
*/

func findShortestPath(nums [][]int) int {
	if nums == nil || len(nums) == 0 || len(nums[0]) == 0 {
		return -1
	}
	var (
		r    = len(nums)
		c    = len(nums[0])
		dirs = [][2]int{
			{-1, -1},
			{-1, 0},
			{-1, 1},
			{0, -1},
			{0, 1},
			{1, -1},
			{1, 0},
			{1, 1},
		}
		que     = [][2]int{}
		pathLen = 0
	)
	que = append(que, [2]int{0, 0})

	for len(que) > 0 {
		pathLen++
		for size := len(que); size > 0; size-- {
			e := que[0]
			que = que[1:]

			if e[0] == r-1 && e[1] == c-1 {
				return pathLen
			}

			nums[e[0]][e[1]] = 1
			for i := 0; i < len(dirs); i++ {
				t := [2]int{e[0] + dirs[i][0], e[1] + dirs[i][1]}
				if t[0] < 0 || t[0] >= r || t[1] < 0 || t[1] >= c || nums[t[0]][t[1]] == 1 {
					continue
				}
				que = append(que, t)
			}
		}
	}
	return -1
}

func numSquares(n int) int {
	var (
		squares = generateSquare(n)
		que     = []int{}
		marked  = make([]bool, n+1)
		level   = 0
	)
	que = append(que, n)
	marked[n] = true

	for len(que) > 0 {
		level++
		for size := len(que); size > 0; size-- {
			cur := que[0]
			que = que[1:]
			for _, v := range squares {
				next := cur - v
				if next < 0 {
					break
				}
				if next == 0 {
					return level
				}
				if marked[next] {
					continue
				}
				marked[next] = true
				que = append(que, next)
			}
		}
	}

	return n
}

/*
generate squares squence below n
1,4,9,16,25,36,49,64
3,5,7,9,11,13,15
*/
func generateSquare(n int) []int {
	var res []int
	square, diff := 1, 3
	for square <= n {
		//fmt.Println(square)
		res = append(res, square)
		square += diff
		diff += 2
	}
	return res
}

func ladderLength(beginWord,endWord string, wordList []string) int{
	wordList = append(wordList, beginWord)
	N := len(wordList)
	start, end := N-1,0
	for end < N && wordList[end] != endWord {
		end++
	}
	if end == N {
		return 0
	}

	graphic := buildGraph(wordList)
	return getShortestPath(graphic, start, end)
}

func isConnected(s1, s2 string) bool{
	diff := 0
	for i := 0; i < len(s1) && i < len(s2); i++{
		if s1[i] != s2[i]{
			diff++
		}
	}
	return diff == 1
}
func buildGraph(wordList []string)[][]int {
	N := len(wordList)
	res := [][]int{}
	for i := 0; i < N; i++{
		res = append(res, []int{})
		for j := 0; j < N; j++{
			if isConnected(wordList[i], wordList[j]){
				res[i] = append( res[i], j)
			}
		}
	}
	return res
}

func getShortestPath(graphic [][]int, start, end int) int{
	var (
		que = []int{}
		marked = make([]bool, len(graphic))
		path = 1
	)
	que = append(que, start)
	for len(que) > 0 {
		path++
		for size := len(que); size > 0; size--{
			cur := que[0]
			que = que[1:]
			for _, v := range graphic[cur]	{
				if v == end{
					return path
				}
				if marked[v] {
					continue
				}
				marked[v] = true
				que = append(que, v)
			}
		}
	}
	return 0
}

func main() {
	/*
		var nums = [][]int{
			{0, 1, 0, 1},
			{1, 0, 1, 0},
			{1, 0, 0, 1},
			{1, 0, 0, 0},
		}
		fmt.Println(findShortestPath(nums))
	*/
	//fmt.Println(numSquares(13))
	fmt.Println(ladderLength("hit", "dog", []string{"hot","dot","dog","lot","log"}))
}
