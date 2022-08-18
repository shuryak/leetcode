func missingNumber(nums []int) int {
    var sum int
    for _, num := range nums {
        sum += num
    }
    
    // Ожидаемая сумма (арифметическая прогрессия) с n+1 элементом минус фактическая сумма
    return (len(nums) + 1) * len(nums) / 2 - sum
}
