package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("test.txt")
    if err != nil {
        fmt.Println("Error occurred when reading the file: ", err)
        return
    }
    defer file.Close()

    // get the file size
    stat, err := file.Stat()
    if err != nil {
        return
    }

    // read the file
    bs := make([]byte, stat.Size())
    _, err = file.Read(bs)

    if err != nil {
        return
    }

    str := string(bs)
    fmt.Println(str)
}
