package board

import (
	"reflect"
	"testing"
)

var (
	baseBoard = New([9][9]uint8{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})
)

func TestBoard_PossibleValuesFor(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name  string
		board Board
		args  args
		want  []uint8
	}{
		{
			name:  "return a list of possible values for the position",
			board: baseBoard,
			args:  args{x: 2, y: 0},
			want:  []uint8{1, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.board.PossibleValuesFor(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PossibleValuesFor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_Copy(t *testing.T) {
	tests := []struct {
		name  string
		board Board
		want  Board
	}{
		{
			name:  "returns the expected board",
			board: baseBoard,
			want: New([9][9]uint8{
				{5, 3, 0, 0, 7, 0, 0, 0, 0},
				{6, 0, 0, 1, 9, 5, 0, 0, 0},
				{0, 9, 8, 0, 0, 0, 0, 6, 0},
				{8, 0, 0, 0, 6, 0, 0, 0, 3},
				{4, 0, 0, 8, 0, 3, 0, 0, 1},
				{7, 0, 0, 0, 2, 0, 0, 0, 6},
				{0, 0, 0, 4, 1, 9, 0, 0, 5},
				{0, 0, 0, 0, 8, 0, 0, 7, 9},
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.board.Copy()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Copy() = %v, want %v", got, tt.want)
			}

			if &got[0:cap(got)][cap(got)-1] == &tt.board[0:cap(tt.board)][cap(tt.board)-1] {
				t.Errorf("Copy() returned the same instance")
			}
		})
	}
}

func TestBoard_isHorizontallyValid(t *testing.T) {
	type args struct {
		value uint8
		x     int
		y     int
	}
	tests := []struct {
		name  string
		board Board
		args  args
		want  bool
	}{
		{
			name:  "is valid",
			board: baseBoard,
			args: args{
				value: 9,
				x:     2,
				y:     0,
			},
			want: true,
		},
		{
			name:  "is not valid",
			board: baseBoard,
			args: args{
				value: 7,
				x:     2,
				y:     0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.board.isHorizontallyValid(tt.args.value, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("isHorizontallyValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_isVerticallyValid(t *testing.T) {
	type args struct {
		value uint8
		x     int
		y     int
	}
	tests := []struct {
		name  string
		board Board
		args  args
		want  bool
	}{
		{
			name:  "is valid",
			board: baseBoard,
			args: args{
				value: 9,
				x:     2,
				y:     0,
			},
			want: true,
		},
		{
			name:  "is not valid",
			board: baseBoard,
			args: args{
				value: 8,
				x:     2,
				y:     0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.board.isVerticallyValid(tt.args.value, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("isVerticallyValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
