package questions

func numIslands(grid [][]byte) int {

	if grid == nil || len(grid) == 0 {
		return 0
	}

	rowLength := len(grid)
	columnLength := len(grid[0])
	numberOfIslands := 0
	for r := 0; r < rowLength; r++ {
		for c := 0; c < columnLength; c++ {
			if grid[r][c] == '1' {
				numberOfIslands++
				dfs(grid, r, c)
			}
		}
	}
	return numberOfIslands
}

func dfs(grid [][]byte, r int, c int) {
	rowLength := len(grid)
	columnLength := len(grid[0])
	if r < 0 || c < 0 || r >= rowLength || c >= columnLength || grid[r][c] == '0' {
		return
	}
	grid[r][c] = '0'
	dfs(grid, r-1, c)
	dfs(grid, r+1, c)
	dfs(grid, r, c-1)
	dfs(grid, r, c+1)
}
