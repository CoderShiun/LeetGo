package cells_with_odd_values_in_a_matrix

func OddCells(n int, m int, indices [][]int) int {
	arr := make([][]int, n)

	for i, _ := range arr {
		arr[i] = make([]int, m)
	}

	for _, v := range indices {
		x, y := v[0], v[1]

		for a := range arr[0] {
			arr[x][a] ++
		}

		for b := range arr {
			arr[b][y] ++
		}
	}

	odds := 0
	for i := range arr {
		for j := range arr[0] {
			if arr[i][j]%2 == 1 {
				odds++
			}
		}
	}
	return odds
}
