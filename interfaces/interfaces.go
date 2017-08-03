package main

import (
	"fmt"
	"strconv"
)

//A basic Person struct
type Person struct {
	name string
	age  int
}

//Some slices of ints, floats and Persons
type IntSlice []int
type Float32Slice []float32
type PersonSlice []Person

type MaxInterface interface {
	// Len is the number of elements in the collection.
	Len() int
	//Get returns the element with index i in the collection
	Get(i int) interface{}
	//Bigger returns whether the element at index i is bigger that the j one
	Bigger(i, j int) bool
}

//Len implementation for our three types
func (x IntSlice) Len() int     { return len(x) }
func (x Float32Slice) Len() int { return len(x) }
func (x PersonSlice) Len() int  { return len(x) }

//Get implementation for our three types
func (x IntSlice) Get(i int) interface{}     { return x[i] }
func (x Float32Slice) Get(i int) interface{} { return x[i] }
func (x PersonSlice) Get(i int) interface{}  { return x[i] }

//Bigger implementation for our three types
func (x IntSlice) Bigger(i, j int) bool {
	if x[i] > x[j] { //comparing two int
		return true
	}
	return false
}

func (x Float32Slice) Bigger(i, j int) bool {
	if x[i] > x[j] { //comparing two float32
		return true
	}
	return false
}

func (x PersonSlice) Bigger(i, j int) bool {
	if x[i].age > x[j].age { //comparing two Person ages
		return true
	}
	return false
}

//Person implements fmt.Stringer interface
func (p Person) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

/*
 Returns a bool and a value
 - The bool is set to true if there is a MAX in the collection
 - The value is set to the MAX value or nil, if the bool is false
*/
func Max(data MaxInterface) (ok bool, max interface{}) {
	if data.Len() == 0 {
		return false, nil //no elements in the collection, no Max value
	}
	if data.Len() == 1 { //Only one element, return it alongside with true
		return true, data.Get(1)
	}
	max = data.Get(0) //the first element is the max for now
	m := 0
	for i := 1; i < data.Len(); i++ {
		if data.Bigger(i, m) { //we found a bigger value in our slice
			max = data.Get(i)
			m = i
		}
	}
	return true, max
}

func main() {
	islice := IntSlice{1, 2, 44, 6, 44, 222}
	fslice := Float32Slice{1.99, 3.14, 24.8}
	group := PersonSlice{
		Person{name: "Bart", age: 24},
		Person{name: "Bob", age: 23},
		Person{name: "Gertrude", age: 104},
		Person{name: "Paul", age: 44},
		Person{name: "Sam", age: 34},
		Person{name: "Jack", age: 54},
		Person{name: "Martha", age: 74},
		Person{name: "Leo", age: 4},
	}

	//Use Max function with these different collections
	_, m := Max(islice)
	fmt.Println("The biggest integer in islice is :", m)
	_, m = Max(fslice)
	fmt.Println("The biggest float in fslice is :", m)
	_, m = Max(group)
	fmt.Println("The oldest person in the group is:", m)
}
