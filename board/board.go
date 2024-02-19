package board

// Board is a data structure that represents a 9x9 sudoku board (with y major).
//
// Empty cells are stored as zeros, following sudoku's rule that cell must be
// filled with numbers from 1 to 9.
type Board [9][9]uint8

// New initializes a new Board using the passed cells.
//
// Remember that empty cells must use zeros.
func New(cells [9][9]uint8) Board {
	return cells
}

// PossibleValuesFor finds the values that can fit in Board for the given cell
// [x,y]
func (board Board) PossibleValuesFor(x int, y int) []uint8 {
	values := make([]uint8, 0, 9)

	if board[y][x] > 0 {
		return values
	}

	for i := uint8(1); i < 10; i++ {
		if !board.isHorizontallyValid(i, x, y) {
			continue
		}

		if !board.isVerticallyValid(i, x, y) {
			continue
		}

		if !board.isSectorValid(i, x, y) {
			continue
		}

		values = append(values, i)
	}

	return values
}

// isHorizontallyValid checks that the given value doesn't violate sudoku's
// horizontal rule.
func (board Board) isHorizontallyValid(value uint8, x int, y int) bool {
	for i := 0; i < 9; i++ {
		if i == x {
			continue
		}

		if board[y][i] == value {
			return false
		}
	}

	return true
}

// isVerticallyValid checks that the given value doesn't violate sudoku's
// vertical rule.
func (board Board) isVerticallyValid(value uint8, x, y int) bool {
	for i := 0; i < 9; i++ {
		if i == y {
			continue
		}

		if board[i][x] == value {
			return false
		}
	}

	return true
}

// isSectorValid checks that the given value doesn't violate sudoku's vertical
// rule.
func (board Board) isSectorValid(value uint8, x int, y int) bool {
	var (
		minX, maxX = getAxisLimits(x)
		minY, maxY = getAxisLimits(y)
	)

	for i := minX; i < maxX; i++ {
		for j := minY; j < maxY; j++ {
			if i == x && j == y {
				continue
			}

			if board[j][i] == value {
				return false
			}
		}
	}

	return true
}

// Copy returns a new copy of the board
func (board Board) Copy() Board {
	cp := [9][9]uint8{}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			cp[i][j] = board[i][j]
		}
	}

	return cp
}
