package main

import (
	"flag"
	"fmt"
	"os"
)

func mainWithExit() int {
	defer fmt.Println("difiere")

	fmt.Println("funcionnnnnn")

	mitoken := os.Getenv("repotoken")
	fmt.Println("mitoken:", mitoken)
	if mitoken == "porongoORG" {
		fmt.Println("OKKKKKKK")
	}

	return (0)

}

func main() {
	json_flag := flag.String("json", "default", "json content")
	is_PR := flag.String("pr", "false", "identifies if the triggered Github event is for a PR")

	flag.Parse()

	fmt.Printf("json_flag:%s is_PR:%s\n", *json_flag, *is_PR)

	ret := mainWithExit()

	os.Exit(ret)
}
