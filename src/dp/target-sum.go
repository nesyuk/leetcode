package dp

func FindTargetSumWays(nums []int, target int) int {
    maxSum := 0
    for _, n := range nums {
        maxSum += n
    }
    if target > maxSum || target < -1*maxSum {
        return 0
    }
    sums := make([]int, 2*maxSum + 1)
    sums[nums[0] + maxSum] = 1
    sums[-1*nums[0] + maxSum] += 1
    
    for numIdx := 1; numIdx < len(nums); numIdx++ {
        num := nums[numIdx]
        sumsCopy := make([]int, 2*maxSum + 1)
        for sumIdx := 0; sumIdx < 2*maxSum + 1; sumIdx++ {
            if sums[sumIdx] > 0 {
                prevSum := sumIdx - maxSum
                sumsCopy[prevSum + num + maxSum] += sums[sumIdx]
                sumsCopy[prevSum - num + maxSum] += sums[sumIdx]
            }
        }
        sums = sumsCopy
    }
    return sums[target + maxSum]
}
