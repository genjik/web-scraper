package webscraper

import (
    "golang.org/x/net/html"
)

func parent(n *html.Node) *html.Node {
    return n.Parent
}
func prevSibling(n *html.Node) *html.Node {
    return n.PrevSibling
}
func nextSibling(n *html.Node) *html.Node {
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

// Returns first parent element that is ElementNode. Otherwise, returns nil
func (e Element) parent() Element {
    temp := e.node.Parent
    return traverse(temp, parent)
}

// Returns first child element that is ElementNode. Otherwise, returns nil
func (e Element) firstChild() Element {
    temp := e.node.FirstChild
    return traverse(temp, nextSibling)
}

// Returns first previous sibling element that is ElementNode.
// Otherwise, returns nil
func (e Element) prevSibling() Element {
    temp := e.node.PrevSibling
    return traverse(temp, prevSibling)
}

// Returns first next sibling element that is ElementNode.
// Otherwise, returns nil
func (e Element) nextSibling() Element {
    temp := e.node.NextSibling
    return traverse(temp, nextSibling)
}


func getNextSibling(e Element) Element {
    temp := e.node.NextSibling
    return traverse(temp, nextSibling)
}
func getPrevSibling(e Element) Element {
    temp := e.node.PrevSibling
    return traverse(temp, prevSibling)
}
func getParent(e Element) Element {
    temp := e.node.Parent
    return traverse(temp, parent)
}

func (e Element) findElement(getSibling func(e Element) Element, tag string, attrs []string) Element {
    pseudoEl := createPseudoEl(tag, attrs)

    if (e.node == nil) {
        return Element{}
    }

    // getParent() or getPrevSibling() or getNextSibling()
    temp := getSibling(e)

    for temp != (Element{}) {
        if temp.compareTo(pseudoEl) == true {
            return temp
        }
        // getParent() or getPrevSibling() or getNextSibling()
        temp = getSibling(temp)
    }

    return Element{}
}

func (e Element) findElements(getSibling func(e Element) Element, tag string, limit int, attrs []string) []Element {
    var elements []Element
    pseudoEl := createPseudoEl(tag, attrs)

    if (e.node == nil) {
        return elements
    }

    // getParent() or getPrevSibling() or getNextSibling()
    temp := getSibling(e)

    for temp != (Element{}) {
        if limit == 0 {
            break
        }

        if temp.compareTo(pseudoEl) == true {
            elements = append(elements, temp)
            limit -= 1
        }

        // getParent() or getPrevSibling() or getNextSibling()
        temp = getSibling(temp)
    }

    return elements
}
