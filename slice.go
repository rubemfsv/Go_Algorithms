// Write a program which prompts the user to enter integers and stores
//  the integers in a sorted slice. The program should be written as a loop.
//  Before entering the loop, the program should create an empty integer slice
//  of size (length) 3. During each pass through the loop, the program prompts
//  the user to enter an integer to be added to the slice. The program adds the
//  integer to the slice, sorts the slice, and prints the contents of the slice in sorted
//  order. The slice must grow in size to accommodate any number of integers
//  which the user decides to enter. The program should only quit (exiting the loop)
//  when the user enters the character ‘X’ instead of an integer.

package main

import (
	"fmt"
	"sort"
)

func main() {
	var ele rune
	var size int
	var sli = make([]int, 0, 3)
	size = cap(sli)
	for i := 0; i <= size; i++ {
		if i >= len(sli) {
			size = size + 1
		}
		ele = 0
		fmt.Println("Enter a number to add: ")
		fmt.Scan(&ele)
		if ele == 0 {
			fmt.Println("Stopping!")
			break
		}
		sli = append(sli, int(ele))
	}
	sort.Ints(sli)
	fmt.Println(sli)
}
