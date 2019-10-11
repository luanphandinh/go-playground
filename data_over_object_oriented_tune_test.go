// This example comparing:
// For a project that running large amount of working over time
//
// Instead of saving flag `active` inside a object (Object Oriented Programming)
// -> Then traverse through the list object if active then doSomething()

// We instead have an flags of active flag (true or false) (Data Oriented Programming)
// Loop through the flags then doSomething()
// It could help acquire a positive performance impact
//
// A good notice that sometimes we don't need to put a boolean flag inside an object
// It faster to loop through flags array for `bool`(1 byte)
// instead of looping through flag inside heavy objects (> 1 byte) just for a bool
//
// Guidance from: code::dive conference 2014 - Scott Meyers: Cpu Caches and Why You Care
// https://www.youtube.com/watch?v=WDIkqP4JbkE
package go_playground

import "testing"

// Just make sure this func is work on something
func doSomething() {
	count := 0
	for i := 0; i < 1000; i++ {
		count += 1
	}
}

type Obj struct {
	active bool
}

// Same doSomething() but for object
func (obj *Obj) doSomething() {
	doSomething()
}

var prepared bool
var flags []bool
var objects []*Obj

// Here we prepare 10000 records
func prepareData() {
	if prepared {
		return
	}

	for i := 0; i < 10000; i++ {
		flags = append(flags, true)
		obj := &Obj{true}
		objects = append(objects, obj)
	}
}

func dataDoSomething() {
	for i, val := range flags {
		if val {
			objects[i].doSomething()
		}
	}
}

func objectsDoSomething() {
	for _, obj := range objects {
		if obj.active {
			obj.doSomething()
		}
	}
}

func BenchmarkDataOriented(b *testing.B) {
	prepareData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dataDoSomething()
	}
}

func BenchmarkObjectOriented(b *testing.B) {
	prepareData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		objectsDoSomething()
	}
}
