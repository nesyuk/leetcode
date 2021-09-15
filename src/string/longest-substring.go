package string

func LengthOfLongestSubstring(s string) int {
    maxLen := 0
    seen := map[rune]int{}
    runes := []rune(s)
    
    for i, j := 0, 0; j < len(s); {        
        seen[runes[j]]++
        for seen[runes[j]] > 1 {
                seen[runes[i]]--
                i++
        }
        j++
        if maxLen < j - i {
            maxLen = j - i
        }
    }
    return maxLen
}
