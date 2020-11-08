package main

import "fmt"

/*
Given a singly linked list of integers l and an integer k, remove all elements from list l that have a value equal to k.

Note: Try to solve this task in O(n) time using O(1) additional space, where n is the number of elements in the list,
since this is what you will ll be asked to do during an interview.
*/

//TODO describe pros/cons of such pattern
/*

LL versus Slice/Array:
+ you dont need to predefine how big is your list/array
+ append is less performance consuming

- size predefining problem resolved by using slices
- slice == object that consists of : pointer to array, length, and capasity (max length)


*/

// ListNode is an implementation of singly linked list
type ListNode struct {
	Value int
	Next  *ListNode
}

func removeKFromList(providedList *ListNode, ValueToRemove int) *ListNode {
	var nodeStorage ListNode        // define sinlgy linked list object to perform `for loop`
	nodeStorage.Next = providedList // save our current pointer as next Node !!!
	traverser := &nodeStorage       // pointer to our storage, needed to manage .Next connections !!!

	for providedList != nil { // next->next loop until reaching Next element of *ListNode, equals nil, means tail
		if providedList.Value == ValueToRemove {
			// using traverser, we rearrage `Next` pointer, skipping ValueToRemove
			traverser.Next = providedList.Next
			// going futher via our main loop
			providedList = providedList.Next
		} else {
			// not needed to rearrange, keep pointers as is
			traverser = providedList
			// going futher via our main loop
			providedList = providedList.Next
		}
	}

	// traverser successfully rearranged `Next` chains in our nodeStorage
	return nodeStorage.Next
}

func (providedList *ListNode) appendToLL(ValueToAppend int) {
	val := ListNode{
		Value: ValueToAppend,
	}
	fmt.Printf("Hey, planning to append element %v to our LL %v\n", val, providedList)

	// if no head
	if providedList == nil {
		providedList = &val
		fmt.Printf("Hey, this LL was empty, we are first to append, viola: %v\n", providedList)
		return
	}

	// DO NOT FORGET TO redefine pointers to iterate w/ loop
	for n := providedList; n != nil; n = n.Next {
		if n.Next == nil {
			n.Next = &val

			fmt.Printf("Hey, we are found tail! Added element to our LL %v\n", n)
			return
		}
		fmt.Println("Searching, sniffing tail....")
		fmt.Printf("here is next: %v\n", n.Next)
	}
}

func (providedList *ListNode) printLL() {
	for e := providedList; e != nil; e = e.Next {
		fmt.Printf("Here is element of LL: %v\n", e.Value)
	}
}

func main() {
	basicLinkedList := &ListNode{}
	fmt.Printf("Here is empty ll: %v\n", basicLinkedList)

	basicLinkedList.appendToLL(1)
	basicLinkedList.appendToLL(2)
	basicLinkedList.appendToLL(3)
	basicLinkedList.appendToLL(4)
	basicLinkedList.printLL()

	fmt.Println("Now we delete el equals 3")
	basicLinkedList = removeKFromList(basicLinkedList, 3)
	basicLinkedList.printLL()
}
