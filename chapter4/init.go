package trans

import "math"

// Pi is a float64 number
var Pi float64

func init() {
	Pi = 4 * math.Atan(1) // init() function computes Pi
}
