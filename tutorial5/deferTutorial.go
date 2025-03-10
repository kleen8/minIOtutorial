// this is to show that defer is used to wait for something to happen and it is last in first out
package main

import "fmt"

func main() {
    fmt.Println("this prints first")

    defer fmt.Println("This prints last")
    defer fmt.Println("This prints seconds last")

    fmt.Println("This prints secondly")
}
