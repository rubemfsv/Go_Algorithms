// As part of this program, you should write a function called BubbleSort() which takes a slice of 
// integers as an argument and returns nothing. The BubbleSort() function should modify the slice 
// so that the elements are in sorted order.

// A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position
//  of two adjacent elements in the slice. You should write a Swap() function which performs this
//   operation. Your Swap() function should take two arguments, a slice of integers and an index value
//    i which indicates a position in the slice. The Swap() function should return nothing, but it should 
//    swap the contents of the slice in position i with the contents in position i+1.

package main
 
import "fmt"
 
func main() {
 
	slice, err := scanSlice(10)
	
	if err != nil {
        fmt.Println(err)
        return
	} else {
		fmt.Println("\n--- Unsorted --- \n\n", slice)
		bubblesort(slice)
		fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
	}
}
 
// Scan a slice of size 10
func scanSlice(n int) ([]int, error) {
    x := make([]int, n)
    y := make([]interface{}, len(x))
    for i := range x {
        y[i] = &x[i]
    }
    n, err := fmt.Scanln(y...)
    x = x[:n]
    return x, err
}

//BubbleSort algorithm
func bubblesort(items []int) {
    var (
        n = len(items)
        sorted = false
    )
    for !sorted {
        swapped := false
        for i := 0; i < n-1; i++ {
            if items[i] > items[i+1] {
                items[i+1], items[i] = items[i], items[i+1]
                swapped = true
            }
        }
        if !swapped {
            sorted = true
        }
        n = n - 1
    }
}