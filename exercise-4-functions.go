package main

// pingala is a function that returns
// a function that returns an int.
func pingala() func() int {
	count := 0
	num1 := 0
	num2 := 1
	return func() int {
		if count == 0 {
			count++
			return 0
		} else if count == 1 {
			count++
			return 1
		}
		next := num1 + num2
		num1 = num2
		num2 = next
		return next
	}
}
