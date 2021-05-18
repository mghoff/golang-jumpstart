package main

import "fmt"

func add(x int, y int) int {
    return x + y
}

func Boolean() bool {
    var b bool
    b = true
    return b
}

func showvar() {
    var a string = "mystring"
    fmt.Println(a)

    var e int
    fmt.Println(e)

    var b, c int = 1, 2
    fmt.Println(b, c)

    d := "love golang"

    f := a + " : " + d
    fmt.Println(f)
}
