package main

import "fmt"

func main() {
	var name string
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	fmt.Print("1")
}
