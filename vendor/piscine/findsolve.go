package piscine

func CheckCell(n int, count *int) bool {
	if n <= 1 || n >= 20 {
		return false
	}
	(*count)++
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

func ContinuationUp(solve [][]int, i, j, up int) {
	if solve[i][j] > 1 && solve[i][j] < 20 {
		solve[i][j] += up
		if i > 0 {
			ContinuationUp(solve, i-1, j, up)
		}
		if i < 4 {
			ContinuationUp(solve, i+1, j, up)
		}
		if j > 0 {
			ContinuationUp(solve, i, j-1, up)
		}
		if j < 4 {
			ContinuationUp(solve, i, j+1, up)
		}
	}
}

func CheckConnect(solve [][]int) bool {
	var res bool = true
	var up int = 100
	var col int = -1
	var row int = 0
	for {
		if col < 4 {
			ContinuationUp(solve, col+1, row, up)
			break
		} else if row < 4 {
			ContinuationUp(solve, 0, row+1, up)
			break
		} else {
			return res
		}
	}
	for i := range solve {
		for j := range solve[i] {
			if solve[i][j] > 1 && solve[i][j] < 20 {
				res = false
			}
		}
	}
	for i := range solve {
		for j := range solve[i] {
			if solve[i][j] > up {
				solve[i][j] -= up
			}
		}
	}
	return res
}

func CheckSolve(solve [][]int) bool {
	for i := range solve {
		for j := range solve[i] {
			if solve[i][j] > 0 && solve[i][j] < 10 && !CheckRowCol(solve, i, j) {
				return false
			}
			if solve[i][j] == 20 {
				if (i > 0 && solve[i-1][j] == 20) ||
					(i < 4 && solve[i+1][j] == 20) ||
					(j > 0 && solve[i][j-1] == 20) ||
					(j < 4 && solve[i][j+1] == 20) {
					return false
				}
			}
		}
	}
	if !CheckConnect(solve) {
		return false
	}
	return true
}

func CountUpSolve(solve [][]int) bool {
	for i := range solve {
		for j := range solve[i] {
			if solve[i][j] > 19 {
				solve[i][j] = 10
			} else if solve[i][j] == 10 {
				solve[i][j] = 20
				return true
			}
		}
	}
	return false
}

func FindSolve(solve [][]int) bool {
	for {
		if CheckSolve(solve) {
			return true
		}
		if !CountUpSolve(solve) {
			return false
		}
	}
}
