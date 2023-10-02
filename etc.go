package main


func some( a int, b  int) int {
	a = a + b
	b = a - b
	a = a - b
	return a
}