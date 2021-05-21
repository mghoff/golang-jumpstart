package main

import "fmt"

func pointer1() {
    var a *int
    b := 6

    a = &b

    fmt.Println("Address ", a)
    fmt.Println("Value", *a)
}

func pointer2() {
    c := 5
    // increment(c)
    // increment(c)
    // fmt.Println("Called from pointer2() ", c)

    incrementbypointer(&c)
    incrementbypointer(&c)
    fmt.Println("Called from pointer2() ", c)
}

func increment(c int) {
    c++
    fmt.Println("Called within increment() ", c)
}

func incrementbypointer(c *int) {
    *c++
    fmt.Println("Called within incrementbypointer() ", *c)
}
