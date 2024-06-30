package main

import (
	"os"
	"piscine"
)

func main() {
	args := os.Args
	if piscine.CheckArgs(args) == 0 {
		piscine.PrintStrLn("Error")
		return
	}
	solve := make([][]int, 5)
	for i := range solve {
		solve[i] = make([]int, 5)
	}
	piscine.SetSolve(args[1:], solve)
	if !piscine.FindSolve(solve) {
		piscine.PrintStrLn("Error")
		return
	}
	piscine.PrintSolve(solve)
}
