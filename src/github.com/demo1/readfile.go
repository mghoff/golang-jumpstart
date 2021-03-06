package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

func readfiledemo() {
    // read1()
    // read2()
    read3()
}

func read1() {
    dat, err := ioutil.ReadFile("dunkirk.txt")
    if err != nil {
        panic(err)
    } else {
        // fmt.Println(dat)
        fmt.Println(string(dat))
    }
}

func read2() {
    f, _ := os.Open("dunkirk.txt")
    b := make([]byte, 60)

    for {
        r, _ := f.Read(b)
        if r == 0 {
            break
        }
        fmt.Println(r)
        fmt.Println(string(b))
    }
}


func read3() {
    f, _ := os.Open("dunkirk.txt")
    b := make([]byte, 60)

    for {
        r, err := f.Read(b)
        if err != nil && err != io.EOF {
            panic(err)
        }
        if r == 0 {
            break
        }
        fmt.Println(string(b[:r]))
    }

}
