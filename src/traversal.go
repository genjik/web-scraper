package webscraper

import (
    "golang.org/x/net/html"
)

// Returns first parent element that is ElementNode. Otherwise, returns nil
func (e Element) parent() Element {
    temp := e.node.Parent
    return traverseUp(temp)
}

// Returns first child element that is ElementNode. Otherwise, returns nil
func (e Element) firstChild() Element {
    temp := e.node.FirstChild
    return traverseForward(temp)
}

// Returns first previous sibling element that is ElementNode.
// Otherwise, returns nil
func (e Element) prevSibling() Element {
    temp := e.node.PrevSibling
    return traverseBackward(temp)
}

// Returns first next sibling element that is ElementNode.
// Otherwise, returns nil
func (e Element) nextSibling() Element {
    temp := e.node.NextSibling
    return traverseForward(temp)
}

func traverseUp(temp *html.Node) Element {
    for temp != nil {
        if temp.Type == html.ElementNode && temp.Data != "" {
            return Element{temp}
        }
        temp = temp.Parent
    }
    return Element{}
}

func traverseBackward(temp *html.Node) Element {
    for temp != nil {
        if temp.Type == html.ElementNode && temp.Data != "" {
            return Element{temp}
        }
        temp = temp.PrevSibling
    }
    return Element{}
}

func traverseForward(temp *html.Node) Element {
    for temp != nil {
        if temp.Type == html.ElementNode && temp.Data != "" {
            return Element{temp}
        }
        temp = temp.NextSibling
    }
    return Element{}
}
