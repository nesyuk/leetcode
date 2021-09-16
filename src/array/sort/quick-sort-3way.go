package sorting

import (
	"math/rand"
	"time"
)

func QuickSort3Way(array []int) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(array), func(i, j int){ array[i], array[j] = array[j], array[i] })
	quickSort3Way(array, 0, len(array)-1)
}

func quickSort3Way(array []int, start, end int) {
	if start >= end {
		return
	}
	partition := array[start]
	left, current, right := start, start + 1, end
	for current <= right {
		if array[current] < partition {
			array[current], array[left] = array[left], array[current]
			left++
			current++
		} else if array[current] > partition {
			array[current], array[right] = array[right], array[current]
			right--
		} else {
			current++
		}
	}
	quickSort3Way(array, start, left - 1)
	quickSort3Way(array, right + 1, end)
}
