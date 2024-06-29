package piscine

import "ft"

func PrintSolve(solve [][]int) {
	for i := range solve {
		for j := range solve[i] {
			if solve[i][j]/10 == 2 {
				ft.PrintRune('B')
			} else if solve[i][j]/10 == 1 {
				ft.PrintRune('.')
			} else {
				ft.PrintRune(rune(solve[i][j] + '0'))
			}
		}
		ft.PrintRune('\n')
	}
}
