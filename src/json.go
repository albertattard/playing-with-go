package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type person struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     uint   `json:"age"`
}

type persons []person

func (person person) String() string {
	return fmt.Sprintf("%v %v is %v years old", person.Name, strings.ToUpper(person.Surname), person.Age)
}

func (person person) print() {
	fmt.Printf("%v\n", person)
}

func (persons persons) printAll() {
	for _, person := range persons {
		person.print()
	}
}

func runJsonMarshalDemo() {
	albert := person{
		Name:    "Albert",
		Surname: "Attard",
		Age:     42,
	}

	source := persons{albert}
	fmt.Println("Source")
	source.printAll()

	bytes, e := json.Marshal(source)
	if e != nil {
		fmt.Printf("Failed to marshal data to JSON - %v\n", e)
		return
	}

	fmt.Printf("JSON %v\n", string(bytes))

	var unmarshalled persons
	e = json.Unmarshal(bytes, &unmarshalled)
	if e != nil {
		fmt.Printf("Failed to unmarshal JSON to data - %v\n", e)
		return
	}

	fmt.Println("Unmarshalled")
	unmarshalled.printAll()
}

func runJsonEncoderDemo() {
	albert := person{
		Name:    "Albert",
		Surname: "Attard",
		Age:     42,
	}

	source := persons{albert}
	fmt.Println("Source")
	source.printAll()

	e := json.NewEncoder(os.Stdout).Encode(source)
	if e != nil {
		fmt.Printf("Failed to encode data to JSON - %v\n", e)
		return
	}

	var decoded persons
	e = json.NewDecoder(strings.NewReader(`[{"name":"Albert","surname":"Attard","age":42}]`)).Decode(&decoded)
	if e != nil {
		fmt.Printf("Failed to decoded JSON to data - %v\n", e)
		return
	}

	fmt.Println("Decoded")
	decoded.printAll()
}
