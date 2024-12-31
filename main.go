package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	// subArray[0] changes testArray[3] because slices share the same underlying array
	var testArray [9]int32
	subArray := testArray[3:6]

	fmt.Printf("testArray[3] = %d\n", testArray[3]) // testArray[3] = 0

	subArray[0] = 1 // subArray[0] conneted to testArray[3]

	fmt.Printf("testArray[3] = %d\n", testArray[3]) // testArray[3] = 1

	subArray0Ptr := &subArray[0]
	fmt.Printf("subArray0Ptr = %d -> ", subArray0Ptr)
	fmt.Println(subArray0Ptr)
	fmt.Println(reflect.TypeOf(subArray0Ptr))

	// Get the address of testArray[2] by subtracting the size of one element from the pointer to subArray[0]
	testArray2Ptr := (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(subArray0Ptr)) - unsafe.Sizeof(testArray[0])))

	fmt.Printf("testArray2Ptr = %d -> ", testArray2Ptr)

	// Calculate the difference in memory addresses
	difference := uintptr(unsafe.Pointer(subArray0Ptr)) - uintptr(unsafe.Pointer(testArray2Ptr))

	fmt.Printf("Memory difference between subArray[0] and testArray[2]: %d bytes (%s)\n", difference, reflect.TypeOf(difference))

	fmt.Printf("Size of int32: %d bytes\n", unsafe.Sizeof(testArray[0]))

	fmt.Printf("testArray[2] = %d\n", testArray[2]) // testArray[2] = 0
	*testArray2Ptr = 2
	fmt.Printf("testArray[2] = %d\n", testArray[2]) // testArray[2] = 2

	// testSlice and subSlice points to different underlying arrays
	testSlice := []int{0, 0, 0, 0, 0, 0}
	subSlice := testSlice[3:6]

	// Append to testSlice cause reallocation
	testSlice = append(testSlice, 7)
	subSlice[0] = 1

	// testSlice[3] = 0, because the underlying array of testSlice reallocated,
	// and subSlice now references the original array, while testSlice points to a new array.
	fmt.Printf("testSlice[3] = %d\n", testSlice[3])
}
