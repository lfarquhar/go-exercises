package main

import "fmt"

/* type Direction int

const (
	North Direction = iota
	East
	South
	West
) */

type emailStatus int

const (
	emailBounced emailStatus = iota
	emailInvalid
	emailDelivered
	emailOpened
)

func main() {
	fmt.Println(emailBounced, emailInvalid, emailDelivered, emailOpened)

	var X = emailStatus(8)

	fmt.Println(X)
}
