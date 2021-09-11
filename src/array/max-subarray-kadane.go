package array

func MaxSubArrayKadane(nums []int) int {
	runningSum := nums[0]
	maxSum := runningSum
	for idx := 1; idx < len(nums); idx++ {
		if nums[idx] < nums[idx] + runningSum {
			runningSum += nums[idx]
		} else {
			runningSum = nums[idx]
		}
		if maxSum < runningSum {
			maxSum = runningSum
		}
	}
	return maxSum
}

