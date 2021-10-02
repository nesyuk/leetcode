package binary_search

func findPeakElement2(nums []int) int {
    left := 0
    right := len(nums) - 1
    for left + 1 < right {
        mid := left + (right - left)/2
        if nums[mid] < nums[mid+1] {
                left = mid
        } else if nums[mid] < nums[mid-1] {
                right = mid
        } else {
            return mid
        }
    }
    if nums[left] < nums[right] {
        return right
    }
    return left
}
