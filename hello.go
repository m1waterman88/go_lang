package main

import "fmt"

func main() {
	// Try out variables and string interpolation with variations.
	var age, ram int = 2020 - 1988, 32
	name := "Mike"

	// Printf: newline and spacing not provided for you -- manual.
	fmt.Printf("1. Hello, " + name + "!\n")

	// Interpolate with %v for any type.
	fmt.Printf("2. Hello, %v!\n", name)

	// Sprintf: return string without printing; manual newline and spacing.
	s1 := fmt.Sprintf("3. Hello, %v!\n", name)
	fmt.Printf(s1)

	fmt.Printf(fmt.Sprintf("4. Hello, %v!\n", name))

	fmt.Print("5. "+name+" is ", age, " years old and has ", ram, " GB of RAM!\n")

	fmt.Println("6. "+name+" is", age, "years old and has", ram, "GB of RAM!")

	fmt.Println("7.", name, "is", age, "years old and has", ram, "GB of RAM!")
}
