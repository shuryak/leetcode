// Жадный алгоритм O(n)

func maxProfit(prices []int) int {
    buy := prices[0]
    maxProfit := 0
    for i := 1; i < len(prices); i++ {
        // Закупить лучше тогда, когда цена ниже
        if buy > prices[i] {
            buy = prices[i]
        } else if prices[i] - buy > maxProfit { // Если закупать не выгодно, значит нужно проверить, выгодно ли продавать
            maxProfit = prices[i] - buy
        }
    }
    return maxProfit
}
