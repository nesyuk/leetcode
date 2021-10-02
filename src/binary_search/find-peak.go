package binary_search

import "math"

func findPeakElement(nums []int) int {
        left := 0
        right := len(nums) - 1
        for left < right {
                mid := left + (right - left)/2
                if nums[mid] > nums[mid+1] {
                        right = mid
                } else {
                        left = mid + 1
                }
        }
        return left
}


func findPeakElementNaive(nums []int) int {
    nums = append([]int{math.MinInt64}, nums...)
    nums = append(nums, math.MinInt64)
    left := 1
    right := len(nums)-2

    for left < right {
        idx := left + (right - left) / 2
        if nums[idx] <= nums[idx+1] {
            left = idx+1
        } else if nums[idx] <= nums[idx-1] {
            right = idx-1
        } else {
            return idx-1
        }
    }
    return left-1
}
