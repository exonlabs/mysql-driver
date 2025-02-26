package mysql

import "cmp"

// From Go v1.24
// The min built-in function returns the smallest value of a fixed number of
// arguments of [cmp.Ordered] types. There must be at least one argument.
// If T is a floating-point type and any of the arguments are NaNs,
// min will return NaN.
func min[T cmp.Ordered](x T, y ...T) T {
	r := x
	for _, v := range y {
		if v < r {
			r = v
		}
	}
	return r
}
