package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // Generate 3 random numbers between 0 and 9
    var a = rand.Intn(10) + 1
    var b = rand.Intn(10) + 1
    var c = rand.Intn(10) + 1

    fmt.Println("The result of a * b * c is:", a * b * c)
}
