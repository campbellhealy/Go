package main

/*
This is code added to the file. Although in testing, VS Code formatter notes that certain imports
are not required at the initial test and removes them from the code.

This could cause an issue if not noticed later in the coding excerise

import (
    "fmt"
    "log"
    "net/http"
)
*/

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
