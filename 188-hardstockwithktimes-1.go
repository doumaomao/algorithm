// 假设有一个数组，其中第i个元素是第i天给定股票的价格。

// 设计算法以找指定K次交易下的最大利润

// 不得同时进行多笔交易（即，必须在再次购买之前卖出股票)

package algorithm

func maxProfit(k int, prices []int) int {
	// 交易总天数
	n := len(prices)
	// 如果k次大于交易总天数的一半，则表示以尽可能多的交易获取利润即可
	if k >= (n >> 1) {
		return buyAsMany(prices)
	}
	// 构造二维数组dp[k+1][n]
	dp := constructMatrix(k+1, n)
	// 按照推导后的公式思路进行循环
	tmp := 0
	for i := 1; i < k+1; i++ {
		tmp = -prices[0]
		for j := 1; j < n; j++ {
			dp[i][j] = max(dp[i][j-1], prices[j]+tmp)
			tmp = max(tmp, dp[i-1][j-1]-prices[j])
		}

	}
	return dp[k][n-1]
}

// 尽可能买多次的利润
func buyAsMany(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	profit := 0
	for i := 1; i < len(prices); i++ {
		profit = profit + max(0, prices[i]-prices[i-1])
	}
	return profit
}

// 构造二维数组
func constructMatrix(rows, cols int) [][]int {
	m := make([][]int, rows)
	for i := 0; i < rows; i++ {
		m[i] = make([]int, cols)
	}
	return m
}

// 最大值求解
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
