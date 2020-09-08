package main

import (
    "golang.org/x/net/html"
    "fmt"
)

var _ = fmt.Print

type Element struct {
    node *html.Node
}

func (e Element) FindOne(tag string, recursive bool, attrs ...string) Element {
    //validatedAttrs := validateAttrs(attrs)

    if (e.node == nil) {
        return Element{}
    }

    temp := e.node.FirstChild
    for temp != nil {
        //if Element{}.contains(tag, validatedAttrs) == true {
        //    return Element{temp}
        //}
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
