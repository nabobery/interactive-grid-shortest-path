// import React, { useState, useEffect } from "react";
import React, { useState} from "react";
import "./Grid.css";

const Grid = () => {
  const gridSize = 20;
  const totalCells = gridSize * gridSize;
  const [start, setStart] = useState(null);
  const [dest, setDest] = useState(null);
  const [result, setResult] = useState(null);

  // Handle cell click: first click sets the start, second sets destination; further clicks reset.
  const handleCellClick = (index) => {
    if (start === null) {
      setStart(index);
      console.log("Start cell: ", index);
    } else if (dest === null && index !== start) {
      setDest(index);
      console.log("Destination cell: ", index);
    } else {
      // Reset selections and clear previous results.
      setStart(index);
      setDest(null);
      setResult(null);
    }
  };

  // Submit the start and destination to the backend and update result with the BFS response.
  const handleSubmit = () => {
    if (start !== null && dest !== null) {
      let payload = {
        x1: Math.floor(start / gridSize),
        y1: start % gridSize,
        x2: Math.floor(dest / gridSize),
        y2: dest % gridSize,
      };
      fetch("http://localhost:8000/find-path", {
        method: "POST",
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
        },
        body: JSON.stringify(payload),
      })
        .then((response) => response.json())
        .then((data) => {
          console.log("BFS Result Received: ", data);
          setResult(data);
        })
        .catch((error) => {
          console.error("Error fetching path: ", error);
        });
    }
  };

  // Check if a grid cell (given its row and col) is part of the returned BFS path.
  const isCellInPath = (row, col) => {
    console.log("Cell in path ran for row, col: ", row, col)
    if (result && result.path && result.path.length > 0) {
      return result.path.some((coord) => coord[0] === row && coord[1] === col);
    }
    return false;
  };

  // useEffect(() => {
  //   console.log(result)
  // }, [result])

  return (
    <div>
      <div className="grid">
        {Array.from({ length: totalCells }).map((_, index) => {
          const row = Math.floor(index / gridSize);
          const col = index % gridSize;
          return (
            <div
              key={index}
              className={`cell ${index === start ? "start" : ""} ${
                index === dest ? "dest" : ""
              } ${isCellInPath(row, col) ? "shortest-path" : ""}`}
              onClick={() => handleCellClick(index)}
            >
              {`${row}, ${col}`}
            </div>
          );
        })}
      </div>
      <button onClick={handleSubmit}>Submit</button>
      {result && (
        <div className="result">
          <p>Shortest Distance: {result.distance}</p>
        </div>
      )}
    </div>
  );
};

export default Grid;