package controllers

import (
	"encoding/json"
	"net/http"
)

// bfs implements the Breadth-First Search algorithm to find the shortest path
// in a 20×20 grid from (x1, y1) to (x2, y2).
func bfs(x1 int, y1 int, x2 int, y2 int) ([][2]int, int) {
	queue := make([][2]int, 0)
	visited := make([][]bool, 20)
	parent := make([][]*Point, 20)
	for i := 0; i < 20; i++ {
		visited[i] = make([]bool, 20)
		parent[i] = make([]*Point, 20)
	}

	queue = append(queue, [2]int{x1, y1})
	visited[x1][y1] = true

	found := false
	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		if cell[0] == x2 && cell[1] == y2 {
			found = true
			break
		}
		for _, d := range Dirs {
			nx := cell[0] + d[0]
			ny := cell[1] + d[1]
			if nx >= 0 && nx < 20 && ny >= 0 && ny < 20 && !visited[nx][ny] {
				visited[nx][ny] = true
				temp := new(Point)
				temp.x = cell[0]
				temp.y = cell[1]
				parent[nx][ny] = temp
				queue = append(queue, [2]int{nx, ny})
			}
		}
	}

	if !found {
		return nil, 0 // Should not happen on an obstacle-free grid.
	}

	// Reconstruct the path from destination to start.
	var path [][2]int
	cur := [2]int{x2, y2}
	for {
		path = append([][2]int{cur}, path...)
		if cur[0] == x1 && cur[1] == y1 {
			break
		}
		pt := parent[cur[0]][cur[1]]
		if pt == nil {
			break
		}
		cur = [2]int{pt.x, pt.y}
	}
	distance := len(path) - 1
	return path, distance
}

func BFSHandler(w http.ResponseWriter, r *http.Request) {
	// Set response headers.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var input RequestBody
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Decode error! Please check your JSON formatting.", http.StatusBadRequest)
		return
	}

	path, distance := bfs(input.X1, input.Y1, input.X2, input.Y2)
	json.NewEncoder(w).Encode(Response{MinDist: distance, Path: path})
}
