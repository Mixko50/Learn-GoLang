package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

type Animal struct {
	name   string `required max : "100"`
	origin string
}

type Bird struct {
	Animal
	speedKPH int
	canFly   bool
}

func main() {

	fmt.Println("----- Maps -----")
	statePopu := make(map[string]int)
	statePopu = map[string]int{
		"Cali": 345,
		"Tex":  100,
		"TO":   23,
		"New":  1234,
		"Ohio": 213,
	}

	fmt.Println(statePopu)
	fmt.Println(statePopu["Tex"])
	statePopu["Cali"] = 1000
	fmt.Println(statePopu["Cali"])
	delete(statePopu, "New")
	fmt.Println(statePopu)

	_, ok := statePopu["Oho"]
	fmt.Println(ok) //ok -> check the key in the map or not and return boolean.
	fmt.Println(len(statePopu))

	sp := statePopu
	delete(sp, "Ohio")
	fmt.Println(statePopu)
	fmt.Println(sp) //Point to statePopu

	fmt.Println("----- Structs -----")
	aDoctor := Doctor{
		3,
		"Mixko",
		[]string{
			"Apisit",
			"BB",
			"SSSS",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.companions)
	fmt.Println(aDoctor.companions[1])

	bDoctor := struct{ name string }{name: "Mixko50"}
	fmt.Println(bDoctor)
	// cDoctor := &bDoctor
	cDoctor := bDoctor
	cDoctor.name = "Apisit"
	fmt.Println(bDoctor)
	fmt.Println(cDoctor)

	nok := Bird{
		Animal:   Animal{name: "Emu", origin: "Thai"},
		speedKPH: 12,
		canFly:   true,
	}

	fmt.Println(nok)

	tag := reflect.TypeOf(Animal{})
	field, _ := tag.FieldByName("name")
	fmt.Println(field.Tag)

}
