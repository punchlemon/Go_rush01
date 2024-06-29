package piscine

import "ft"

func PrintStrLn(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}
