/*
module = collection of Go packages stored in a file tree with a go.mod
file at its root.
go.mod file defines the module's module path
*/

// creating a new module
package hello

func Hello() string {
	return "Hello, world."
}