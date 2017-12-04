package stdext

// The standard library only offers typical math functions for float64
// arguments. Presumably because doing otherwise would result in an explosion
// in the number of functions needed to cover all the built-in numeric types.
// TODO(DH): Do this in p-m-g?

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Min64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Max64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Abs64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
