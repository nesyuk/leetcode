package array

// all n in nums are in [1...n]
func findDuplicateRecursive(nums []int) int {
    return swap(nums, 0)
}

func swap(nums []int, idx int) int {
    if idx == nums[idx] {
        return idx
    }
    next := nums[idx]
    nums[idx] = idx
    return swap(nums, next)
}

func findDuplicateIterative(nums []int) int {
	idx := 0
	for ; nums[idx] != idx; {
		next := nums[idx]
		nums[idx] = idx
		idx = next
	}
	return idx
}
