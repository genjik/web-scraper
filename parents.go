package webscraper

import (
    "golang.org/x/net/html"
)

func (e Element) Parent() Element {
    var element Element
    
    parent := e.node.Parent

    if parent != nil && parent.Type == html.ElementNode {
        element = Element{parent}
    }

    return element 
}
