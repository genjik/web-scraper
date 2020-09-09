package webscraper

import (
    "golang.org/x/net/html"
    "fmt"
    "io"
)

var _ = fmt.Print

type Element struct {
    node *html.Node
}

func GetRootElement(r io.Reader) (*Element, error) {
    root, err := html.Parse(r)

    if err != nil {
        return nil, err
    }

    return &Element{root.FirstChild}, nil
}

func (e Element) FindOne(tag string, recursive bool, attrs ...string) Element {
    pseudoEl := createPseudoEl(tag, attrs)

    if (e.node == nil) {
        return Element{}
    }

    temp := e.firstChild()

    if recursive == true {
        return findOneR(temp, pseudoEl)
    }

    return findOne(temp, pseudoEl)
}

func findOne(e Element, pseudoEl Element) Element {
    temp := e

    for temp != (Element{}) {
        if temp.compareTo(pseudoEl) == true {
            return temp
        }

        temp = temp.nextSibling()
    }
    return Element{}
}

func findOneR(e Element, pseudoEl Element) Element {
    temp := e

    for temp != (Element{}) {
        if temp.compareTo(pseudoEl) == true {
            return temp
        }

        if temp.firstChild() != (Element{}) {
            found := findOneR(temp.firstChild(), pseudoEl)
            if found != (Element{}) {
                return found
            }
        }

        temp = temp.nextSibling()
    }
    return Element{}
}

//func (e Element) FindAll(sel selector, limit int, recursive bool) []Element {}

func (e Element) FindParent(tag string, attrs ...string) Element {
    pseudoEl := createPseudoEl(tag, attrs)

    if (e.node == nil) {
        return Element{}
    }

    temp := e.parent()

    for temp != (Element{}) {
        if temp.compareTo(pseudoEl) == true {
            return temp
        }
        temp = temp.parent()
    }

    return Element{}
}

//func (e Element) FindParents(selector, limit int) []Element {}

func (e Element) FindNextSibling(tag string, attrs ...string) Element {
    pseudoEl := createPseudoEl(tag, attrs)

    if (e.node == nil) {
        return Element{}
    }

    temp := e.nextSibling()

    for temp != (Element{}) {
        if temp.compareTo(pseudoEl) == true {
            return temp
        }
        temp = temp.nextSibling()
    }

    return Element{}
}

//func (e Element) FindNextSiblings(selector, limit int) []Element{}

func (e Element) FindPrevSibling(tag string, attrs ...string) Element {
    pseudoEl := createPseudoEl(tag, attrs)

    if (e.node == nil) {
        return Element{}
    }

    temp := e.prevSibling()

    for temp != (Element{}) {
        if temp.compareTo(pseudoEl) == true {
            return temp
        }
        temp = temp.prevSibling()
    }

    return Element{}
}

//func (e Element) FindPrevSiblings(selector, limit int) []Element {}
