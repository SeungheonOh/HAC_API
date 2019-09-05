# HAC_API
HAC_API is a convenient method to fetch your classes and grades for HomeAccessCenter

## Example
```
package main

import (
    "github.com/SeungheonOh/HAC_API
    "fmt"
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
```