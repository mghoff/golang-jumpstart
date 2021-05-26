package main

import "fmt"

type userprofile interface {
    getname() string
    getaddress() string
    getcontact() map[string]string
}

func getuserinfo(u userprofile) {
    fmt.Println("NAME", u.getname())
    fmt.Println("CONTACT", u.getcontact()["email"])
}

func interfacedemo() {
    s := profile {
        name:       "Sherlock Holmes",
        address:    "221B Baker St",
        email:      "sherlock.holmes@gmail.com",
        phone:      "18811904",
    }

    getuserinfo(s)
}
