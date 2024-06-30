package piscine

func CheckArgs(args []string) int {
	if len(args) != 6 {
		return 0
	}
	for _, s := range args[1:] {
		count := 0
		for _, c := range s {
			count++
			if c != '.' && (c < '2' || c > '9') {
				return 0
			}
		}
		if count != 5 {
			return 0
		}
	}
	return 1
}
