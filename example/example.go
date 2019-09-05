package main

import (
	"fmt"
	hac "github.com/SeungheonOh/HAC_API"
)

func main() {
	a, err := hac.NewHAC("UserName", "Password", "Database") // Initialize connection
	if err != nil {
		panic(err)
	}

	g, _ := a.Grades() // Fetch Classes and Grades into Class structure
	/*
	   type Class struct {
	       ClassName string
	       ClassAvg  uint8
	   }
	*/
	fmt.Println(g)
}
