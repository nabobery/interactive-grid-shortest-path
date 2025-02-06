package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// DFS-related functions have been moved to this file.
// NOTE: We are not using DFS in production because its worst-case time complexity is exponential (approximately O(4^n))
// which makes it impractical for larger grids. BFS, on the other hand, runs in O(n²) on a 20×20 grid.

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func dfs_helper(x1 int, y1 int, x2 int, y2 int, currPath, minPath *[][2]int, minDist *int, currDist int, visited [][]bool) {
	if x1 < 0 || y1 < 0 || x1 >= 20 || y1 >= 20 || visited[x1][y1] {
		return
	}

	visited[x1][y1] = true
	*currPath = append(*currPath, [2]int{x1, y1})

	fmt.Printf("Visiting (%d, %d), dist=%d, minDist=%d\n", x1, y1, currDist, *minDist)

	if x1 == x2 && y1 == y2 {
		if currDist < *minDist {
			*minDist = currDist
			fmt.Printf("Found path with distance %d\n", currDist)
			*minPath = make([][2]int, len(*currPath))
			copy(*minPath, *currPath)
		}
	} else {
		if currDist < *minDist {
			for _, dir := range Dirs {
				dfs_helper(x1+dir[0], y1+dir[1], x2, y2, currPath, minPath, minDist, currDist+1, visited)
			}
		}
	}

	// Backtracking.
	visited[x1][y1] = false
	*currPath = (*currPath)[:len(*currPath)-1]
}

func dfs(x1 int, y1 int, x2 int, y2 int) ([][2]int, int) {
	var currPath, minPath [][2]int
	minDist := absDiffInt(x2, x1) + absDiffInt(y2, y1) // Manhattan distance as an initial bound

	visited := make([][]bool, 20)
	for i := 0; i < 20; i++ {
		visited[i] = make([]bool, 20)
	}

	dfs_helper(x1, y1, x2, y2, &currPath, &minPath, &minDist, 0, visited)
	return minPath, minDist
}

func handle_dfs(w http.ResponseWriter, r *http.Request) {
	// Set response headers.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var input RequestBody
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Decode error! Please check your JSON formatting.", http.StatusBadRequest)
		return
	}

	path, distance := dfs(input.X1, input.Y1, input.X2, input.Y2)
	json.NewEncoder(w).Encode(Response{MinDist: distance, Path: path})
}