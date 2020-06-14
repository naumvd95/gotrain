package mergesort

import "testing"

var testDataset []int

// generating fake dataset
func init() {
	for i := 0; i < 1000000; i++ {
		testDataset = append(testDataset, i)
	}
}

/*
Functions of the form
func BenchmarkXxx(*testing.B)
are considered benchmarks, and are executed by the "go test" command when its -bench flag is provided.
*/
//BenchmarkMergeSortSingleThread tests performance of MergeSortSingleThread
func BenchmarkMergeSortSingleThread(b *testing.B) {
	// N means number of times to run benchmark
	for n := 0; n < b.N; n++ {
		MergeSortSingleThread(testDataset)
	}
}

//BenchmarkMergeSortMultiThread tests performance of MergeSortMultiThread
func BenchmarkMergeSortMultiThread(b *testing.B) {
	// N means number of times to run benchmark
	for n := 0; n < b.N; n++ {
		MergeSortMultiThread(testDataset)
	}
}
