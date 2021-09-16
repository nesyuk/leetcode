package sorting

import (
	"math/rand"
	"time"
)

func QuickSort(array []int) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(array), func(i, j int){ array[i], array[j] = array[j], array[i] })
	quickSort(array, 0, len(array)-1)
}

func quickSort(array []int, left, right int) {
	if left >= right {
		return
	}
	partition := array[left]
	i, j := left + 1, right
	for i <= j {
		for ; i <= j && array[i] <= partition; i++ {}
		for ; j >= i && array[j] >= partition; j-- {}
		if i < j {
			array[i], array[j] = array[j], array[i]
		}
	}
	array[left], array[j] = array[j], array[left]
	quickSort(array, left, j-1)
	quickSort(array, j+1, right)
}
