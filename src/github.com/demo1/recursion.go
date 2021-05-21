package main

import "fmt"

func recursion(n int) {
    f := fact(n)
    fmt.Println(n, "factorial is", f)
}

func fact(n int) int {
    if n == 1 {
        return 1
    } else {
        return n * fact(n - 1)
    }
}


func recursionchallenge(n int) {
    if n == 1 {
        fmt.Println("*")
    } else {
        for i := 1; i <= n; i++ {
            fmt.Print("* ")
        }
        fmt.Println()
        recursionchallenge(n-1)
    }
}
