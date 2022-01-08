package basic

const (
	i = iota + 1
	j
	k
)

const (
	one = 1 << iota // 1
	two // 2
	three // 4
	four = "four" // "four"
	five // "four"
	six = 1 << iota // 32
)
