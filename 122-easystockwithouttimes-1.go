// 假设有一个数组，其中第i个元素是第i天给定股票的价格。

// 设计算法以找到最大利润,可以根据需要完成尽可能多的交易（即，多次买入并卖出一股股票）。

package algorithm

func maxProfit(prices []int) int {
	if prices == nil {
		return 0
	}
	var max int = 0

	for i := 0; i < len(prices)-1; i++ {
		tmp_max := prices[i+1] - prices[i]
		if tmp_max > 0 {
			max = max + tmp_max
		}
	}
	return max
}
