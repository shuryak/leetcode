func maxProfit(prices []int) int {
	var profit int
    maxIdx := findIndexOfMax(prices)
	for idx, val := range prices {
		if idx == len(prices)-1 {
			break
		}
        if prices[idx] == prices[idx+1] {
            continue
        }
        if idx >= maxIdx {
            maxIdx = idx+1+findIndexOfMax(prices[idx+1:])
        }
		pre := prices[maxIdx] - val
		if pre > profit {
			profit = pre
		}
	}

	return profit
}

func findIndexOfMax(nums []int) int {
	resVal := nums[0]
	resIdx := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] > resVal {
			resVal = nums[i]
			resIdx = i
		}
	}

	return resIdx
}
