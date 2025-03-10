// Go uses interfaces instead of classical inheritance
package main

import "fmt"

type Speaker interface {
    Speak()
}

type Dog struct{}
func (d Dog) Speak() { fmt.Println("Woof") }

type Cat struct{}
func (c Cat) Speak() { fmt.Println("Meaow") }

func main() {
    var s Speaker = Dog{}
    s.Speak()

    s = Cat{}
    s.Speak()
}
