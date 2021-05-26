package main

import "fmt"

type profile struct {
    name    string
    address string
    email   string
    phone   string
}

func (p profile) getname() string {
    return p.name
}

func (p profile) getaddress() string {
    return p.address
}

func (p profile) getcontact() map[string]string {
    m := map[string]string{"email": p.email, "phone": p.phone}
    return m
}


func methodsdemo() {

    s := profile {
        name:       "Sherlock Holmes",
        address:    "221B Baker St",
        email:      "sherlock.holmes@gmail.com",
        phone:      "18811904",
    }

    fmt.Println("NAME:", s.getname())
    fmt.Println("EMAIL", s.getcontact()["email"])
}
