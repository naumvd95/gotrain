package main

type Stack []int

func (s *Stack) push(v int) {
	*s = append(*s, v)
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) pop() (int, bool) {
	var res int
	if s.isEmpty() {
		return res, false
	}

	topIdx := len(*s) - 1
	res = (*s)[topIdx]
	// cutoff tail
	(*s) = (*s)[:topIdx]

	return res, true
}

func (s *Stack) top() (int, bool) {
	var res int
	if s.isEmpty() {
		return res, false
	}

	// do not cutoff tail, just return it
	topIdx := len(*s) - 1
	res = (*s)[topIdx]
	return res, true
}

/* algo:
1. create stack and hanlde in it 'k' numbers only
2. after filling in 'k' numbers, you need to compare current num with
value in stack and switch them if necessary
3. Note that you can switch not only last num from stack, but any number of nums, if you have enough numbers to exchange in your
original array
*/
func getMaxNumber(a []int, k int) int {
	var res int
	var s Stack

	stackCounter := 0
	for i := 0; i < len(a); i++ {
		currentNum := a[i]
		// fill in stack until we reach k elements
		if stackCounter < k {
			s.push(currentNum)
			continue
		}

		//otherwise we need to compare the numbers
		allowedCompareDepth := len(a) - i - 1 // how many numbers left

	}

	return res
}

func main() {
	input := []int{7, 2, 5, 9}
}
