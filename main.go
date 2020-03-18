package main

import (
	"fmt"
)

func main() {
	cli := NewClient()
	list, err := cli.GetList()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(list["git"])
}
