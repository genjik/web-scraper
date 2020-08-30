package main // change it

import (
    "fmt"
    "golang.org/x/net/html"
)

var _ = fmt.Println

func (e Element) FindAllChildren(limit int) []Element {
    var elements []Element
    temp := e.node.FirstChild

    for temp != nil {
        if limit == 0 {
            break
        }
        if temp.Type == html.ElementNode {
            elements = append(elements, Element{temp})
            limit -= 1
        }
        temp = temp.NextSibling
    }

    return elements
}


func (e Element) FindChildrenByClass(class string, limit int) []Element {
    var elements []Element
    
    temp := e.node.FirstChild

    for temp != nil {
        if limit == 0 {
            break
        }
        contains := containsSel(temp.Attr, html.Attribute{"", "class", class}, "class")
        if temp.Type == html.ElementNode && contains {
            elements = append(elements, Element{temp})
            limit -= 1
        }
        temp = temp.NextSibling
    }

    return elements
}

func (e Element) FindChildrenByElement(element string, limit int) []Element {
    var elements []Element
    
    temp := e.node.FirstChild

    for temp != nil {
        if limit == 0 {
            break
        }

        isEqual := compareTypeAndData(Element{temp}, Element{&html.Node{Type: html.ElementNode, Data: element}})

        if temp.Type == html.ElementNode && isEqual {
            elements = append(elements, Element{temp})
            limit -= 1
        }
        temp = temp.NextSibling
    }

    return elements
}

//func (n *Element) FindChildById(id string) Element {}
//
