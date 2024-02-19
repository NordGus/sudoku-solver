package solver

import (
	"sudoku-solver/board"
	"sync"
)

type Solver struct {
	board board.Board
}

func New(board board.Board) *Solver {
	return &Solver{
		board: board,
	}
}

func (s Solver) Solve() []board.Board {
	var (
		solutions = make(chan board.Board, 3)
		output    = make([]board.Board, 0, 5)
		wg        = new(sync.WaitGroup)
	)

	wg.Add(1)

	go solve(wg, solutions, s.board, 0, 0)

	go func(wg *sync.WaitGroup, solutions chan board.Board) {
		wg.Wait()
		close(solutions)
	}(wg, solutions)

	for solution := range solutions {
		output = append(output, solution)
	}

	return output
}

func solve(wg *sync.WaitGroup, solution chan<- board.Board, current board.Board, x int, y int) {
	defer wg.Done()

	if y == 9 {
		solution <- current
		return
	}

	var (
		nextX = x + 1
		nextY = y
	)

	if nextX == 9 {
		nextX = 0
		nextY = y + 1
	}

	if current[y][x] > 0 {
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
