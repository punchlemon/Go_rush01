package main

import (
	"os"
	"piscine"
)

func main() {
	args := os.Args
	count := 0
	for range args {
		count++
	}
	if count != 6 {
		piscine.PrintStr("Error")
	} else {
		for _, s := range args[1:] {
			piscine.PrintStr(s)
			piscine.PrintStr("\n")
		}
	}
}