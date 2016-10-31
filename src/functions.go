package main

import "fmt"

func add(x int, y int) int {
    // パラメータは変数名の後ろに型名を書く！
    return x + y
}

func main() {
    fmt.Println(add(114, 514))
}
