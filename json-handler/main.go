package main

import (
	"fmt"
	"os"
)

func mainWithExit() int {
	defer fmt.Println("difiere")

	fmt.Println("funcionnnnnn")

	return (35)

}

func main() {

	ret := mainWithExit()

	os.Exit(ret)
}
