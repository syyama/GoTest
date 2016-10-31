package main

import "fmt"

func main()  {
    var i, j int = 1, 2

    // 代入文を使い暗黙的な肩宣言ができる
    k := 3
    c, python, java := true, false, "no!"

    fmt.Println(i, j, k, c, python, java);
}
