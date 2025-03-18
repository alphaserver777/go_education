package main

func main() {
	GetSum(-5, 3)

}

func GetSum(a int, b int) int {
	if a == b {
		return a
	}

	min := a
	max := b
	if a > b {
		min = b
		max = a
	}

	sum := 0
	for i := min; i <= max; i++ {
		sum += i
	}

	return sum
}
