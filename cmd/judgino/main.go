package main

import (
	"fmt"
	"os"

	"github.com/SUT-technology/judgino/cmd/judgino/command"
)

func main() {
	err := command.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
