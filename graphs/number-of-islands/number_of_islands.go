package main

func NumberOfIslands(grid [][]byte) int {

	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	cont := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				cont++
				//grid[i][j] = '0'
				dfs(grid, i, j)
			}
		}

	}
	return cont
}

func dfs(grid [][]byte, r, c int) {

	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] == '0' {
		return
	}

	grid[r][c] = '0'

	dfs(grid, r-1, c)
	dfs(grid, r+1, c)
	dfs(grid, r, c-1)
	dfs(grid, r, c+1)
}
