func missingNumber(nums []int) int {
    var result int
    
    for i, num := range nums {
        result ^= i ^ num
    }
    
    result ^= len(nums)
    
    return result
}
