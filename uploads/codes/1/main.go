package main

import (
	"fmt"
	"time"
)

func main() {
	var name string
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	time.Sleep(time.Second * 5)
	fmt.Println("1")
}
