package array

// all n in nums are in [1...n]
func findDuplicate(nums []int) int {
    for _, n := range nums {
        // treat number as an index, check if we visited number at this index or not
        // use "negative marking" to set number as "visited"
        idx := n
        if idx < 0 {
            idx *= -1
        }
        if nums[idx] < 0 {
            // index was visited, so return the number
            return idx
        }
        // mark as visited
        nums[idx] = -1 * nums[idx]
    }
    return -1
}
