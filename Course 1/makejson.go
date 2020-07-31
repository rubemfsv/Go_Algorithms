// Write a program which prompts the user to first enter a name, and then enter an address. 
// Your program should create a map and add the name and address to the map using the keys 
// “name” and “address”, respectively. Your program should use Marshal() to create a JSON
//  object from the map, and then your program should print the JSON object.
package main

import (
	"fmt"
	"encoding/json"
	"os"
	"bufio"
)
	
func main() {
	var user map[string]string
	user = make(map[string]string)
	fmt.Printf("Insert your name:\n")
	br := bufio.NewReader(os.Stdin)
	bname, _, _ := br.ReadLine()
	name := string(bname)
	user["name"]=name
	fmt.Printf("Insert your address:\n")
	br1 := bufio.NewReader(os.Stdin)
	baddress, _, _ := br1.ReadLine()
	address := string(baddress)
	user["address"]=address
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Encoding failed!")
	} else {
		fmt.Println("Encoded data: ", b)
		fmt.Println("Decoded data: ", string(b))
	}

}
