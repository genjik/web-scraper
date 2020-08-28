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

func containsClass(attrs1, attrs2 []html.Attribute) bool {
    for _, attr := range attrs1 {
        if attr.Key != "class" {
            continue
        }
        for _, attr2 := range attrs2 {
            if attr2.Key != "class" {
                continue
            }

            receiverClasses := strings.Split(attr.Val, " ")
            classes := strings.Split(attr2.Val, " ")
            
            sort.Strings(receiverClasses)
            sort.Strings(classes)

            if len(receiverClasses) < len(classes) {
                return false
            }

            count := 0
            for _, str := range receiverClasses {
                for _, str2 := range classes {
                    if str == str2 {
                        count += 1
                    }
                }
            }

            if len(receiverClasses) > len(classes) &&
            count == len(classes) {
                    return true
            }

            if len(receiverClasses) == len(classes) &&
            count == len(receiverClasses) {
                    return true
            }
        }
    }
    return false
}
