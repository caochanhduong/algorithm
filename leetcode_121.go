func maxProfit(prices []int) int {
    dayCnt := len(prices)
    if dayCnt == 0 || dayCnt == 1 {
        return 0
    }
    buy := 0
    sell := 1
    profit := 0
    var diff int = prices[sell] - prices[buy]
    for sell < dayCnt {
        if prices[buy] < prices[sell] {
            diff = prices[sell] - prices[buy]
            if diff > profit {
                profit = diff
            }
            sell++ 
            continue
        }
        buy = sell
        sell = buy + 1
    }
    return profit
}
