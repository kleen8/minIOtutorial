package main

import "fmt"

func main() {
    ch := make(chan int) // Create a channel

    go func() {
        ch <- 42 // Send value into the channel
    }()

    fmt.Println(<-ch) // Receive value from the channel
}

