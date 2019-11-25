package main

<<<<<<< HEAD
import (
	"fmt"
   "runtime"
	"os"
)

func main() {
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}
=======
var a = "G"

func main() {
	n()
	m()
	n()
}

func n() { print(a) }

func m() {
	a := "O"
	print(a)
}
>>>>>>> 12b184ea8c350dba2a34b68fa820e19542c93c81
