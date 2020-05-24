package gridrobo

import (
	"errors"
)

type Point struct {
	Row, Col int
}

func FindPath(maze [][]bool) []*Point {
	if maze == nil || len(maze) == 0 {
		return nil
	}
	if path, ok := findPath(maze, len(maze)-1, len(maze[0])-1, []*Point{}); ok {
		return path
	}
	return nil
}

func findPath(maze [][]bool, row, col int, path []*Point) ([]*Point, bool) {
	if col < 0 || row < 0 || !maze[row][col] {
		return nil, false
	}
	isAtOrigin := (row == 0) && (col == 0)
	path, leftOk := findPath(maze, row, col-1, path)
	path, aboveOk := findPath(maze, row-1, col, path)
	if isAtOrigin || leftOk || aboveOk {
		return append(path, &Point{row, col}), true
	}
	return nil, false
}

type HashsetMap map[interface{}]struct{}

type Hashset struct {
	set HashsetMap
}

func NewHashset(values ...interface{}) *Hashset {
	h := Hashset{make(HashsetMap)}
	for _, v := range values {
		h.set[v] = struct{}{}
	}
	return &h
}

func (h *Hashset) Add(v interface{}) error {
	if _, ok := h.set[v]; ok {
		return errors.New("Cannot add duplicate values.")
	}

	h.set[v] = struct{}{}
	return nil
}

func (h *Hashset) Contains(v interface{}) bool {
	_, ok := h.set[v]
	return ok
}

func FindPath2(maze [][]bool) []*Point {
	if maze == nil || len(maze) == 0 {
		return nil
	}
	if path, ok := findPath2(maze, len(maze)-1, len(maze[0])-1, []*Point{}, NewHashset()); ok {
		return path
	}
	return nil
}

func findPath2(maze [][]bool, row, col int, path []*Point, failedPoints *Hashset) ([]*Point, bool) {
	if col < 0 || row < 0 || !maze[row][col] {
		return nil, false
	}
	p := &Point{row, col}
	if failedPoints.Contains(p) {
		return nil, false
	}
	isAtOrigin := (row == 0) || (col == 0)
	path, leftOk := findPath2(maze, row, col-1, path, failedPoints)
	path, aboveOk := findPath2(maze, row-1, col, path, failedPoints)
	if isAtOrigin || leftOk || aboveOk {
		return append(path, p), true
	}
	failedPoints.Add(p)
	return nil, false
}
