package dp

// leetcode #198

func HouseRobberTopDown(nums []int) int {
    memo := map[int]int{}
    return robDpFn(&nums, len(nums)-1, &memo)
}

func robDpFn(nums *[]int, i int, memo *map[int]int) int {
    if val, exist := (*memo)[i]; exist {
        return val
    }

    cur := (*nums)[i]
    if i >= 2 {
        cur += robDpFn(nums, i-2, memo)
    }
    (*memo)[i] = cur
    if i >= 1 {
        prev := robDpFn(nums, i-1, memo)
        (*memo)[i-1] = prev
        if cur < prev {
            return prev
        }
    }
    return cur
}

func HouseRobberBottomUp(nums []int) int {
    max := make([]int, len(nums))
    max[0] = nums[0]
    max[1] = nums[1]
    max[2] = nums[2] + nums[0]
    for i := 3; i < len(nums); i++ {
        max[i] = nums[i]
        if max[i-2] > max[i-3] {
            max[i] += max[i-2] 
        } else {
            max[i] += max[i-3]
        }
    }
    if max[len(max)-1] > max[len(max)-2] {
        return max[len(max)-1]
    }
    return max[len(max)-2]    
}
