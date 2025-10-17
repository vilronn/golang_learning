package main

import "fmt"

func main() {
    var a string
    fmt.Scan(&a)
    
    if (int(a[0]) + int(a[1]) + int(a[2]) == int(a[3]) + int(a[4]) + int(a[5])) {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}
