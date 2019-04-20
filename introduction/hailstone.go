package algorithms

var Sequence []int

func Hailstone(n int) int {
	var l = 0
	Sequence = make([]int, 0)
	Sequence = append(Sequence, n)
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		Sequence = append(Sequence, n)
		l++
	}

	return l
}
