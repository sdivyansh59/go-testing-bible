package calculator

import "math"

// CalculateIaArmstrong takes in a 3 digit number 'n'
// and return true if it is an Armstring number
// Armstrong number example 371 == 3^3 + 7^3 + 1^3
func CalculateIsArmstrong(n int) bool {
	a := n /100
	b := n % 100 /10
	c := n %10

	return n == int(math.Pow(float64(a),3)+ math.Pow(float64(b),3)+ math.Pow(float64(c),3))

}


// RandomFunction
func RandomFunction(n int) bool {
	if n> 10 {
		return true
	} else {
		return false
	}
}

//1. To check code coverage :- go test ./... -cover
//2. go test ./... -coverprofile=coverage.out  
// go tool cover -html=coverage.out
