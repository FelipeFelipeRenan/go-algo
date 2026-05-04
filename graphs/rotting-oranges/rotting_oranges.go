package main

func orangesRotting(grid [][]int) int {

	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}

	goodOranges := 0
	badOranges := [][2]int{}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				goodOranges += 1
			} else if grid[i][j] == 2 {
				badOranges = append(badOranges, [2]int{i, j})
			}
		}
	}

	minutes := 0

	for len(badOranges) > 0 && goodOranges > 0 {

		minutes++
		tamanhoLote := len(badOranges)

		for i := 0; i < tamanhoLote; i++ {
			atual := badOranges[0]
			badOranges = badOranges[1:]

			r := atual[0]
			c := atual[1]

			if r-1 >= 0 && grid[r-1][c] == 1 {
				grid[r-1][c] = 2
				goodOranges--
				badOranges = append(badOranges, [2]int{r - 1, c})
			}

			if r+1 < len(grid) && grid[r+1][c] == 1 {
				grid[r+1][c] = 2
				goodOranges--
				badOranges = append(badOranges, [2]int{r + 1, c})
			}

			if c-1 >= 0 && grid[r][c-1] == 1 {
				grid[r][c-1] = 2
				goodOranges--
				badOranges = append(badOranges, [2]int{r, c - 1})
			}

			if c+1 < len(grid[0]) && grid[r][c+1] == 1 {
				grid[r][c+1] = 2
				goodOranges--
				badOranges = append(badOranges, [2]int{r, c + 1})
			}
		}
	}
	if goodOranges > 0 {
		return -1
	}
	return minutes
}
