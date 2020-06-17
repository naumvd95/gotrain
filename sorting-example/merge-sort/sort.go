package sort

import (
	"sync"

	"github.com/naumvd95/gotrain/common"
)

func merge(leftChunk, rightChunk []common.CovidUnit) []common.CovidUnit {
	// control size of resulting slice
	res := make([]common.CovidUnit, 0, len(leftChunk)+len(rightChunk))

	// while one of chunks is not empty
	for len(leftChunk) > 0 || len(rightChunk) > 0 {
		// if left already empty, merge all right part
		if len(leftChunk) == 0 {
			return append(res, rightChunk...)
		}
		// if right already empty, merge all left part
		if len(rightChunk) == 0 {
			return append(res, leftChunk...)
		}
		if leftChunk[0].DeathByDate <= rightChunk[0].DeathByDate {
			res = append(res, leftChunk[0])
			// cut off calculated part
			leftChunk = leftChunk[1:]
		} else {
			res = append(res, rightChunk[0])
			// cut off calculated part
			rightChunk = rightChunk[1:]
		}
	}

	return res
}

//MergeSortSingleThread Typical divide->sort->merge alghoritm in single thread mode
func MergeSortSingleThread(dataset []common.CovidUnit) []common.CovidUnit {
	// nothing to sort
	if len(dataset) <= 1 {
		return dataset
	}

	initialDivider := len(dataset) / 2
	var leftChunk []common.CovidUnit
	var rightChunk []common.CovidUnit

	// recursive division
	leftChunk = MergeSortSingleThread(dataset[:initialDivider])
	rightChunk = MergeSortSingleThread(dataset[initialDivider:])

	//merge sorted chunks
	return merge(leftChunk, rightChunk)
}

// create buffered channel up to 100 routines
var sem = make(chan struct{}, 100)

//MergeSortMultiThread Typical divide->sort->merge alghoritm in multi thread mode
func MergeSortMultiThread(dataset []common.CovidUnit) []common.CovidUnit {
	// nothing to sort
	if len(dataset) <= 1 {
		return dataset
	}

	// start initial division of big dataset
	initialDivider := len(dataset) / 2
	var leftChunk []common.CovidUnit
	var rightChunk []common.CovidUnit
	// init wait group to sync states from goroutines
	wg := sync.WaitGroup{}
	wg.Add(2) // 2 threads for left/rigth chunks

	// select like case/switch but for channels
	select {
	// while sem is not full and at least 1 struct(routine) can be forwarded
	case sem <- struct{}{}:
		go func() {
			// recursive division
			leftChunk = MergeSortMultiThread(dataset[:initialDivider])
			wg.Done() // decrement initial value of wg.Add to confirm routine finish
		}()
	// otherwise go w/ single thread
	default:
		// recursive division
		leftChunk = MergeSortMultiThread(dataset[:initialDivider])
		wg.Done() // decrement initial value of wg.Add to confirm routine finish
	}

	select {
	// while sem is not full and at least 1 struct(routine) can be forwarded
	case sem <- struct{}{}:
		go func() {
			rightChunk = MergeSortMultiThread(dataset[initialDivider:])
			wg.Done() // decrement initial value of wg.Add to confirm routine finish
		}()
	// otherwise go w/ single thread
	default:
		rightChunk = MergeSortMultiThread(dataset[initialDivider:])
		wg.Done() // decrement initial value of wg.Add to confirm routine finish
	}

	wg.Wait() // wait decrementing to 0 , i.e. all routines finished
	//merge sorted chunks
	return merge(leftChunk, rightChunk)
}
