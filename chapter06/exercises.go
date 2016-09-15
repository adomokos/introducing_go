package main

import "fmt"

func half(x int) (int, bool) {
    result := x / 2
    even := x % 2 == 0
    return result, even
}

func main() {
    fmt.Println(half(2))
}
