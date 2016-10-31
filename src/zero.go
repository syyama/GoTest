package main

import "fmt"

func main() {
    // 初期値を与えない場合は
    var i int       // 0
    var f float64   // 0
    var b bool      // false
    var s string    // ""

    fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
