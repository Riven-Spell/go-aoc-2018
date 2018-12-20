package utils

func Range(start, end int) []int {
	out := make([]int, end - start)
	for k := range out {
		out[k] = k + start
	}
	return out
}