package piscine

import "ft"

func PrintStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
}