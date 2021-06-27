// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type and create a value of this type. Declare a function
// that can change the value of some field in this struct type. Display the
// value before and after the call to your function.
package main

import "fmt"

// Add imports.

// Declare a type named user.
type user struct {
	Name string
	Age  int
}

// Create a function that changes the value of one of the user fields.
func funcName( /* add pointer parameter, add value parameter */ user *user, name string) {

	// Use the pointer to change the value that the
	// pointer points to.
	user.Name = name
}

func main() {

	// Create a variable of type user and initialize each field.
	jack := user{
		Name: "Jack",
		Age:  24,
	}

	// Display the value of the variable.
	fmt.Println(jack.Name, jack.Age)

	// Share the variable with the function you declared above.
	funcName(&jack, "Jackson")

	// Display the value of the variable.
	fmt.Println(jack.Name, jack.Age)
}
