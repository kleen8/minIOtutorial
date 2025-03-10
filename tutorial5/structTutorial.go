// Go does not have classes but you can define methods on structs like below
package main

import "fmt"

type Person struct {
    Name string
}

func (p *Person) Rename(newName string) {
    p.Name = newName
}

func main() {
    p := Person{ Name: "john" }
    fmt.Println(p.Name)
    p.Rename("doe")
    fmt.Println(p.Name)
}
