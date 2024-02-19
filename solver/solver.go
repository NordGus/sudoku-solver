package solver

import (
	"sudoku-solver/board"
	"sync"
)

const (
	buffSize = 5
)

// Solver is a simple wrapper for the board and call the solving algorithm.
type Solver struct {
	board board.Board
}

// New initializes a new Solver for the given board.
func New(board board.Board) *Solver {
	return &Solver{
		board: board,
	}
}

// Solve starts recursion to solve the given board and works as an orchester
// for the concurrent solving process.
func (s Solver) Solve() []board.Board {
	var (
		solutions = make(chan board.Board, buffSize)
		output    = make([]board.Board, 0, buffSize)
		wg        = new(sync.WaitGroup)
	)

	wg.Add(1) // increase by 1 the amount of goroutines to wait for

	// start solving the board
	go solve(wg, solutions, s.board, 0, 0)

	// simple concurrency control pattern for this kind of concurrency implementation
	go func(wg *sync.WaitGroup, solutions chan board.Board) {
		wg.Wait()        // wait for all goroutines to finish
		close(solutions) // close the solutions channel
	}(wg, solutions)

	for solution := range solutions {
		output = append(output, solution)
	}

	return output
}

// solves fills the board recursively until its full or all possible decision
// branches are exhausted.
func solve(wg *sync.WaitGroup, solution chan<- board.Board, current board.Board, x int, y int) {
	defer wg.Done()

	// if y reaches the board's height it means the board is solved
	if y == board.Height {
		solution <- current
		return
	}

	var (
		nextX = x + 1
		nextY = y
	)

	// if nextX reaches the board's width reset it and move to the next row
	if nextX == board.Width {
		nextX = 0
		nextY = y + 1
	}

	// if the current cell is already filled move to the next cell
	if current[y][x] != board.EmptyCell {
		wg.Add(1)
		go solve(wg, solution, current, nextX, nextY)
		return
	}

	candidates := current.PossibleValuesFor(x, y)

	wg.Add(len(candidates))

	for i := 0; i < len(candidates); i++ {
		next := current.Copy()
		next[y][x] = candidates[i]

		go solve(wg, solution, next, nextX, nextY)
	}
}
