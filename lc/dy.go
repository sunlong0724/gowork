package main

import "fmt"

/*
题目描述：有 N 阶楼梯，每次可以上一阶或者两阶，求有多少种上楼梯的方法。
定义一个数组 dp 存储上楼梯的方法数（为了方便讨论，数组下标从 1 开始），dp[i] 表示走到第 i 个楼梯的方法数目。
第 i 个楼梯可以从第 i-1 和 i-2 个楼梯再走一步到达，走到第 i 个楼梯的方法数为走到第 i-1 和第 i-2 个楼梯的方法数之和。
考虑到 dp[i] 只与 dp[i - 1] 和 dp[i - 2] 有关，因此可以只用两个变量来存储 dp[i - 1] 和 dp[i - 2]，使得原来的 O(N) 空间复杂度优化为 O(1) 复杂度。
public int climbStairs(int n) {
*/
func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}
func climbStairsDy(n int) int {
	if n <= 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
func climbStairsDy2(n int) int {
	if n <= 2 {
		return n
	}
	a, b := 1, 2
	var c int
	for i := 3; i <= n; i++ {
		c = a + b
		a, b = b, c
	}
	return c
}

/*
题目描述：抢劫一排住户，但是不能抢邻近的住户，求最大抢劫量。
定义 dp 数组用来存储最大的抢劫量，其中 dp[i] 表示抢到第 i 个住户时的最大抢劫量。
由于不能抢劫邻近住户，如果抢劫了第 i -1 个住户，那么就不能再抢劫第 i 个住户，所以
//定义dp数组
//初始化
//定义状态转移方程
*/
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	if nums[1] > nums[0] {
		dp[1] = nums[1]
	} else {
		dp[1] = nums[0]
	}
	for i := 2; i < len(nums); i++ {
		if dp[i-2]+nums[i] > dp[i-1] {
			dp[i] = dp[i-2] + nums[i]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(dp)-1]
}

func robDy2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	var (
		dp0 = 0
		dp1 = 0
		dp  = 0
	)
	if nums[1] > nums[0] {
		dp1 = nums[1]
	}
	for i := 2; i < len(nums); i++ {
		if dp0+nums[i] > dp1 {
			dp = dp0 + nums[i]
		} else {
			dp = dp1
		}
		dp0, dp1 = dp1, dp
	}
	return dp
}

/*
213. House Robber II (Medium)
public int rob(int[] nums) {
    if (nums == null || nums.length == 0) {
        return 0;
    }
    int n = nums.length;
    if (n == 1) {
        return nums[0];
    }
    return Math.max(rob(nums, 0, n - 2), rob(nums, 1, n - 1));
}

private int rob(int[] nums, int first, int last) {
    int pre2 = 0, pre1 = 0;
    for (int i = first; i <= last; i++) {
        int cur = Math.max(pre1, pre2 + nums[i]);
        pre2 = pre1;
        pre1 = cur;
    }
    return pre1;
}
*/
func robDy3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	le := len(nums)
	if le == 1 {
		return nums[0]
	}

	v1 := doRobDy3(nums, 0, le-2)
	v2 := doRobDy3(nums, 1, le-1)
	if v1 > v2 {
		return v1
	} else {
		return v2
	}
}
func doRobDy3(nums []int, si, ei int) int {
	pre1, pre2, cur := 0, 0, 0

	for i := si; i <= ei; i++ {
		if pre1 > pre2+nums[i] {
			cur = pre1
		} else {
			cur = pre2 + nums[i]
		}
		pre1, pre2 = cur, pre1
	}
	return cur

}

/*
[[1,3,1],
 [1,5,1],
 [4,2,1]]
Given the above grid map, return 7. Because the path 1→3→1→1→1 minimizes the sum.
题目描述：求从矩阵的左上角到右下角的最小路径和，每次只能向右和向下移动。
确实不能bfs，bfs的话就是贪心了。
动规
1.定义dp数组 dp[i][j]
2.初始化   dp[0][0] = 1
3.找出状态转移方程  dp[i][j] = min(dp[i-1][j],dp[i][j-1])
4.返回值 dp[i][j]
5.再优化
*/
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	if len(grid) == 1 {
		return grid[0][0]
	}
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[i]))
	}
	dp[0][0] = grid[0][0]
	for j := 1; j < len(grid[0]); j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < len(grid); i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[i]); j++ {
			v1 := dp[i-1][j] + grid[i][j]
			v2 := dp[i][j-1] + grid[i][j]
			if v1 < v2 {
				dp[i][j] = v1
			} else {
				dp[i][j] = v2
			}
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}
func minPathSum2(grid [][]int)int {
	if len(grid) == 0 {
		return 0
	}
	if len(grid) == 1{
		return grid[0][0]
	}
	/*
	1.定义dp  dp[i]
	2.初始化  dp[j] = grid[0][j]
	3.寻找状态转移方程  dp[i] = min(dp[i] + grid[i][j], dp[i-1] + grid[i][j])
	4.返回值 dp[i]
	 */
	dp := make([]int, len(grid[0]))
	dp[0] = grid[0][0]
	for j := 1; j < len(grid[0]); j++{
		dp[j] = grid[0][j] + dp[j-1]
	}
	for i := 1; i < len(grid); i++{
		for j := 0; j < len(grid[i]); j++{
			if j == 0 || dp[j] < dp[j-1]{
				dp[j] = dp[j] + grid[i][j]
			}else{
				dp[j] = dp[j-1] + grid[i][j]
			}
		}
	}
	return dp[len(dp)-1]
}
/*
统计从矩阵左上角到右下角的路径总数，每次只能向右或者向下移动。
 */
func uniquePaths(m,n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if m == 1 || n == 1 {
		return 1
	}

	dp := make([]int, n+1)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++{
			if i == 1 {
				dp[j] = 1
			}else{
				dp[j] = dp[j] + dp[j-1]
			}
		}
	}
	return dp[n]
}

/*
303. Range Sum Query - Immutable (Easy)
Given nums = [-2, 0, 3, -5, 2, -1]
sumRange(0, 2) -> 1
sumRange(2, 5) -> -1
sumRange(0, 5) -> -3
求区间 i ~ j 的和，可以转换为 sum[j + 1] - sum[i]，其中 sum[i] 为 0 ~ i - 1 的和。
 */

/*
数组中等差递增子区间的个数
413. Arithmetic Slices (Medium)
A = [0, 1, 2, 3, 4]
return: 6, for 3 arithmetic slices in A:
[0, 1, 2],
[1, 2, 3],
[0, 1, 2, 3],
[0, 1, 2, 3, 4],
[ 1, 2, 3, 4],
[2, 3, 4]
dp[i] 表示以 A[i] 为结尾的等差递增子区间的个数。

当 A[i] - A[i-1] == A[i-1] - A[i-2]，那么 [A[i-2], A[i-1], A[i]] 构成一个等差递增子区间。而且在以 A[i-1] 为结尾的递增子区间的后面再加上一个 A[i]，一样可以构成新的递增子区间。

dp[2] = 1
    [0, 1, 2]
dp[3] = dp[2] + 1 = 2
    [0, 1, 2, 3], // [0, 1, 2] 之后加一个 3
    [1, 2, 3]     // 新的递增子区间
dp[4] = dp[3] + 1 = 3
    [0, 1, 2, 3, 4], // [0, 1, 2, 3] 之后加一个 4
    [1, 2, 3, 4],    // [1, 2, 3] 之后加一个 4
    [2, 3, 4]        // 新的递增子区间
综上，在 A[i] - A[i-1] == A[i-1] - A[i-2] 时，dp[i] = dp[i-1] + 1。

因为递增子区间不一定以最后一个元素为结尾，可以是任意一个元素结尾，因此需要返回 dp 数组累加的结果。

public int numberOfArithmeticSlices(int[] A) {
 */

func numberOfArithmeticSlices(A []int) int{
	/*
	1. 定义dp  dp[i] 以A[i]结尾的等差递增子序列个数
	2. 初始化  dp[0] dp[1] = 0,0
	3.状态转移方程    当A[i]-A[i-1] == A[i-1]-A[i-2]时  dp[i] = dp[i-1] + 1
	4. 返回dp[i]的数组和
	 */
	if len(A) < 3 {
		return 0
	}
	dp := make([]int, len(A))
	dp[2] = 1
	for i := 3; i < len(A); i++{
		if A[i] - A[i-1] == A[i-1]-A[i-2]{
			dp[i] = dp[i-1]+1
		}
	}
	ret := 0
	for _, v := range dp{
		ret += v
	}
	return ret
}



/*
1. 分割整数的最大乘积
343. Integer Break (Medim)
题目描述：For example, given n = 2, return 1 (2 = 1 + 1); given n = 10, return 36 (10 = 3 + 3 + 4).
public int integerBreak(int n) {
 */
/*
分割整数构成字母字符串
91. Decode Ways (Medium)
题目描述：Given encoded message "12", it could be decoded as "AB" (1 2) or "L" (12).

public int numDecodings(String s) {
 */

/*
最长递增子序列
已知一个序列 {S1, S2,...,Sn}，取出若干数组成新的序列 {Si1, Si2,..., Sim}，其中 i1、i2 ... im 保持递增，即新序列中各个数仍然保持原数列中的先后顺序，称新序列为原序列的一个 子序列 。

如果在子序列中，当下标 ix > iy 时，Six > Siy，称子序列为原序列的一个 递增子序列 。

定义一个数组 dp 存储最长递增子序列的长度，dp[n] 表示以 Sn 结尾的序列的最长递增子序列长度。对于一个递增子序列 {Si1, Si2,...,Sim}，如果 im < n 并且 Sim < Sn，此时 {Si1, Si2,..., Sim, Sn} 为一个递增子序列，递增子序列的长度增加 1。满足上述条件的递增子序列中，长度最长的那个递增子序列就是要找的，在长度最长的递增子序列上加上 Sn 就构成了以 Sn 为结尾的最长递增子序列。因此 dp[n] = max{ dp[i]+1 | Si < Sn && i < n} 。

因为在求 dp[n] 时可能无法找到一个满足条件的递增子序列，此时 {Sn} 就构成了递增子序列，需要对前面的求解方程做修改，令 dp[n] 最小为 1，即：



对于一个长度为 N 的序列，最长递增子序列并不一定会以 SN 为结尾，因此 dp[N] 不是序列的最长递增子序列的长度，需要遍历 dp 数组找出最大值才是所要的结果，max{ dp[i] | 1 <= i <= N} 即为所求。

1. 最长递增子序列
300. Longest Increasing Subsequence (Medium)

Leetcode / 力扣

public int lengthOfLIS(int[] nums) {
 */
func main() {
	//fmt.Println(climbStairs(5))
	//fmt.Println(climbStairsDy(5))
	//fmt.Println(climbStairsDy2(5))
	//fmt.Println( rob([]int{1,2,3,4,5,6,5,8}) )
	//fmt.Println( robDy2([]int{1,2,3,4,5,6,5,8}) )
	//fmt.Println( robDy3([]int{1,2,3,4,5,6,5}) )
	/*
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(minPathSum(grid))
	fmt.Println(minPathSum2(grid))
	 */
	//fmt.Println(uniquePaths(3,3))
	fmt.Println(numberOfArithmeticSlices([]int{0,1,2,3,4}))
}
