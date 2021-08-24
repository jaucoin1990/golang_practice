package main
import "fmt"
import "reflect"

var x = 42 //DECLARED the VARIABLE with IDENTIFIER of "x" and VALUE of 42
var y = "Example string" //DECLARED the VARIABLE with IDENTIFIER y and VALUE "Example string"
var z = true //DECLARED the VARIABLE with IDENTIFIER z and VALUE true

func main() {
	fmt.Println(x)
	fmt.Printf("This type: %T\n", x)
	fmt.Println(y)
	fmt.Println("I'm showing this TYPE by using reflect.", "TYPE:", reflect.TypeOf(y))
	fmt.Println(z)
	fmt.Printf("This type: %T\n", z)
}