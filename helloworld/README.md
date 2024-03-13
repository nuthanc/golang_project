### Hello World and Running the Program

* Create a file called main.go(Even hello.go worked)
* Run code using `go run main.go`
  * `go run` compiles and executes the file
* `go build` just compiles it and we get an executable
  * Execute by calling `./main`
* `go fmt` for formatting code in each file in the current directory
* `go install` for Compiling and Installing a package
* `go get` for Downloading the raw source code of someone else's package
* `go test` for running any tests associated with the current project

### Go Packages

* Package == Project or Workspace
  * Collection of common source code files
* Every first line in go file should start with a package
* Types of Packages
  * Executable: Generates a file that we can run
  * Reusable: Code used as helpers
* package `main` is executable and it generates an executable with `go build`
* Any other package name would not generate an executable with `go build`, so is a Reusable package
* Also package main should have a `main` function
  * This is run automatically when the program runs

### Import Statements

* import from another package
* "fmt" is part of Standard library of Go
* Go's standard packages: https://pkg.go.dev/std

### Functions and Organization

* `func` for functions
* File Organization: package, imports and functions

### Array and Slice

* Array: Fixed length list of things
* Slice: An array that can grow or shrink
* They both must have identical types of elements
* Example of empty slice: `cards := []string{}` and with 1 element `cards := []string{"Diamond"}`
* Iterating with range
```go
for index, card := range extra_cards {
  fmt.Println(index, card)
}
```

### OO approach vs Approach in Go

* https://app.diagrams.net/#Hnuthanc%2Fgolang_project%2Fmain%2Fdiagrams%2F03%2Fdiagrams.xml
* In Object-Oriented programming languages like Python and Java, we usually create a `Deck class` with `cards property` and `methods like print, shuffle, saveToFile`
* But Go doesn't have classes
* So we define a new **type** in Go and then attach Functions with a **receiver**
* `type deck []string` in deck.go and use it in main.go like so `cards := deck{"Five of Spades", "Six of Diamonds"}`
  * It's like a `subclass` of Slice of string in OO terms

### Function Receivers

```go
package main

import "fmt"

type deck []string

// By convention single letter that matches the abbreviation of the receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Can rename it to instance as well
func (instance deck) first() string {
	return instance[0]
}
```
* Any variable of type "deck" now gets access to the "print" and "first" methods
* In the above example, 
  * d is the "Actual copy" of the deck we're working on(Similar to "this" in JS)
  * deck is the "type" to which we want to attach this method to. Every variable of type "deck" can call this function on itself

### Implementation of Methods

* https://app.diagrams.net/#Hnuthanc%2Fgolang_project%2Fmain%2Fdiagrams%2F03%2Fdiagrams.xml
* `newDeck` doesn't need to have receivers as we are not going to work on existing deck
* cardSuits and cardValues are slices of strings instead of deck because they are just Half values and we also don't require methods associated with deck
* Since `deal` function is in the same package, we can just call it in `main.go`

### Byte Slices

* https://app.diagrams.net/#Hnuthanc%2Fgolang_project%2Fmain%2Fdiagrams%2F03%2Fdiagrams.xml#%7B%22pageId%22%3A%22ee34c4b0-917b-729f-15d4-373b03a1a540%22%7D
* Think them as String
* Slice of ASCII character code
* asciitable.com
* Type Conversion example `stringSlice := []string(d)`
* 0666 Anyone can read and write to this file
* Author used `rand.NewSource`, `rand.New` and `time.Now().UnixNano` for seed(for randomness)

### Creating a go.mod File

```txt
go: go.mod file not found in current directory or any parent directory; see 'go help modules'

Inside the cards project directory run the following:

go mod init helloworld

Then, you will be able to use the run test function from within VSCode, and/or run go test from the terminal.
```

### Testing with Go

* https://app.diagrams.net/#Hnuthanc%2Fgolang_project%2Fmain%2Fdiagrams%2F03%2Fdiagrams.xml#%7B%22pageId%22%3A%22f9f645f6-c63f-0355-c867-11bf8a31ca29%22%7D
* `go mod init helloworld` to create a module
* To make a test, create a new file ending in `_test.go`, example `deck_test.go`
* To run all tests in a package, run `go test`
* To see *Run tests* option in Go, go to Settings -> Search "Go test enable code lens" -> Enable if true, enabled code lens for running and debugging tests (runtest)
  * If it's already enabled but still not showing, try to create at least one function in the test file, empty function is enough. After creating at least one function, the button should be showing.

##### Deciding What to Test

* https://app.diagrams.net/#Hnuthanc%2Fgolang_project%2Fmain%2Fdiagrams%2F03%2Fdiagrams.xml#%7B%22pageId%22%3A%2225d2e42e-2f1e-fff1-c784-e222438bc7f0%22%7D
* What do you really care about the newDeck function
  * Length
  * Items check etc
* It's really upto the developer to have the checks
* D 23: tests
* D 24: tests
* Test functions automatically called with testHandler `t *testing.T`
* `go test` or VS code run test
* Fail the test, then pass it
* Cleanup needs to be handled by us if something in the middle crashes(D 27: cleanup)
* D 26: save and load -> Delete at beginning and end and use "_decktesting" filename(temporary)

### Packages

* To use variables or functions from another package, without module configuration, we need to have that file in GOROOT or GOPATH
* GOROOT: Path where Go binaries is installed
* GOPATH: Environment variable where path to Go workspace is set
* Folder and package name(exception of main package) should be the same
* For getting external package(like github), first we need to download using `go get github.com/pioz/faker`
  * This will download to our GOPATH

### Scoping

* Package Private(Package Scope) and Package Public(Public Scope)
* We can use variables from another file without any changes if it is in the same package
  * Run using `go run .` instead of specifying multiple files
* Uppercase 1st letter variable/function is exported(Public scope) and can be accessed from other packages, where lowercase 1st letter variable/function is accessed only within that package

### init function

* This is called by Go runtime
* It does not have any arguments or return values
* Runs before main function
* Used to initialize global variables, establishing db connection and logging
* There can be multiple init functions in a package
  * The execution order is based Lexicographically(Alphabetical)
* There can even be multiple init functions in a file

### Go Modules

* `go mod init <module-name>`
* go get will add the dependencies in go.mod and go.sum
* go.sum containing hash and go.mod containing require
* `go get github.com/pioz/faker@latest` or `go get -u` for latest
* `go get github.com/pioz/faker@v0.1.1`
* `go mod tidy` for cleaning up go.sum

### Vendoring

* `go mod vendor -v`
* vendor directory -> Entire package and modules.txt

### dep tool

* Turn off go module `go env -w GO111MODULE=off`
* Project code in GOPATH directory
* Download dep from github.com/golang/dep/releases or Google for it
* `dep init`
* Dependencies of Go project in *Gopkg.lock* and *Gopkg.toml*
* `dep ensure`
* `dep status` for all dependencies
* `dep check`