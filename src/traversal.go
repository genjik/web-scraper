package webscraper

import (
    "golang.org/x/net/html"
)

// Returns first parent element that is ElementNode. Otherwise, returns nil
func (e Element) parent() Element {
    temp := e.node.Parent
    return traverse(temp, getParent)
}

// Returns first child element that is ElementNode. Otherwise, returns nil
func (e Element) firstChild() Element {
    temp := e.node.FirstChild
    return traverse(temp, getNextSibling)
}

// Returns first previous sibling element that is ElementNode.
// Otherwise, returns nil
func (e Element) prevSibling() Element {
    temp := e.node.PrevSibling
    return traverse(temp, getPrevSibling)
}

// Returns first next sibling element that is ElementNode.
// Otherwise, returns nil
func (e Element) nextSibling() Element {
    temp := e.node.NextSibling
    return traverse(temp, getNextSibling)
}

func getParent(n *html.Node) *html.Node {
    return n.Parent
}
func getPrevSibling(n *html.Node) *html.Node {
    return n.PrevSibling
}
func getNextSibling(n *html.Node) *html.Node {
    return n.NextSibling
}

func traverse(temp *html.Node, t func(n *html.Node) *html.Node) Element {
    for temp != nil {
        if temp.Type == html.ElementNode && temp.Data != "" {
            return Element{temp}
        }
        temp = t(temp)
    }

    return Element{}
}

func nextSibling(e Element) Element {
    temp := e.node.NextSibling
    return traverse(temp, getNextSibling)
}
func prevSibling(e Element) Element {
    temp := e.node.PrevSibling
    return traverse(temp, getPrevSibling)
}
func parent(e Element) Element {
    temp := e.node.Parent
    return traverse(temp, getParent)
}

func findElement(e Element, getSibling func(e Element) Element, tag string, attrs []string) Element {
    pseudoEl := createPseudoEl(tag, attrs)

    if (e.node == nil) {
        return Element{}
    }

    // Either prevSibling() or nextSibling()
    temp := getSibling(e)

    for temp != (Element{}) {
        if temp.compareTo(pseudoEl) == true {
            return temp
        }
        temp = getSibling(temp)
    }

    return Element{}
}

func findElements(e Element, getSibling func(e Element) Element, tag string, limit int, attrs []string) []Element {
    var elements []Element
    pseudoEl := createPseudoEl(tag, attrs)

    if (e.node == nil) {
        return elements
    }

    temp := getSibling(e)

    for temp != (Element{}) {
        if limit == 0 {
            break
        }

        if temp.compareTo(pseudoEl) == true {
            elements = append(elements, temp)
            limit -= 1
        }

        temp = getSibling(temp)
    }

    return elements
}
