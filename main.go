package main

import (
    "fmt"
    "strings"
    "os"
)

var _ = fmt.Println
var _ = strings.NewReader
var _ = os.Open

func main() {
    //r, err := os.Open("test/index.html")

    //if err != nil {
    //    fmt.Println(err)
    //    return
    //}
    r := strings.NewReader("jsonjson")

    root, err := GetRootElement(r)

    if err != nil {
        fmt.Println(err)
        return 
    }

    fmt.Printf("%+v\n", root.node.LastChild)
    //fmt.Printf("%+v\n", root.node.LastChild) //body
    //fmt.Printf("%+v\n", root.node.LastChild.FirstChild) //body > div1
    //fmt.Printf("%+v\n", root.node.LastChild.FirstChild.NextSibling) //body > div2
    //fmt.Printf("%+v\n", root.node.LastChild.FirstChild.NextSibling.NextSibling) //body > div2
}
