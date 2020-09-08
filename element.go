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
    return findOneR(temp, pseudoEl)
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

//func (e Element) FindParent(selector) Element {}
//func (e Element) FindParents(selector, limit int) []Element {}
//
//func (e Element) FindNextSibling(selector) Element {}
//func (e Element) FindNextSiblings(selector, limit int) []Element{}
//
//func (e Element) FindPrevSibling(selector) Element {}
//func (e Element) FindPrevSiblings(selector, limit int) []Element {}
