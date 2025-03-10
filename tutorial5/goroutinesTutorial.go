package main

import (
    "fmt"
    "time"
)

func sayHello() {
    fmt.Println("Hello from goroutine")
}

func sayBye() {
    fmt.Println("Bye from goroutine")
}

func main() {
    go sayHello() // Run in a separate goroutine
    go sayBye()
    time.Sleep(time.Second) // Give it time to run
}
