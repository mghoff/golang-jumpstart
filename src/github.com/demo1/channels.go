package main

import "fmt"

func channeldemo() {

    ch := make(chan string)

    go func(msg string) {
        ch <- msg
    }("Hello!")

    recv := <- ch

    fmt.Println(recv)
}
