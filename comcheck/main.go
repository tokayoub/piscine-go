package main

import (
	"fmt"
	"os"
)

func main() {
	gg := os.Args[1:]
	for _, v := range gg {
		if v == "01" || v == "galaxy" || v == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
