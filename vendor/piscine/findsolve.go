package piscine

func CheckCell(n int, count *int) bool {
	if n > 0 && n < 20 {
		(*count)++
	} else if n > 19 {
		return false
	}
	return true
}

func CheckRowCol(solve [][]int, i, j int) bool {
	count := 0
	for k := 0; k <= 4-i; k++ {
		if !CheckCell(solve[i+k][j], &count) {
			break
		}
	}
	for k := -1; -i <= k; k-- {
		if !CheckCell(solve[i+k][j], &count) {
			break
		}
	}
	for k := 1; k <= 4-j; k++ {
		if !CheckCell(solve[i][j+k], &count) {
			break
		}
	}
	for k := -1; -j <= k; k-- {
		if !CheckCell(solve[i][j+k], &count) {
			break
		}
	}
	return count == solve[i][j]
}

func CheckSolve(solve [][]int) bool {
	for i := range solve {
		for j := range solve[i] {
			if solve[i][j] > 0 && solve[i][j] < 10 && !CheckRowCol(solve, i, j) {
				return false
			}
		}
	}
	return true
}

func CountUpSolve(solve [][]int) bool {
	for i := range solve {
		for j := range solve[i] {
			if solve[i][j] > 19 {
				solve[i][j] = 10
			} else if solve[i][j] == 10 {
				solve[i][j] = 20;
				return true
			}
		}
	}
	return false
}

func FindSolve(solve [][]int) bool {
	for {
		res := !CountUpSolve(solve)
		if CheckSolve(solve) {
			return true
		} else if res {
			return false
		}
	}
}
