package main

type Person struct {
    Dni     string `json:"dni,omitempty"`
    Name    string `json:"name,omitempty"`
    Surname string `json:"surname,omitempty"`
    Gender  string `json:"gender,omitempty"`
    Addr    string `json:"addr,omitempty"`
    Phone   string `json:"phone,omitempty"`
    Weight  int    `json:"weight,omitempty"`
}

type PersonList struct {
    Persons []Person `json:"persons,omitempty"`
}
