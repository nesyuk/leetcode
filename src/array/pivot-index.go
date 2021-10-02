func PivotIndex(nums []int) int {
    totalSum := 0
    for _, num := range nums {
        totalSum += num
    }
    leftSum := 0
    for idx := 0; idx < len(nums); idx++ {
        if leftSum == totalSum - leftSum - nums[idx] {
            return idx
        }
        leftSum += nums[idx]
    }
    return -1
}
