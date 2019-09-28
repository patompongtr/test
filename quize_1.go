package main

import "fmt"

func main() {
	fruits := []string{"Apple", "Orange", "Plum", "Promenganate", "grape"}
	fmt.Println("fruits :", fruits)
	//some := make[]string
	some := fruits[1:3:3]

	some = append(some, "Banana")
	fmt.Println("fruits :", fruits)
	fmt.Println("some :", some)

}
