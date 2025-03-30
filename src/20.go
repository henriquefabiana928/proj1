package main

import "fmt"

func main() {
    // Example of using goroutines
    go doSomething()
}

func doSomething() {
    fmt.Println("Goroutine 1 started")
    time.Sleep(2 * time.Second)
    fmt.Println("Goroutine 1 finished")
}
