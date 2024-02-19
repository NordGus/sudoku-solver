package board

import "testing"

func Test_getAxisLimits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name       string
		args       args
		wantBottom int
		wantUpper  int
	}{
		{
			name:       "first quadrant",
			args:       args{n: 2},
			wantBottom: 0,
			wantUpper:  3,
		},
		{
			name:       "second quadrant",
			args:       args{n: 5},
			wantBottom: 3,
			wantUpper:  6,
		},
		{
			name:       "third quadrant",
			args:       args{n: 8},
			wantBottom: 6,
			wantUpper:  9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBottom, gotUpper := getAxisLimits(tt.args.n)
			if gotBottom != tt.wantBottom {
				t.Errorf("getAxisLimits() gotBottom = %v, want %v", gotBottom, tt.wantBottom)
			}
			if gotUpper != tt.wantUpper {
				t.Errorf("getAxisLimits() gotUpper = %v, want %v", gotUpper, tt.wantUpper)
			}
		})
	}
}
