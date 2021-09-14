package sqrt

import "math"

func IntSqrtBinarySearch(x int) int {
    min := 0
    max := x / 2 // sqrt(x) <= x / 2 for all x 

    for min <= max {
        mid := min + (max - min)/2
        midSquared := mid*mid
        if midSquared <= x && (mid + 1)*(mid + 1) > x {
            return mid
        }
        if midSquared < x {
            min = mid + 1
        } else if midSquared > x {
            max = mid - 1
        }
    }
    return min
}

func SqrtPocketCalculator(x int) int {
	// sqrt(x) = e ^ ln(sqrt(x)) = e ^ ln(x^1/2) = e ^ (1/2 * ln(x))
	return int(math.Pow(math.E, math.Log(float64(x)) / 2))
}

func SqrtRecursiveBitShifts(x int) int {
	// x = 2*y + r, where r is 0 or 1
	// sqrt(x) = 2 * sqrt(x / 2^2) until base case which is either 1 or 2
	if x < 2 {
		return x
	}
	leftSqrt := SqrtRecursiveBitShifts(x >> 2) << 1
	rightSqrt := leftSqrt + 1
	if rightSqrt*rightSqrt > x {
		return leftSqrt
	}
	return rightSqrt
}


// also called Babylonian or Heron's method
func SqrtNewtonsMethod(x int) int {
	if x < 2 {
		return x
	}
	precision := 0.1
	X := float64(x)
	z := X
	for z*z >= X + precision {
		z = (z + X / z) / 2
	}
	
	return int(z)
}
