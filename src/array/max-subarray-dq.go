package array

func MaxSubArrayDivideConquer(nums []int) int {
    return findMaxSubArrayDq(nums, 0, len(nums)-1)
}

func findMaxSubArrayDq(nums []int, fromIdx, toIdx int) int {
    if fromIdx == toIdx {
        return nums[fromIdx]
    }
    midIdx := (fromIdx + toIdx)/2
    max := nums[midIdx]
    if midIdx-1 >= fromIdx {
        maxLeftSum := findMaxSubArrayDq(nums, fromIdx, midIdx-1)
        if maxLeftSum > max {
            max = maxLeftSum
        }
    }
    if toIdx >= midIdx + 1 {
        maxRightSum := findMaxSubArrayDq(nums, midIdx + 1, toIdx)
        if maxRightSum > max {
            max = maxRightSum
        }
    }
    maxMidSum := nums[midIdx]
    
    maxLeft := 0
    currentLeft := 0
    
    for i := midIdx-1; i >= fromIdx; i-- {
        currentLeft += nums[i]
        if maxLeft < currentLeft {
            maxLeft = currentLeft
        }
    }
    
    maxRight := 0
    currentRight := 0
    
    for j := midIdx+1; j <= toIdx; j++ {
        currentRight += nums[j]
        if maxRight < currentRight {
            maxRight = currentRight
        }
    }
    
    maxMidSum = maxLeft + maxRight + nums[midIdx]
    if maxMidSum > max {
        return maxMidSum
    }
    return max
}

