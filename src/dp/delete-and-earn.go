package dp

var MaxInt = int(uint(0) >> 1) 

// leetcode #740

func DeleteAndEarn(nums []int) int {
    counts := map[int]int{}
    minNum, maxNum := MaxInt, 0
    for _, n := range nums {
        counts[n]++
        minNum = min(minNum, n)
        maxNum = max(maxNum, n)
    }
    
    prevPrevState := value(counts, minNum)
    prevState := max(value(counts, minNum+1), prevPrevState)
    
    for n := minNum+2; n <= maxNum; n++ {
        prevPrevState, prevState = prevState, max(value(counts, n) + prevPrevState, prevState)
    }
    
    return prevState
}

func value(counts map[int]int, n int) int {
    return counts[n]*n
}

func max(v1, v2 int) int {
    if v1 > v2 {
        return v1
    }
    return v2
}
