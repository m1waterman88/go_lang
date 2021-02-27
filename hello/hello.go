package main

import (
	"fmt"
	"time"
)

func main() {
	// Try out variable variations.
	age := age(8, 3)

	//var age, gigsOfRam int = time.Now().Year() - 1988, 32
	var gigsOfRam int = 32
	var firstName string = "Mike"

	// Printf: newline and spacing not provided for you -- manual.
	fmt.Printf("1. Hello, " + firstName + "!\n")

	// Interpolate with %v for any type.
	fmt.Printf("2. Hello, %v!\n", firstName)

	// Sprintf: return string without printing; manual newline and spacing.
	s1 := fmt.Sprintf("3. Hello, %v!\n", firstName)
	fmt.Printf(s1)

	fmt.Printf(fmt.Sprintf("4. Hello, %v!\n", firstName))

	fmt.Print("5. "+firstName+" is ", age, " years old and has ", gigsOfRam, " GB of RAM!\n")

	fmt.Println("6. "+firstName+" is", age, "years old and has", gigsOfRam, "GB of RAM!")

	fmt.Println("7.", firstName, "is", age, "years old and has", gigsOfRam, "GB of RAM!")
}

func age(birthMonth int, birthDay int) int {
	if time.Now().Month() >= time.Month(birthMonth) && time.Now().Day() >= birthDay {
		return time.Now().Year() - 1988
	}

	return (time.Now().Year() - 1988) - 1
}
