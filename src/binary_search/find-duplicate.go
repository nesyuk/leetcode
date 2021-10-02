package binary_search

import "sort"

func findDuplicate(nums []int) int {
    copyNums := nums
    sort.Ints(copyNums)
    left, right := 0, len(copyNums) - 1
    for left < right {
        mid := left + (right - left)/2
        if copyNums[mid] > mid {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return copyNums[left]
}
