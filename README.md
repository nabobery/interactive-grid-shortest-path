# Grid Pathfinding Solver

This project is a grid pathfinding solver that demonstrates finding the shortest path on a 20×20 grid using the Breadth-First Search (BFS) algorithm on the backend together with a React frontend.

## Backend

The backend is written in Go and now includes:

- **BFS Handler** (`bfs_handler.go`):  
  Implements the BFS algorithm to compute the shortest path. BFS is used since it guarantees the shortest path in an unweighted grid—with a time complexity of O(n²) for a grid of n cells.

- **DFS Handler** (`dfs_handler.go`):  
  Contains a DFS implementation. However, DFS is not used here because its exponential worst-case time complexity (approximately O(4^(n^2))) makes it impractical for larger grids.

- **Main Server** (`main.go`):  
  Sets up the HTTP server, routing the `/find-path` endpoint to the BFS handler.

### Running the Backend

1. Make sure you have Go installed.
2. Navigate to the `backend` directory.
3. Run the server (ensuring all Go files are included):

   ```bash
   go run main.go
   ```

4. The server will be available at `127.0.0.1:8000`.

## Frontend

The frontend is built with React.

- **Grid.jsx**:  
  Displays a 20×20 grid. Click a cell to set the start point, click another cell for the destination, and then click "Submit" to retrieve the path from the backend. The resulting path is highlighted in blue, and the shortest distance is displayed.

- **Grid.css**:  
  Contains styles for the grid. The `.shortest-path` class colors the path cells blue.

### Running the Frontend

1. Navigate to the `frontend/react-app` directory.
2. Install dependencies:

   ```bash
   npm install
   ```

3. Start the development server:

   ```bash
   npm start
   ```

4. Open `http://localhost:5173` in your browser.

## Algorithm Comparison

- **BFS (Breadth-First Search)**:  
  Runs in O(n²) on a grid of n cells and ensures that the shortest path is found. This makes it very suitable for our 20×20 grid pathfinding.

- **DFS (Depth-First Search)**:  
  Has an exponential worst-case time complexity (approximately O(4^n)) since it might recursively explore many redundant paths. For these reasons, DFS is avoided in favor of BFS in this project.

## Conclusion

This project demonstrates how a performant backend BFS pathfinding solution can be integrated with a modern React frontend to visually display the shortest path and distance on a grid.
