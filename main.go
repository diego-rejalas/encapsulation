package main

import "fmt"

func main() {
	Go := Course{
		"Go desde 0",
		12.64,
		false,
		[]uint{12, 34, 67},
		map[uint]string{
			1: "Introduccion",
			2: "Estructuras",
			3: "Maps",
		},
	}

	// css := Course{Name: "Css from scratch", IsFree: false}

	// Js := Course{}
	// Js.Name = "Course Js"
	// Js.IsFree = false

	// fmt.Println(Go.Classes)
	// fmt.Printf("%+v:\n", css)
	// fmt.Printf("%+v:\n", Js)

	// Go.PrintClasses()
	Go.ChangePrice(23.4)
	fmt.Println(Go.Price)
}
