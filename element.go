package main //change it

import (
    "golang.org/x/net/html"
    "strings"
    "io"
    "sort"
)

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

// Compares html.Node.Type and html.Node.Data of two elements
func (e *Element) compareTypeAndData(e2 Element) bool {
    if e.node.Type != e2.node.Type {
        return false
    }

    if strings.ToLower(e.node.Data) != strings.ToLower(e2.node.Data) {
        return false
    }

    return true
}

func containsClass(attributes []html.Attribute, attribute html.Attribute) bool {
    for _, attr := range attributes {
        if attr.Key != "class" {
            continue
        }
        if attribute.Key != "class" {
            return false
        }

        aClasses := strings.Split(attr.Val, " ")
        bClasses := strings.Split(attribute.Val, " ")
        
        sort.Strings(aClasses)
        sort.Strings(bClasses)

        if len(aClasses) < len(bClasses) {
            return false
        }

        count := 0
        for _, str := range aClasses {
            for _, str2 := range bClasses {
                if str == str2 {
                    count += 1
                }
            }
        }

        if len(aClasses) > len(bClasses) && count == len(bClasses) {
            return true
        }

        if len(aClasses) == len(bClasses) && count == len(aClasses) {
            return true
        }
    }
    return false
}
