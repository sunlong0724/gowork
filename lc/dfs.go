package main

import (
	"fmt"
	"sort"
	"strconv"
)

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
func numIslands(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	var nums int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != 0 {
				nums++
				dfs1(grid, i, j)
			}
		}
	}
	return nums
}
func dfs1(grid [][]int, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == 0 {
		return
	}
	grid[i][j] = 0
	dirs := [][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	for _, d := range dirs {
		dfs1(grid, i+d[0], j+d[1])
	}
}

func findCircleNum(M [][]int) int {
	hasVisited := make([]bool, len(M))
	var nums int
	for i := 0; i < len(M); i++ {
		if !hasVisited[i] {
			nums++
			dfs3(M, i, hasVisited)
		}
	}
	return nums
}

func dfs3(M [][]int, i int, hasVisited []bool) {
	hasVisited[i] = true
	for j := 0; j < len(M[i]); j++ {
		if M[i][j] == 1 && !hasVisited[j] {
			dfs3(M, j, hasVisited)
		}
	}
}
func letterCombinations(digits string) [][]byte {
	digitKey := [][]byte{
		{},
		{},
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
		{'j', 'k', 'l'},
		{'m', 'n', 'o'},
		{'p', 'q', 'r', 's'},
		{'t', 'u', 'v'},
		{'w', 'x', 'y', 'z'},
	}
	if len(digits) == 0 {
		return [][]byte{}
	}
	var res [][]byte
	combinate([]byte{}, &res, []byte(digits), digitKey)
	return res
}

func combinate(curS []byte, res *[][]byte, digits []byte, digitKey [][]byte) {
	if len(curS) == len(digits) {
		d := make([]byte, len(curS))
		copy(d, curS)
		*res = append(*res, d)
		return
	}

	idx := digits[len(curS)] - '0'
	for _, c := range digitKey[idx] {
		curS = append(curS, c)
		combinate(curS, res, digits, digitKey)
		curS = curS[:len(curS)-1]
	}
}

/*
255255255128

class Solution {
public:
    vector<string> restoreIpAddresses(string s) {
        vector<string> res;
        string cur;
        back_track(0, cur, 0, s, res);
        return res;
    }

    bool check(string& s){
        if (s.empty()) return false;
        if (s[0] == '0') return s.size() == 1;
        return atoi(s.c_str()) > 0 && atoi(s.c_str()) <= 256;
    }

    void back_track(int i, string& cur, int kth, const string& s, vector<string>& res){
        if (kth >= 4){
            if ( cur.size() == s.size() + 3)
                res.push_back(cur);
            return;
        }

        for (int j = 1; j <= 3 && i < s.size() ; ++j){
            if (s.size() - (i + j) > 3*(3-kth) ) continue;
            string tmp = s.substr(i, j);
            fprintf(stdout,"%s,", tmp.c_str());
            if (check(tmp)){
                string bak = cur;
                if (!cur.empty()) cur.append(".");
                cur.append(tmp);
                back_track(i+j, cur, kth+1, s, res);
                cur = bak;
            }
        }
    }
};
*/
func restoreIpAddresses(strIp string) []string {
	var res []string
	doRestoreIpAddress(0, 0, []byte{}, strIp, &res)
	return res
}
func check(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] == '0' {
		return len(s) == 1
	}
	v, _ := strconv.Atoi(s)
	return v > 0 && v <= 255
}

func doRestoreIpAddress(idx, kth int, curS []byte, strIp string, res *[]string) {
	if kth == 4 {
		*res = append(*res, string(curS))
		return
	}
	for j := 1; j <= 3 && idx < len(strIp); j++ {
		if len(strIp)-(idx+j) > 3*(3-kth) {
			continue
		}
		tmp := []byte(strIp)[idx : idx+j]
		if check(string(tmp)) {
			bak := curS
			if len(curS) > 0 {
				curS = append(curS, '.')
			}
			curS = append(curS, tmp...)
			doRestoreIpAddress(idx+j, kth+1, curS, strIp, res)
			curS = bak
		}
	}
}

/*
For example,
Given board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]
word = "ABCCED", -> returns true,
word = "SEE", -> returns true,
word = "ABCB", -> returns false.
*/
func exist(board [][]byte, word string) bool {
	hasVisited := make([][]bool, len(board))
	for i := range hasVisited {
		hasVisited[i] = make([]bool, len(board[i]))
	}
	for i, s := range board {
		for j, _ := range s {
			if checkWord(i, j, 0, board, hasVisited, word) {
				return true
			}
		}
	}
	return false
}

func checkWord(i, j, k int, board [][]byte, hasVisited [][]bool, word string) bool {
	if k == len(word) {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) || hasVisited[i][j] || board[i][j] != word[k] {
		return false
	}

	dirs := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	hasVisited[i][j] = true
	for _, d := range dirs {
		b := checkWord(i+d[0], j+d[1], k+1, board, hasVisited, word)
		if b {
			return true
		}
	}
	hasVisited[i][j] = false
	return false

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BinaryTreePath(root *TreeNode) [][]byte {
	var res [][]byte
	backTraceTree(root, []byte{}, &res)
	return res
}

func backTraceTree(root *TreeNode, curPath []byte, res *[][]byte) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		curPath = append(curPath, []byte(strconv.Itoa(root.Val))...)
		d := make([]byte, len(curPath))
		copy(d, curPath)
		*res = append(*res, d)
		return
	}
	curPath = append(curPath, []byte( strconv.Itoa(root.Val))...)
	curPath = append(curPath, []byte("->")...)
	backTraceTree(root.Left, curPath, res)
	backTraceTree(root.Right, curPath, res)
}

func permute(nums []int) [][]int {
	var res [][]int
	hasVisited := make([]bool, len(nums))
	doPermute([]int{}, nums, hasVisited, &res)
	return res
}

func doPermute(curArr, nums []int, hasVisited []bool, res *[][]int) {
	if len(curArr) == len(nums) {
		d := make([]int, len(curArr))
		copy(d, curArr)
		*res = append(*res, d)
		return
	}

	for j := 0; j < len(nums); j++ {
		if hasVisited[j] {
			continue
		}
		hasVisited[j] = true
		curArr = append(curArr, nums[j])
		doPermute(curArr, nums, hasVisited, res)
		curArr = curArr[:len(curArr)-1]
		hasVisited[j] = false
	}
}

/*
[1,1,2] have the following unique permutations:
[[1,1,2], [1,2,1], [2,1,1]]
数组元素可能含有相同的元素，进行排列时就有可能出现重复的排列，要求重复的排列只返回一个。

在实现上，和 Permutations 不同的是要先排序，然后在添加一个元素时，判断这个元素是否等于前一个元素，如果等于，并且前一个元素还未访问，那么就跳过这个元素。
*/
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	hasVisited := make([]bool, len(nums))
	var res [][]int
	doPermuteUnique([]int{}, hasVisited, &res, nums)
	return res
}
func doPermuteUnique(curS []int, hasVisited []bool, res *[][]int, nums []int) {
	//fmt.Println(curS)
	if len(curS) == len(nums) {
		d := make([]int, len(curS))
		copy(d, curS)
		*res = append(*res, d)
		return
	}

	for i := 0; i < len(nums); i++ {
		if hasVisited[i] || i > 0 && nums[i-1] == nums[i] && !hasVisited[i-1] {
			continue
		}
		hasVisited[i] = true
		curS = append(curS, nums[i])
		doPermuteUnique(curS, hasVisited, res, nums)
		curS = curS[:len(curS)-1]
		hasVisited[i] = false
	}
}

/*
If n = 4 and k = 2, a solution is:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
public List<List<Integer>> combine(int n, int k) {
*/
func combine(n, k int) [][]int {
	var res [][]int
	doCombine([]int{}, 1, n, k, &res)
	return res
}

func doCombine(curS []int, i, n, k int, res *[][]int) {
	if len(curS) == k {
		d := make([]int, len(curS))
		copy(d, curS)
		*res = append(*res, d)
		return
	}
	for j := i; j <= n; j++ {
		curS = append(curS, j)
		doCombine(curS, j+1, n, k, res)
		curS = curS[:len(curS)-1]
	}
}

/*
given candidate set [2, 3, 6, 7] and target 7,
A solution set is:
[[7],[2, 2, 3]]
public List<List<Integer>> combinationSum(int[] candidates, int target)

*/
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	sort.Ints(candidates)
	for i := 0; i < len(candidates); i++{
		doCombinationSum([]int{}, candidates, i,target, &res)
	}
	return res
}

func doCombinationSum(curS, candidates []int, i, target int, res *[][]int) {
	if target == 0 {
		d := make([]int, len(curS))
		copy(d, curS)
		*res = append(*res, d)
		return
	}
	for j := i; j < len(candidates); j++{
		if target-candidates[j] >= 0 {
			curS = append(curS, candidates[j])
			doCombinationSum(curS, candidates, j, target-candidates[i], res)
			curS = curS[:len(curS)-1]
		}
	}

}
/*
For example, given candidate set [10, 1, 2, 7, 6, 1, 5] and target 8,
A solution set is:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
 */
func combinationSum2(candidates []int, target int) [][]int{
	var res [][]int
	hasVisited := make([]bool, len(candidates))
	sort.Ints(candidates)

	doCombinationSum2([]int{}, 0, candidates, hasVisited, target, &res)

	return res

}

func doCombinationSum2(curS []int, i int, candidates []int, hasVisited []bool, target int, res *[][]int) {
	if target == 0 {
		d := make([]int, len(curS))
		copy(d, curS)
		*res = append(*res, d)
		return
	}

	for j := i; j < len(candidates); j++{
		if hasVisited[j] || j > 0 && candidates[j-1] == candidates[j] && !hasVisited[j-1]{
			continue
		}

		hasVisited[j] = true
		curS = append(curS, candidates[j])
		doCombinationSum2(curS, j + 1, candidates, hasVisited, target-candidates[j], res)
		hasVisited[j] = false
		curS = curS[:len(curS)-1]
	}

}

/*
Input: k = 3, n = 9

Output:

[[1,2,6], [1,3,5], [2,3,4]]
从 1-9 数字中选出 k 个数不重复的数，使得它们的和为 n。

public List<List<Integer>> combinationSum3(int k, int n) {
 */

func combinationSum3(k,n int) [][]int{
	var res [][]int
	doCombinationSum3([]int{}, 1, k, n, &res)
	return  res
}

func doCombinationSum3(curS []int, i, k,n int, res *[][]int){
	//fmt.Println(curS, n)
	if n == 0 && len(curS) == k{
		d := make([]int, len(curS))
		copy(d, curS)
		*res = append(*res, d)
		return
	}

	for j := i; j <= n; j++{
		if n - j >= 0{
			curS = append(curS, j)
			doCombinationSum3(curS, j+1, k, n - j, res)
			curS = curS[:len(curS)-1]
		}
	}
}
/*
找出集合的所有子集，子集不能重复，[1, 2] 和 [2, 1] 这种子集算重复

public List<List<Integer>> subsets(int[] nums) {
 */
func subsets(nums []int) [][]int{
	var res [][]int
	sort.Ints(nums)
	doSubsets([]int{}, 0, nums, &res)
	return res
}

func doSubsets(curS []int, i int, nums []int, res *[][]int){
	if len(curS) > 0{
		d := make([]int, len(curS))
		copy(d, curS)
		*res = append(*res, d)
		//return
	}

	for j := i; j < len(nums); j++{
		curS = append(curS, nums[j])
		doSubsets(curS, j+1, nums, res)
		curS = curS[:len(curS)-1]
	}

}

/*
For example,
If nums = [1,2,2], a solution is:

[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]
public List<List<Integer>> subsetsWithDup(int[] nums) {
 */
func subsetsWithDup(nums []int) [][]int{
	var res [][]int
	hasVisited := make([]bool, len(nums))
	doSubsetsWithDup([]int{}, 0, hasVisited, nums, &res)
	res = append(res, []int{})
	return res
}

func doSubsetsWithDup(curS []int, i int, hasVisited []bool, nums []int, res *[][]int){
	if len(curS) > 0{
		d := make([]int, len(curS))
		copy(d, curS)
		*res = append(*res, d)
	}

	for j := i; j < len(nums);j++{
		if hasVisited[j] || j > 0 && nums[j-1] == nums[j] && !hasVisited[j-1] {
			continue
		}
		hasVisited[j] = true
		curS = append(curS, nums[j])
		doSubsetsWithDup(curS, j+1, hasVisited, nums, res)
		curS = curS[:len(curS)-1]
		hasVisited[j] = false
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
	/*
		M := [][]int{
		{1,1,0},
		{1,1,0},
		{0,0,1},
		}
		fmt.Println(findCircleNum(M))

	*/
	/*
		res := letterCombinations("23")
		for _, s := range res{
			fmt.Println()
			for _, c := range  s{
				fmt.Printf("%c", c)
			}
		}
	*/
	//fmt.Println( restoreIpAddresses("250125128"))
	/*
		board := [][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}
		fmt.Println( exist(board, "ABCCED") )
	*/
	/*
		 s:=[]TreeNode {
			{1, nil,nil},
			{ 2, nil, nil},
			{3, nil, nil},
			 {5, nil, nil},
		 }
		 s[0].Left, s[0].Right = &s[1], &s[2]
		 s[1].Left = &s[3]

		res := BinaryTreePath(&s[0])
		for _, s := range res{
			fmt.Println()
			fmt.Printf("%s", string(s))
		}
		fmt.Println()
	*/
	/*
		var s []byte
		fmt.Printf("%V %T\n", s, s)
		fmt.Printf("%v %T\n", s, s)
	*/
	/*
		res := permute([]int{1,2,3,4,5})
	*/
	/*
		res := permuteUnique([]int{1,2,4,3,2})
	*/
	/*
		res := combine(9,2)
	*/
	/*
	res := combinationSum([]int{2, 3, 6, 7}, 7)
	}
	 */
	/*
	res := combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	 */
	/*
	res := combinationSum3(3, 9)
	 */
	/*
	res := subsets([]int{1,2,3,4,5})
	 */
	res := subsetsWithDup([]int{1,2,2})
	for _,s := range res{
		fmt.Println(s)
	}
}
