package main

import (
    "fmt"
    "math"
)

type Abser interface{
    Abs() float64
}

// エラーになるので要注意！
func main()  {
    var a Abser
    f := MyFloat(-math.Sqrt2)
    v := Vertex{3, 4}

    a = f
    a = &v

    a = v

    fmt.Println(a.Abs())
}