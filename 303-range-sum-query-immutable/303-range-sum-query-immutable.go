package main

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	result := NumArray{
		sums: make([]int, len(nums)),
	}

	result.sums[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		result.sums[i] = result.sums[i-1] + nums[i]
	}

	return result
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.sums[right]
	}

	// Всю сумму левее left вычитаем из результата
	return this.sums[right] - this.sums[left-1]
}
