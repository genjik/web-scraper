package main

import (
    "golang.org/x/net/html"
    "strings"
    "sort"
    "fmt"
)

var _ = fmt.Print

type Element struct {
    node *html.Node
}

func (e Element) FindOne(tag string, recursive bool, attrs ...string) Element {
    validatedAttrs := validateAttrs(attrs)

    if len(validatedAttrs) < 0 {
        return Element{}
    }

    if (e.node == nil) {
        return Element{}
    }

    temp := e.node.FirstChild
    for temp != nil {
        
    }

    return Element{}
}

func (e Element) contains(tag string, attrs []string) bool {
    if e.node.Type != html.ElementNode {
        return false
    }

    if e.node.Data != tag {
        return false
    }

    if (len(attrs) / 2) > len(e.node.Attr) {
        return false
    }

    for _, i := range e.node.Attr {
                
    }

    return true
}

func validateAttrs(attrs []string) []string {
    var newAttrs []string 

    if len(attrs) < 1 { 
        return newAttrs
    }

    // if the len(attrs) == odd, remove last value
    if len(attrs) % 2 != 0 { 
        attrs = attrs[:len(attrs)-1]
    }

    for i:=0; i < len(attrs); i++ {
        if i % 2 != 0 { 
            continue
        }

        count := 0

        // iterates even indexes from the attrs[0] till attrs[i]
        for j:=0; j < i; j++ {
            if j % 2 != 0 {
                continue
            }

            if attrs[i] == attrs[j] {
                count += 1
            }
        }

        if count < 1 {
            newAttrs = append(newAttrs, attrs[i], attrs[i+1])
        }
    }

    return newAttrs
}

func main() {
}

// Functions that are used for testing
func compareStr(s1, s2 []string) bool {
    if len(s1) != len(s2) {
        return false
    }

    sort.Strings(s1)
    sort.Strings(s2)

    for i, v := range s1 {
        if strings.ToLower(v) != strings.ToLower(s2[i]) {
            return false
        }
    }

    return true
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
