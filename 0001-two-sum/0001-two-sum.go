func twoSum(nums []int, target int) []int {
    indexes := make(map[int]int) // (value): (index of this value in nums)

    for i := 0; i < len(nums); i++ {
        diff := target - nums[i]
        if diffIdx, ok := indexes[diff]; ok {
            return []int{diffIdx, i}
        } else {
            indexes[nums[i]] = i
        }
    }

    return []int{}
}
