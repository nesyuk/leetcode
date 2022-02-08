package array

// 287. Find the Duplicate Number

func findDuplicate2(nums []int) int {
    slow, fast := nums[0], nums[nums[0]]
    
    for slow != fast {
        slow = nums[slow]
        fast = nums[nums[fast]]
    }
    
    slow = nums[0]
    fast = nums[fast]
    
    for ; slow != fast; {
        slow = nums[slow]
        fast = nums[fast]
    }
    return slow
}
