package main

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

func main() {
	Go := Course{
		"Go desde Cero",
		12.34,
		false,
		[]uint{12, 56, 89},
		map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Maps",
		},
	}

	css := Course{Name: "CSS desde cero", IsFree: true}
	js := Course{}
	js.Name = "Curso JS"
	js.UserIDs = []uint{12, 67}

	fmt.Println(Go.Name)
	fmt.Printf("%+v\n", css)
	fmt.Printf("%+v", js)
}
