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
	if count != 6 || piscine.CheckArgs(args) == 0 {
		piscine.PrintStrLn("Error")
	} else {
		for _, s := range args[1:] {
			piscine.PrintStrLn(s)
		}
	}
}