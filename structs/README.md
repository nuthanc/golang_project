### Struct

- Collection of properties that are related together
- https://app.diagrams.net/#Hnuthanc%2Fgolang_project%2Fmain%2Fdiagrams%2F04%2Fdiagrams.xml#%7B%22pageId%22%3A%22d5842bfb-84b7-b4a5-3ead-a080b650881c%22%7D
- In python, this is like dictionary
- D 13: people -> Define what fields the struct has
- Check `main.go` for examples
- string: "", int: 0, float: 0, bool: false

* `%+v` for fields and values of struct

```go
package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main() {
	// 1st way
	alex := person{"Alex", "Harrold"}
	fmt.Println(alex)

	// 2nd way: Author recommends this because
	// it would work even if order changed
	naruto := person{firstName: "Naruto", lastName: "Uzumaki"}
	fmt.Println(naruto.lastName)

	// 3rd way: For Zero values instead of initializing
	var lee person
	// Default value of "" for string
	fmt.Println(lee.firstName)
	fmt.Printf("%+v\n", lee)
	lee.firstName = "Rock"
	lee.lastName = "Lee"
	fmt.Printf("%+v", lee)
}
```

### Pointers

* https://app.diagrams.net/#Hnuthanc%2Fgolang_project%2Fmain%2Fdiagrams%2F04%2Fdiagrams.xml
* A pointer holds the memory address of a value.
* Go is `Pass by Value` language
* Go `copies whatever is passed`(Differs from JS objects where the value of the variable is a reference)
```go
alexPointer := &alex
alexPointer.updateName("Naruto")

func (pointerToPerson *person) updateName(firstName string) {
  (*pointerToPerson).firstName = firstName;
}

```

### Go shortcut with Pointers

* Go automatically converts to Pointer if value is passed to Pointer receiver
```go
// Allows to call updateName even though
// alex is of type person and not *person
alex.updateName("Naruto")

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
  // But even the below works!
  // pointerToPerson.firstName = newFirstName
}
```

### Gotchas with Pointers

* Whenever you pass an integer, float, string, or struct into a function, Go creates a copy of each argument, and these copies are used inside of the function
* But this is different with Slices, Maps, Channels, Pointers and Functions
* Slice internally creates an Array and a structure that records the length of the slice, the capacity of the slice, and a *reference to the underlying array*
* https://app.diagrams.net/#Hnuthanc%2Fgolang_project%2Fmain%2Fdiagrams%2F04%2Fdiagrams.xml#%7B%22pageId%22%3A%222eb5d477-d379-13a7-dc45-afbad306a350%22%7D
* D 19: slice in funcs
* D 20: Pass by Value
* Even though Slice is copied, it is still pointing to the original Array
* D 21: Value Types and Reference Types
* **Value Types**: int, float, string, bool, struct
  * Use Pointers to change these things in a function
* **Reference Types**: slice, map, channel, pointer, function