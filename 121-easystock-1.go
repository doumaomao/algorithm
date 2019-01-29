// 假设有一个数组，其中第i个元素是第i天给定股票的价格。

// 如果只被允许完成最多一笔交易（即买入并卖出一股股票），请设计一个算法以找到最大利润。

// 该方法存在两层循环，时间复杂度较高

package algorithm

func maxProfit(prices []int) int {

	var max int = 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			tmp_max := prices[j] - prices[i]
			if tmp_max > max {
				max = tmp_max
			}
		}
	}
	return max
}
