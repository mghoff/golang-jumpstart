package main

import (
    "fmt"
    // io
    "io/ioutil"
    "os"
)

func writefiledemo() {
    // write1()
    write2()
}

func write1() {
    b := []byte("This is a text string.\nIt is to be written to file.\n")

    ioutil.WriteFile("txtfile_writefiledemo.txt", b, 0644)

    fmt.Println("writeFile Complete")
}

func write2() {
    f, _ := os.OpenFile("txtfile_writefiledemo.txt", os.O_APPEND | os.O_WRONLY, 0777)
    b := []byte("This is an new line added to original file using write2().")
    defer f.Close()

    f.Write(b)
    fmt.Println(f.Stat())
}
