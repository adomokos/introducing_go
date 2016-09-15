package main

import (
    "os"
)

func main() {
    file, err := os.Create("new_file.txt")
    if err != nil {
        // handle the error here
        return
    }
    defer file.Close()

    file.WriteString("test")
}
