package main

import (
	"fmt"
	"time"
)

func main() {
	var age = age(8, 3, 1988)
	var firstName string = "Mike"
	var gigsOfRam int = 32

	// Printf: newline and spacing not provided for you -- manual.
	fmt.Printf("1. Hello, " + firstName + "!\n")

	// Interpolate with %v for any type.
	fmt.Printf("2. Hello, %v!\n", firstName)

	// Sprintf: return string without printing; manual newline and spacing.
	stringFormat := fmt.Sprintf("3. Hello, %v!\n", firstName)
	fmt.Printf(stringFormat)

	fmt.Printf(fmt.Sprintf("4. Hello, %v!\n", firstName))

	fmt.Print("5. "+firstName+" is ", age, " years old and has ", gigsOfRam, " GB of RAM!\n")

	fmt.Println("6. "+firstName+" is", age, "years old and has", gigsOfRam, "GB of RAM!")

	fmt.Println("7.", firstName, "is", age, "years old and has", gigsOfRam, "GB of RAM!")
}

// Return an age for a provided birth month, birth day, and year
func age(birthMonth, birthDay, birthYear int) int {
	if time.Now().Month() >= time.Month(birthMonth) && time.Now().Day() >= birthDay {
		return time.Now().Year() - birthYear
	}

	return (time.Now().Year() - birthYear) - 1
}
