package webscraper

import (
    "golang.org/x/net/html"
)

func (e Element) firstChild() Element {
    temp := e.node.FirstChild
    return traverse(temp)
}

func (e Element) nextSibling() Element {
    temp := e.node.NextSibling
    return traverse(temp)
}

func traverse(temp *html.Node) Element {
    for temp != nil {
        if temp.Type == html.ElementNode && temp.Data != "" {
            return Element{temp}
        }
        temp = temp.NextSibling
    }
    return Element{}
}
