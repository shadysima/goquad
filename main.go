package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var a, b, c float64

	//parsing command line args (strings) as float64 values
	//also checks to see if correct number of args is passed to file

	if len(os.Args) == 4 {
		if floatVal, err := strconv.ParseFloat(os.Args[1], 64); err == nil {
			a = floatVal
		}
		if floatVal, err := strconv.ParseFloat(os.Args[2], 64); err == nil {
			b = floatVal
		}
		if floatVal, err := strconv.ParseFloat(os.Args[3], 64); err == nil {
			c = floatVal
		}
	} else if len(os.Args) == 1 {
		DisplayDoc()
		return
	} else {
		fmt.Println("Error. Incorrect number of arguments passed. Exiting.")
		return
	}

	//quadratic calculations

	var disc float64 = math.Pow(b, 2) - (4 * a * c)
	var posDisc float64 = math.Sqrt(math.Abs(disc))
	var root1, root2 float64

	if disc < 0 {
		fmt.Printf("Root 1: %0.3f + %0.3fi\n", (-b / (2 * a)), posDisc)
		fmt.Printf("Root 2: %0.3f - %0.3fi\n", (-b / (2 * a)), posDisc)
		return
	} else if disc == 0 {
		fmt.Printf("Root 1: %0.3f\n", (-b / (2 * a)))
		fmt.Printf("Root 2: %0.3f\n", (-b / (2 * a)))
		return
	} else {
		root1 = (-b + math.Sqrt(disc)) / (2 * a)
		root2 = (-b - math.Sqrt(disc)) / (2 * a)
		fmt.Printf("Root 1: %0.3f\n", root1)
		fmt.Printf("Root 2: %0.3f\n", root2)
		return
	}
}
