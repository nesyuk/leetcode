package binary_search

func FindFirstBrokenBuild(builds []int) int {
	fromIdx := 0
	toIdx := len(builds) - 1

	for fromIdx < toIdx {
		midIdx := fromIdx + (toIdx - fromIdx) / 2

		// if build not broken, set toIdx to mid
		if builds[midIdx] == 1 {
			toIdx = midIdx
		} else {
			fromIdx = midIdx + 1
		}
	}
	if fromIdx == toIdx && builds[fromIdx] != 1 {
		return -1
	}
	return fromIdx
}
