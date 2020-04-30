package testmod

import "fmt"

// Print prints the formatted salutation message
func Print(salutation, name string) {
	fmt.Printf("%s, %s", salutation, name)
}
