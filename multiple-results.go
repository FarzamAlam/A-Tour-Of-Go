package main

func main() {
	a, b := swap("Hello", "world")
	print(a, " ", b)
}

func swap(a, b string) (string, string) {
	return b, a
}
