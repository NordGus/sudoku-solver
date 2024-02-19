package board

// getAxisLimits is a helper function that helps determine the Board sector
// boundaries for any axis. When n is the X position it will return the limits
// for the X axis and vice versa.
func getAxisLimits(n int) (bottom, upper int) {
	bottom, upper = 6, 9

	if n < 6 {
		bottom, upper = 3, 6
	}

	if n < 3 {
		bottom, upper = 0, 3
	}

	return
}
