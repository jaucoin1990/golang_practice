package main
import "fmt"
func main() {
	var x = "This string was assigned to a variable using normal declaration."
	fmt.Println(x)

	var y string
	fmt.Println(y)

	z := "This string was assigned to a variable using short declaration."
	fmt.Println(z)
}