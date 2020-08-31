package main //change it

import (
    "fmt"
    "golang.org/x/net/html"
)

var _ = fmt.Println

func (e Element) Parent() Element {
    var element Element
    
    parent := e.node.Parent

    if parent != nil && parent.Type == html.ElementNode {
        element = Element{parent}
    }

    return element 
}
