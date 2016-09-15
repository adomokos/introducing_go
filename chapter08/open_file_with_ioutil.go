package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    bs, err := ioutil.ReadFile("test.txt")
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }

    str := string(bs)
    fmt.Println(str)
}
