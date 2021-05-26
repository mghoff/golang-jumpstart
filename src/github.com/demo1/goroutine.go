package main

import (
    "fmt"
    // "time"
)

func goroutinedemo(n int) {
    count(n)
    // go count(n) // need "time" package
    // time.Sleep(time.Second * 1)
}

func count(num int) {
    for i := 0; i <= num; i++ {
            fmt.Println(i)
    }
}
