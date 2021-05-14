package main // this will tell compiler to create an executable

import (
	"fmt"
	_	"gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
	"example.com/modpkgs/justAnotherPackage"
)

func main() { // starting point of application
	fmt.Println("This is Modules and Packages Program")
	fmt.Println("variable declared in another Package: ", justAnotherPackage.exportedVariable)
}
