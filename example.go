package main

import (
	"./hac"
	"fmt"
)

func main() {
	a, err := hac.NewHAC("UserName", "Password", "Database")
	if err != nil {
		panic(err)
	}

	g, _ := a.Grades()
	fmt.Println(g)
}
