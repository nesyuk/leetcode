package dp

// leetcode #1137

func TribonacciBottomUp(n int) int {
    if n  == 0 {
        return 0
    }
    if n < 3 {
        return 1
    }
    t1, t2, t3 := 0, 1, 1
    for i := 3; i <= n; i++ {
        t1, t2, t3 = t2, t3, t3 + t2 + t1
    }
    return t3
}

func TribonacciTopDown(n int) int {
    memo := map[int]int{}
    return tribonacci(n, &memo)
}

func tribonacci(n int, memo *map[int]int) int {
    if t, exist := (*memo)[n]; exist {
        return t
    }
    if n  == 0 {
        return 0
    }
    if n < 3 {
        return 1
    }
    (*memo)[n] = tribonacci(n-1, memo) + tribonacci(n-2, memo) + tribonacci(n-3, memo)
    return (*memo)[n] 
}
