package controllers

// dirs defines the 4 possible directions (up, right, down, left).
var Dirs [4][2]int = [4][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

// request_body defines the expected JSON input.
type RequestBody struct {
	X1 int `json:"x1"` // Starting X coordinate
	Y1 int `json:"y1"` // Starting Y coordinate
	X2 int `json:"x2"` // Destination X coordinate
	Y2 int `json:"y2"` // Destination Y coordinate
}

// response defines the JSON response returned by the server.
type Response struct {
	MinDist int      `json:"distance"`
	Path    [][2]int `json:"path"`
}

// Point is used to keep track of the parent for each cell.
type Point struct {
	x int
	y int
}