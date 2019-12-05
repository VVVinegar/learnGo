package main

import (
	"fmt"
	"os"
)

func main() {
	maze := readMaze("maze/maze.in")

	step := walk(maze, point{0, 0}, point{len(maze)-1, len(maze[0])-1})

	for _, i := range step {
		for _, j := range i{
			fmt.Printf("%3d", j)
		}
		fmt.Println()
	}
}

func walk(maze [][]int, start point, end point) [][]int {
	step := make([][]int, len(maze))

	for i := range step {
		step[i] = make([]int, len(maze[i]))
	}

	q := []point{start}

	for len(q) > 0 {
		currentPoint := q[0]
		q = q[1:]

		if currentPoint == end {
			break
		}

		for _, dir := range dirs {
			nextPoint := currentPoint.add(dir)

			val, ok := nextPoint.at(maze)
			// 有障碍
			if !ok || val == 1{
				continue
			}

			val, ok = nextPoint.at(step)
			// 走过
			if !ok || val != 0{
				continue
			}

			// 回到起点
			if nextPoint == start {
				continue
			}

			currentStep, _ := currentPoint.at(step)
			step[nextPoint.i][nextPoint.j] = currentStep + 1

			q = append(q, nextPoint)

		}
	}

	return step
}

type point struct {
	i int
	j int
}

func (p point) add(p2 point) point {
	return point{
		p.i + p2.i,
		p.j + p2.j,
	}
}

var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i<0 || p.i > len(grid)-1 {
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

func readMaze(fileName string) [][]int {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	var row, column int

	fmt.Fscanf(file, "%d %d", &row, &column)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, column)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}
