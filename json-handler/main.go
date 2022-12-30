package main

import (
	"fmt"
	"os"
)

func mainWithExit() int {
	defer fmt.Println("difiere")

	fmt.Println("funcionnnnnn")
	
	mitoken := os.Getenv("repotoken")
	fmt.Println("mitoken:", mitoken)
	if (mitoken == "porongoORG") {
		fmt.Println("OKKKKKKK")
	}

	return (0)

}

func main() {

	ret := mainWithExit()

	os.Exit(ret)
}
