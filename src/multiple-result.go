package main

import "fmt"

func swap(x, y string) (string, string) {
    // 関数は複数の戻り値を返すことが可能
    return y, x
}

func main() {
    a, b := swap("hello", "world")
    fmt.Println(a, b)
}