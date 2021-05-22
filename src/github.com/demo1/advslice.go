package main

import "fmt"

func advslice() {
    s := make([]int, 3)
    for i := 0; i < 3; i++ {
        s[i] = i + 1
    }

    s = append(s, 4, 5, 6, 7, 8, 9)

    fmt.Println("s = ", s)


    // copy slice
    b := make([]int, len(s))
    copy(b, s)
    fmt.Println("b = ", b)


    //  cut slice - remove 3,4
    // s = append(s[:1], s[4:]...)
    // fmt.Println("cut(3,4): s =", s)


    // delete by index
    s = delbyidx(2, s)
    fmt.Println("delete by index:", s)
}

func delbyidx(i int, a []int) []int {
    a = append(a[:i-1], a[i:]...)
    return a
}
