package islands

func numIslands(grid [][]byte) int {
	// create and size visited grid to keep track of visited nodes
	visited := make([][]bool, len(grid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}
	islands := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				dfs(i, j, grid, visited)
				islands++
			}
		}
	}
	return islands
}

// dfs checks adjacent nodes ( N, S, W, E) to see if they are 'land' and will mark them visited,
//it will do this recursively but may be done iterative with a stack, ( or bfs with a queue)
func dfs(i, j int, grid [][]byte, visited [][]bool) {
	// base case here
	// if out of bounds, return
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || visited[i][j] || grid[i][j] == '0' {
		return
	}
	// rest
	visited[i][j] = true
	dfs(i+1, j, grid, visited)
	dfs(i-1, j, grid, visited)
	dfs(i, j+1, grid, visited)
	dfs(i, j-1, grid, visited)
}
