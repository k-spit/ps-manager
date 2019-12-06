package utils

//Even checks if number is even
func Even(number int) bool {
	return number%2 == 0
}

//Odd checks if number is odd
func Odd(number int) bool {
	return !Even(number)
}
