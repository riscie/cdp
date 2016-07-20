package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"os"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	clipboard.WriteAll(pwd)
}
