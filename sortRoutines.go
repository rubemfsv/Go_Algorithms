// Write a program to sort an array of integers. The program should partition the array into 4 parts, 
// each of which is sorted by a different goroutine. Each partition should be of approximately equal 
// size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

// The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ 
// of the array should print the subarray that it will sort. When sorting is complete, the main goroutine
//  should print the entire sorted list.

package main
import (
		"fmt"
		"bufio"
		"os"
		"strings"
		"strconv"
		"sync")


func Swap(slice []int, i int) {
	aux := slice[i];
	slice[i] = slice[i+1];
	slice[i+1] = aux;
}
	
func BubbleSort(slice []int, wg *sync.WaitGroup) {
	defer wg.Done();
	var swaps int = 1;
	for (swaps > 0) {
		swaps = 0;
		for i:=0;i<(len(slice) - 1);i++ {
			if (slice[i] > slice[i+1]) {
				Swap(slice, i);
				swaps++;
			}
		}
	}
	fmt.Printf("Sorted slice: %v \n", slice);
}

type Sli struct {
	sl1 []int
	sl2 []int
	sl3 []int
	sl4 []int
}

func divideSlice (slice []int) Sli {
	n := len(slice);
	a := n/4;
	b := n/4;
	c := n/4;
	switch n%4 {
		case 1:
			a++;
		case 2:
			a++;
			b++;
		case 3:
			a++;
			b++;
			c++;
	}
	slice1 := slice[0:a];
	slice2 := slice[a:a+b];
	slice3 := slice[a+b:a+b+c];
	slice4 := slice[a+b+c:];
	sli := Sli{sl1: slice1, sl2: slice2, sl3: slice3, sl4: slice4};
	return sli;
}

func unifySlice(slice Sli) []int {
	x := append(slice.sl1, slice.sl2...);
	x = append(x, slice.sl3...);
	x = append(x, slice.sl4...);
	return x;
}

func main() {
	fmt.Println(" Enter a series of integers separated by spaces: ");
	fmt.Print( " > " );

	scanner := bufio.NewScanner(os.Stdin);
	scanner.Scan()
	line := strings.Split(scanner.Text(), " ");
	integers := make([]int, len(line));
	for i,v := range line {
		integers[i],_ = strconv.Atoi(v);
	}

	slices := divideSlice(integers);

	var wg sync.WaitGroup;
	wg.Add(4);

	go BubbleSort(slices.sl1, &wg);
	go BubbleSort(slices.sl2, &wg);
	go BubbleSort(slices.sl3, &wg);
	go BubbleSort(slices.sl4, &wg);
	wg.Wait();

	integers = unifySlice(slices);
	wg.Add(1);
	BubbleSort(integers, &wg);
	wg.Wait();
}