package main // change it

import (
    "fmt"
    "golang.org/x/net/html"
)

var _ = fmt.Println


func (e Element) FindAllChildren(recursive bool, limit int) []Element {
    var elements []Element = make([]Element, 0)

    temp := e.node.FirstChild

    if limit < 1 {
        for temp != nil {
            if temp.Type == html.ElementNode {
                elements = append(elements, Element{temp})
            }
            temp = temp.NextSibling
        }
    }

    if limit >= 1 {
        for temp != nil {
            if limit == 0 {
                break
            }
            if temp.Type == html.ElementNode {
                elements = append(elements, Element{temp})
            }
            temp = temp.NextSibling
            limit -= 1
        }
    }
    fmt.Println(len(elements), cap(elements))
    fmt.Println(elements)

    return elements
}

//func (n *Element) FindChildById(id string) Element {}
//
//func (n *Element) FindChildrenByClass(class string, recursive bool, limit int) []Element {}
//
//func (n *Element) FindChildrenByElement(element string, recursive bool, limit int) []Element {}
