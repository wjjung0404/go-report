// go package give a modullize and reuse of code
// you can use a small package, make a program

// go give a many pacakge at develop programe function from standard package
// and it's downloaded default
package main

// import
// when you should anywhere package use, write import and package.
import "fmt"

// import a package, GO compiler searching a GOROOT or GOPATH,
// standard package is GOROOT/pkg and when user or Third Party Package, searching a GOPATH.pkg
// GOROOT evviroment are aotomatic install from go download, but GOPATH envirment are use setting

// main pacakge are specially recognized
// package name is main, it's make to share library to executable program
// main() is make Starting point, the Entry point
// when makes a share library, noy use main package or main()
func main() {
	// package name's first part are UPPERCASE, it can use because public stat,
	// but LOWERCASE use, can't use it because private
	fmt.Println()
}
