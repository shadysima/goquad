package main

import (
	"fmt"
)

func DisplayDoc() {
	fmt.Println("\nNAME")
	fmt.Println("\tgoquad - quadratic solver")
	fmt.Println("USAGE")
	fmt.Println("\tgoquad [ARGS]...")
	fmt.Println("DISCRIPTION")
	fmt.Println("\tA quadratic solver written in Go. \n\t[ARGS] are represented in order of the variables in the quadratic equation (ax^2 + bx + c). \n\tFirst [ARGS] passed is 'a', second [ARGS] passed is 'b', and third [ARGS] passed is 'c'.")
	fmt.Println("\n\tno args")
	fmt.Println("\t\tdisplays this manual page")
	fmt.Println("\t3 args")
	fmt.Println("\t\tcalculates roots using quadratic formula based on [ARGS] passed")
	fmt.Println("\n\tAny other number of [ARGS] will cause an error and exit the solver\n")
}
