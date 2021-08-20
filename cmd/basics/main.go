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

// lesson 1 - types and assignment.

package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	// LESSON 1 - Variable Declaration/assignment /////////////////////////////

	var name string
	name = "Alice"
	var othername = "Bob"
	yetanothername := "Dave"

	// Variables have to be used - compilation will fail if they are unused:
	fmt.Printf("name=%q, othername=%q, yetanothername=%q\n", name, othername,
		yetanothername)

	// Types
	var isValid bool
	var number int
	var number32 int32
	var unumber64 uint64

	isValid = false // What false and true look like
	isValid = true

	fmt.Printf("isValid=%v, number=%2d, number32=%2d, unumber64=%2d\n", isValid,
		number, number32, unumber64)

	var bob = Person{Name: othername, Age: 24}
	fmt.Printf("Person named %v with age %2d\n", bob.Name, bob.Age)

	// Variables get initialised to their "null" value by default
	var nobody = Person{}
	fmt.Printf("Nobody = %v\n", nobody)

	var personPointer = &bob
	fmt.Printf("pointed-to-person is %q\n", personPointer)
	fmt.Printf("dereferenced pointed-to-person is %q\n", *personPointer)

	// Pointers can be nil but dereferencing one will cause a panic.
	personPointer = nil
	fmt.Printf("nil pointers can be printed: %q\n", personPointer)

	// functions are types too, but more about this later
	var someFunction = func() int {
		return 1
	}

	// Arrays are indicated by starting with a [n] where n is the size
	var team = [2]Person{bob, {Name: "Alice", Age: 25}}
	fmt.Printf("The first member of the team is %v\n", team[0].Name)

	// Ellipses may be used to avoid specifying a length when the contents are
	// being assigned.
	var otherTeam = [...]Person{
		{"Jane", 23},
		{Name: "Fred", Age: 25},
		{"John", 26}, // notice the comma at the end
	}
	fmt.Printf("The other team is: %v\n", otherTeam)

	// In practise we rarely use arrays directly - they are almost always
	// accessed via a "slice". Slices are a big part of golang but I'm going to
	// mention them only briefly to start with.
	// A slice "looks" like an array but it's only a reference to some section
	// of one.

	var teamSlice []Person // No ellipsis (...) or number means it's a "slice"

	teamSlice = otherTeam[1:] // All but the first element
	fmt.Printf("teamSlice [1:] is %q\n", teamSlice)
	teamSlice = otherTeam[:len(otherTeam)-1] // All but the last element
	fmt.Printf("teamSlice [:len()-1] is %q\n", teamSlice)
	teamSlice = otherTeam[2:3] // A range of elements
	fmt.Printf("teamSlice [2:3] is %q\n", teamSlice)

	fmt.Printf("----------------------------------------------------------\n\n")


	// LESSON 2 - Flow Control Statements //////////////////////////////////////
	fmt.Printf(" LESSON 2 \n\n")

	fmt.Printf("Simple if statement:\n")
	if isValid {
		fmt.Printf("  Valid\n")
	} else {
		fmt.Printf("  Invalid\n")
	}

	fmt.Printf("if with pre-assignment:\n")
	if result := someFunction(); result > 0 {
		// result only exists within this if statement.
		fmt.Printf("  result %d > 0\n", result)
	} else {
		// result also has scope within the else-clause
		fmt.Printf("  result %d <= 0\n", result)
	}

	// The following won't work because result only exists within the if clause:
	// fmt.Printf("result is: %d", result)

	fmt.Printf("switch statements:\n")
	aResult := someFunction()
	switch aResult {
	case 0:
		fmt.Printf("  result is 0\n")
	case 1, 2: // multiple possibilities per "case"
		fmt.Printf("  result is 1 or 2\n")
	default:
		fmt.Printf("  result is neither 1 nor 0\n")
	}

	fmt.Printf("switch with comparison operators:\n")
	switch {
	case aResult > 5:
		fmt.Printf("  aResult > 5")
	case aResult <= 1:
		fmt.Printf("  aResult <= 1")
	default:
		fmt.Printf("  1 < aResult <= 5")
	}

	fmt.Printf("switch with assignment\n")
	switch otherResult := someFunction(); {
	case otherResult < 2:
		fmt.Printf("  otherResult == %q\n", otherResult)
	default:
		fmt.Printf("  otherResult >= 1\n")
	}

	// This won't work because otherResult exists only in the switch:
	// fmt.Printf("  otherResult is %d\n", otherResult)

	fmt.Printf("Iterating with range\n")
	// Iterating with "range" is very concise and easy
	for index, person := range team {
		// The scope of index and person is *inside* the for loop - they don't
		// exist outside.
		fmt.Printf("  %d: Name: %v, Age: %2d\n", index, person.Name, person.Age)
	}

	// fmt.Printf("%d", index) // Won't work because it's undeclared

	fmt.Printf("Iterating by counting\n")
	for index := 0; index < len(team); index++ {
		fmt.Printf("  %d: Name: %v, Age: %2d\n", index, team[index].Name,
			team[index].Age)
	}

	fmt.Printf("For can work like 'while' in other languages:\n")
	for someFunction() < 0 {
		fmt.Printf("gets printed while someFunction() returns < 0 \n")
	}

	// LESSON 3 - Maps, make() and more slicing ////////////////////////////////
	fmt.Printf(" LESSON 3 - maps and dynamic arrays \n\n")

	var nameToAge = map[string]int{"Bob": 24, "Alice" : 25, "Jane": 23 }

	fmt.Printf("Alice is %d years old\n", nameToAge["Alice"])

	fmt.Printf("\nIterating over a map\n")
	for name, age := range nameToAge {
		fmt.Printf("   %q is %d years old\n", name, age)
	}


	
	fmt.Printf("\nDynamic Allocaton of slices or maps\n")

	var dynamicPerson = &Person{"Bob", 24} // Dynamically allocating 1 thing

	var dynamicTeam = make([]Person, 0, 3) // initial size is 0, capacity is 3

	// Basically an array of size 5 is created and a slice length 0 is returned
	//

	// Appending upto the capacity 
	dynamicTeam = append(dynamicTeam, Person{"Alice", 25})
	fmt.Printf("After appending Alice dynamicTeam is %v\n", dynamicTeam)
	dynamicTeam = append(dynamicTeam, *dynamicPerson)
	fmt.Printf("After appending Bob dynamicTeam is %v\n", dynamicTeam)
	dynamicTeam = append(dynamicTeam, Person{"Jane", 23})
	fmt.Printf("After appending Jane dynamicTeam is %v\n", dynamicTeam)

	 // The array is now "full" to it's capacity of 3
	 
	 
	// Copy the slice - this is only a copy of the "window" onto the original
	// array that was allocated by make()
	savedSlice := dynamicTeam  

	 // The next append forces a reallocation and copy
	dynamicTeam = append(dynamicTeam, Person{"Fred", 22})


	savedSlice[0] = Person{"Greg", 21} // This changes the original array size 3
	fmt.Printf("After appending Fred savedSlice is %v\n", savedSlice)

	// ... but dynamicTeam is unaffected because the last append reallocated and
	// copied it.
	fmt.Printf("After appending Fred dynamicTeam is %v\n", dynamicTeam)

	// So beware of the potential for a subtle bug here! When you append to an
	// slice the returned slice may not be "looking" at the same memory that
	// it was before the append. 


    // You can also "make()" maps

	myMap := make(map[string]int, 3) //  a map with "capacity" 3
	myMap["Fred"] = 13

	

}
