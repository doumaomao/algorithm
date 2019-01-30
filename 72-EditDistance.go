// 给定两个单词word1和word2，找到将word1转换为word2所需的最小操作数

// 对单词允许3个操作：插入、删除、替换

/*
使用dp[i][j]用来表示字符串1的0~i-1、字符串2的0~j-1的最小编辑距离
边界值如下：dp[i][0] = i、dp[0][j] = j
同时对于两个字符串的子串，都能分为最后一个字符相等或者不等的情况:
1.如果words[i-1] == words[j-1]：dp[i][j] = dp[i-1][j-1]；也就是说当前的编辑距离和位置i和j的字符无关；
2.如果words[i-1] != words[j-1]：则存在三种可能的操作：
2.1向word1插入：dp[i][j] = dp[i][j-1] + 1;
2.2从word1删除：dp[i][j] = dp[i-1][j] + 1;
2.3替换word1元素：dp[i][j] = dp[i-1][j-1] + 1;
*/

package algorithm

func minDistance(word1 string, word2 string) int {
	rows := len(word1)
	cols := len(word2)

	// 异常情况构造
	if rows*cols == 0 {
		return rows + cols
	}
	// 构造二维数组dp[rows+1][cols+1]
	var i, j int
	dp := constructMatrix(rows+1, cols+1)
	for i = 0; i <= rows; i++ {
		dp[i][0] = i
	}
	for j = 0; j <= cols; j++ {
		dp[0][j] = j
	}

	for i = 1; i <= rows; i++ {
		for j = 1; j <= cols; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			}
		}
	}
	return dp[rows][cols]
}

// 构造二维数组
func constructMatrix(rows, cols int) [][]int {
	m := make([][]int, rows)
	for i := 0; i < rows; i++ {
		m[i] = make([]int, cols)
	}
	return m
}

// 最小值求解
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
