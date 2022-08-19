func findDisappearedNumbers(nums []int) []int {
    for _, num := range nums {
        idx := int(math.Abs(float64(num))) - 1
        
        // Ответ в индексах
        if nums[idx] > 0 {
            nums[idx] *= -1
        }
    }
    
    var result []int
    for idx, num := range nums {
        if num > 0 {
            result = append(result, idx + 1)
        }
    }
    
    return result
}