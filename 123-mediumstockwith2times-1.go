// 假设有一个数组，其中第i个元素是第i天给定股票的价格。

// 设计算法以找到最大利润,最多完成2次交易

// 不得同时进行多笔交易（即，必须在再次购买之前卖出股票)
package algorithm

func maxProfit(prices []int) int {

	// 判断是否为有效交易天数
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}

	// 第一次购买盈利(因为是购买所以盈利为负数)
	buy1 := -prices[0]
	// 第二次购买盈利
	buy2 := -prices[0]
	// 第一次出售盈利
	sell1 := 0
	// 第二次出售盈利
	sell2 := 0
	// for循环调整购买出售的值
	var cur_price int
	for i := 1; i < len(prices); i++ {
		// 获取当前price
		cur_price = prices[i]
		// 第一次购买盈利越大越好，表示买的越便宜越好
		buy1 = max(buy1, -cur_price)
		// 第一次售卖在当前sell1、当前价格与第一次购买盈利之和，2者取大
		sell1 = max(sell1, cur_price+buy1)
		// 第二次购买盈利越大越好，从上次盈利价格、第一次售卖盈利减去当前价格之差，2者取大
		buy2 = max(buy2, sell1-cur_price)
		// 第二次售卖盈利越大越好，从上次售卖盈利、当前售出价格与第二次购买盈利之和，2者取大
		sell2 = max(sell2, buy2+cur_price)
	}
	return sell2
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
