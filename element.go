package webscraper

import (
    "golang.org/x/net/html"
    "strings"
    "sort"
    "io"
)

type Element struct {
    node *html.Node
}

// Returns Element struct that containts pointer to <html> tag as a html.Node
func GetRootElement(r io.Reader) (*Element, error) {
    root, err := html.Parse(r)

    if err != nil {
        return nil, err
    }

    return &Element{root.FirstChild}, nil
}

func (e Element) Data() string {
    temp := e.node

    if temp.Type == html.ElementNode {
        if temp.FirstChild != nil {
            if temp.FirstChild.Type == html.TextNode {
                return temp.FirstChild.Data
            }
        }
    }

    return ""
}

// Compares html.Node.Type and html.Node.Data of two elements
func compareTypeAndData(e, e2 Element) bool {
    if e.node.Type != e2.node.Type {
        return false
    }

    if strings.ToLower(e.node.Data) != strings.ToLower(e2.node.Data) {
        return false
    }

    return true
}

// should be used only for class and id
func containsSel(attributes []html.Attribute, attribute html.Attribute, attrKey string) bool {
    for _, attr := range attributes {
        if strings.ToLower(attr.Key) != strings.ToLower(attrKey) {
            continue
        }
        if strings.ToLower(attribute.Key) != strings.ToLower(attrKey) {
            return false
        }

        aClasses := strings.Split(attr.Val, " ")
        bClasses := strings.Split(attribute.Val, " ")
        
        sort.Strings(aClasses)
        sort.Strings(bClasses)

        if hasRepetition(bClasses) > 0 { return false }

        if len(aClasses) < len(bClasses) {
            return false
        }

        count := 0
        for _, str := range aClasses {
            for _, str2 := range bClasses {
                if strings.ToLower(str) == strings.ToLower(str2) {
                    count += 1
                }
            }
        }

        if len(aClasses) == len(bClasses) && count == len(aClasses) {
            return true
        }

        // if len(aClasses) > len(bClasses)
        if n := hasRepetition(aClasses); n > 0 && count == len(bClasses) + n {
            return true
        }
        if count == len(bClasses) { return true }

    }
    return false
}

// Checks if []string has repetitive strings and returns the repetition number
func hasRepetition(val []string) int {
    count := 0
    for i:=0; i < len(val); i++ {
        for j:=i+1; j < len(val); j++ {
            if val[i] == val[j] {
                count += 1
            }
        }
    }

    if count > 2 {
        if count == len(val) - (len(val) - count) { return count-1 }
    }
    return count
}
