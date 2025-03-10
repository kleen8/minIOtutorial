// Go doesn't use exceptions it uses error value's instead of
package main

import (
	"errors"
	"fmt"
	"log"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("Can't divide by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)   
    if err != nil {
        log.Fatalf("error: %s\n", err.Error())
        return
    }
    fmt.Println("Result: ", result)

}
