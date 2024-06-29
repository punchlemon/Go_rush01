package piscine

func SetSolve(args []string, solve [][]int) {
	for i, s := range args {
		for j, c := range s {
			if c == '.' {
				solve[i][j] = 10
			} else {
				solve[i][j] = int(c - '0')
			}
		}
	}
}
