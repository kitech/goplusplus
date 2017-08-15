package gopp

func Clamp(x int, low int, high int) int {
	if x >= low && x <= high {
		return x
	} else if x < low {
		return low
	} else {
		return high
	}
}

// C++17 P0025R0 clamp 函数
