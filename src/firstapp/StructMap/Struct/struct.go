package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number int
	actorName string
	companions []string
	episodes []string
}

type Animal struct {
	Name string
	Origin string
}

type TagAnimal struct {
	Name string `required max:"100"`
	Origin string
}

// Embedding is an alternative to inheritance in go
// Here we are embedding Animal struct inside Bird Struct
type Bird struct {
	Animal
	SpeedKPH float32
	CanFly bool
}

// This syntax is more robust and change proof as it allows us to change the structure of a struct without affecting its underlying functionality
func main() {
	aDoctor := Doctor{
		number: 3,
		actorName: "John Pertwee",
		companions: []string {
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	// Initializing a struct using positional syntax
	// Not recommended to use this syntax due to its inability to accommodate changes made in struct variables in future
	// aDoctor2 := Doctor{
	// 	3,
	// 	"John Pertwee",
	// 	[]string {
	// 		"Liz Shaw",
	// 		"Jo Grant",
	// 		"Sarah Jane Smith",
	// 	},
	// }
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[1])
	// fmt.Println(aDoctor2)

	// Defining an anonymous struct without it's name	
	aDoc := struct{name string}{name: "John Pertwee"}
	anotherDoc := aDoc
	anotherDoc.name = "Tom Baker"
	// Have to use pointer syntax to reference to the same parent instead of creating a copy
	anoDoc := &aDoc
	anoDoc.name = "Sherlock"
	// Unlike maps and slices, structs refer to individual datasets when assigned to a different variable
	fmt.Println(aDoc)
	fmt.Println(anotherDoc)
	fmt.Println(anoDoc)

	b := Bird{
		Animal : Animal{Name: "Emu",Origin: "Austraila"},
		SpeedKPH: 30,
		CanFly: false,
	}

	// b:=Bird{}
	// b.Name = "Emu"
	// b.Origin = "Australia"
	// b.SpeedKPH = 77
	// b.CanFly = false
	fmt.Println(b)
	fmt.Println(b.Name)
	t := reflect.TypeOf(TagAnimal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)

}