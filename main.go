package main

import "fmt"

// the code below is used to create a struct named Stack that contains a slice
// for getting the values of the stack
type Stack struct {
	// the property below is used to create a slice for getting the slice of
	// items in the stack
	items []int
}

// the code below is used to create a method for adding the data to the stack

// the push method will take a pointer to the stack struct and is used to add data
// to the stack and this will not return anything

// the s pointer variable below of type stack is used to get access to the stack struct
func (s *Stack) push(value int) {
	// the code below is to use the append() method for adding the data to the
	// items slice in the stack structure

	// the append method takes two inputs first one is the slice to which we have to
	// add the items and the second one is the value that we need to add
	s.items = append(s.items, value)
}

// the code below is used to create a method that will remove the value at the
// end or top of the stack and then returns the removed value

// the pop method will take a pointer to the stack struct and is used to delete data
// to the stack and will return the removed item
func (s *Stack) pop() int {

	// the property below is used to get the last item to be removed from the items
	// slice in the stack struct
	var removedItem int

	// the code below is used to check that if the items slice is empty then
	// showing that items cannot be removed due to underflow condition
	if len(s.items) == 0 {
		fmt.Println("Underflow condition: Cannot pop since stack is empty")
	} else {

		// the property below is used to get the last item to be removed from the items
		// slice in the stack struct
		removedItem = s.items[len(s.items)-1] // this is used to get the last item
		// from the items slice in the stack struct

		// the code below is used to remove the last item from the items slice in the stack
		// struct

		// the below line of code iterates the items slice from 0th element to
		// the second last element ignoring the last item in the items slice
		s.items = s.items[:len(s.items)-1] // here we are going from the 0th element
		// to the second last element so removing the last item from the stack
	}
	// the code below is used to return the removedItem to the user
	return removedItem

}

// the code below is used to create the stack data structure in go
func main() {
	// the code below is used to create an instance of the stack struct for doing
	// push and pop operations in the stack
	stack := Stack{}

	// the code below is used for debugging purpose to print the initial value of
	// the stack variable
	fmt.Println("The initial data in the stack is:", stack)

	// the code below is used to call the push() method for adding the data to the
	// stack
	stack.push(100)
	stack.push(200)
	stack.push(300)

	// the code below is for debugging purpose
	fmt.Println("The data in the stack is:", stack)

	// the code below is used to call the pop method for removing the last item
	// from the stack and then saving the item to be removed from the items slice
	// in the stack struct in the variable
	itemDeleted := stack.pop()

	// the code below is used to print the item being removed from the items
	// slice in the stack struct
	fmt.Println("The item removed is:", itemDeleted)

	// the below line of code is for debugging purpose
	fmt.Println("The data in the stack after the pop operation is:", stack)
}
