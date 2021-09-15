package bit_manipulation

func KthGrammar(n int, k int) int {
    if n == 1 {
        return 0
    }
    return 1 - (KthGrammar(n-1, (k + 1) >> 1) ^ (k % 2))
}
