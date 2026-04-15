package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	allArgs := os.Args[1:]
	firstArg := allArgs[0]
	afterArgs := allArgs[1:]
	argString := strings.Join(afterArgs, " ")

	if strings.ToLower(firstArg) == "add" {
		fmt.Printf("'%v' succesfully added to the list\n", argString)
	}
}
