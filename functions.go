package main

import "fmt"

func add(x int, y int) int {
    return x + y
}

func nadd(x, y int) int {
    return x + y
}

func addPointer(x [1]int, y [1]int) int {
    return x[0] + y[0]
}

func main() {
    fmt.Println(add(22, 33))
    fmt.Println(nadd(22, 33))
    fmt.Println(addPointer([1] int {1}, [1] int {2}))
}

