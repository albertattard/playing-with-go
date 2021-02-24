package main

import (
	"fmt"
	"sort"
)

type pet struct {
	name string
	age  uint
}

func (p pet) String() string {
	return fmt.Sprintf("%v is %v years old", p.name, p.age)
}

type pets []pet
type byAge struct{ pets }
type byName struct{ pets }

func (a pets) Len() int             { return len(a) }
func (a pets) Swap(i, j int)        { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool  { return a.pets[i].age < a.pets[j].age }
func (a byName) Less(i, j int) bool { return a.pets[i].name < a.pets[j].name }

func runSortDemo() {
	pets := pets{
		pet{"Tina", 5},
		pet{"Fluffy", 2},
		pet{"Fido", 7},
	}

	fmt.Printf("Original:                   %v\n", pets)

	sort.Sort(byAge{pets})
	fmt.Printf("Sorted (by age):            %v\n", pets)

	sort.Slice(pets, func(i, j int) bool { return pets[i].age > pets[j].age })
	fmt.Printf("Sorted (by age reversed):   %v\n", pets)

	sort.Sort(byName{pets})
	fmt.Printf("Sorted (by name):           %v\n", pets)

	sort.Slice(pets, func(i, j int) bool { return pets[i].name > pets[j].name })
	fmt.Printf("Sorted (by name reversed):  %v\n", pets)
}
