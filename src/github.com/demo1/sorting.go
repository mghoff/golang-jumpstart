package main

import (
    "fmt"
    "sort"
)

func sortdemo() {
    s := []string{"z", "x", "b", "a", "y"}
    sort.Strings(s)
    fmt.Println("Sorted String: ", s)

    nums := []int{4, 1, 7, 3, 2, 5, 6}
    sort.Ints(nums)
    fmt.Println("Sorted Integer: ", nums)
}
