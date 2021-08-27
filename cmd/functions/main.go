/* MIT License

Copyright (c) 2021 Timothy Norman Murphy <tnmurphy@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

package main

import (
	"fmt"
	"strings"
)


type Person struct {
	Name string
	Age  int
}

var people = []Person{
		{"Alice", 25},
		{"Bob",   24},
		{"Jane",  23},
		{"Fred",  25},
		{"John",  26}, 
}



func isOldEnough(p *Person) bool {
	return p.Age > 20
}



func selectPeople(filter func(p *Person)bool) []Person {
	// allocate a dynamic array for the result
	selectedPeople := make([]Person, 0, len(people)) 
    // loop over the results
	for _, p := range(people) {
		if filter(&p) {
			selectedPeople = append(selectedPeople, p)
		}
	}

	return selectedPeople
}





func peopleGenerator(filter func(p *Person)bool) func()(Person, bool) {
    // Returns a function which "gets the next" person
    // It demonstrates a continuation where the returned function remembers
    // "currentPersonIndex" so that it is able to issue the next record
    // Note that this avoids allocating a dynamic array which could save
    // memory when there are a large number of records.
	currentPersonIndex := -1
	return func()(Person, bool) {
		currentPersonIndex++
		if currentPersonIndex >= len(people) {
			return Person{}, false
		}
		return people[currentPersonIndex], true
	}
}

func main() {
	fmt.Println("lesson 4 - functions")

	// Function Calls
	if isOldEnough(&people[0]) {
		fmt.Printf("True returned from isOldEnough\n")
	}


	fmt.Printf("\nAnonymous Function passed as an argument:\n")
	for _, person := range(selectPeople(
		func(p *Person)bool {
			return p.Age > 20 
		})) {
		fmt.Printf("Selected %q\n", person.Name)	
	}

	// You can assign anonymous functions to variables
	var myFilter = func(p *Person) bool {
		return strings.ContainsAny(p.Name, "Aa") 
	}

	fmt.Printf("\nAnonymous function assigned to a variable (people with names containing 'a'):\n")
	for _, person := range(selectPeople(myFilter)) {
		fmt.Printf("Person: %v, Age: %d\n", person.Name, person.Age)
	}

	fmt.Printf("\nA continuation\n")
	getNextPerson := peopleGenerator(isOldEnough); 
	for p, ok := getNextPerson(); ok; p, ok = getNextPerson() {
		fmt.Printf("getNextPerson returned: %q\n", p.Name)	
	}

}