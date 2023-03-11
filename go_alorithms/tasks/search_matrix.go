package tasks

//10 11 16 20
//23 26 30 35
//45 48 50 55

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m, n := len(matrix[0]), len(matrix)
	l, r := 0, m*n

	for l < r {
		mid := l + (r-l)/2
		if matrix[mid/m][mid%n] > target {
			mid = l + 1
		} else if matrix[mid/m][mid%n] < target {
			mid = r
		} else {
			return true
		}
	}
	return false
}
