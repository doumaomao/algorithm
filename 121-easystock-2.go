// 假设有一个数组，其中第i个元素是第i天给定股票的价格。

// 如果只被允许完成最多一笔交易（即买入并卖出一股股票），请设计一个算法以找到最大利润。

// 一层循环，通过一层循环标记最低值以及最大差价来寻求最大利润
package algorithm

func maxProfit(prices []int) int {
	// 存储最大值的变量 初始值为0
	ans := 0

	// 判断参数长度是否为0
	if len(prices) == 0 {
		return ans
	}

	// 判断哪日买入
	bought := prices[0]

	// 遍历所有交易日价格
	for i := 0; i < len(prices); i++ {
		// 判断本日是否能够卖出
		if prices[i] > bought {
			// 判断如果本日卖出收益是否最大
			if ans < (prices[i] - bought) {
				ans = prices[i] - bought
			}
		} else {
			// 本日为最低价格
			bought = prices[i]
		}

	}
	return ans
}
