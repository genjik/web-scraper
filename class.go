package main

import (
    "golang.org/x/net/html"
    "strings"
    "sort"
)

func createPseudoEl(tag string, attrs []string) Element {
    if tag == "" {
        return Element{}
    }

    validated := validateAttrs(attrs)

    el := Element{
        &html.Node{
            Type: html.ElementNode,
            Data: tag,
            Attr: validated,
        },
    }

    return el
}

// makes sure the len(slice)=even and there is no recurrence among [even] index
func validateAttrs(attrs []string) []html.Attribute {
    var newAttrs []html.Attribute 

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
            newAttrs = append(newAttrs, html.Attribute{Key: attrs[i], Val: attrs[i+1]})
        }
    }

    return newAttrs
}

func (e Element) compareTo(e2 Element) bool {
    if e != (Element{}) && e2 == (Element{}) ||
       e == (Element{}) && e2 != (Element{}) {
        return false
    }

    if e == (Element{}) && e2 == (Element{}) {
        return true
    }
    
    if e.node.Type != html.ElementNode || e2.node.Type != html.ElementNode {
        return false
    }

    if strings.ToLower(e.node.Data) != strings.ToLower(e2.node.Data) {
        return false
    }
    
    return compareAttrs(e.node.Attr, e2.node.Attr)
}

// The order matters! the first parameter has to be part of real html document
func compareAttrs(attrs []html.Attribute, attrs2 []html.Attribute) bool {
    if len(attrs2) > len(attrs) {
        return false
    }

    if len(attrs) < 1 && len(attrs2) > 0 {
        return false
    }

    count := 0
    for i:=0; i < len(attrs2); i++ {
        for j:=0; j < len(attrs); j++ {
            if strings.ToLower(attrs2[i].Key) == "class" {
                if strings.ToLower(attrs[j].Key) != "class" {
                    continue
                }

                res := containsClass(attrs[j].Val, attrs2[i].Val)
                if res == true {
                    count += 1
                    break
                }
                continue
            }

            if strings.ToLower(attrs2[i].Key) == strings.ToLower(attrs[j].Key) &&
               strings.ToLower(attrs2[i].Val) == strings.ToLower(attrs[j].Val) {
                count += 1 
                break
            }
        }
    }

    if count == len(attrs2) {
        return true
    }

    return false
}

func containsClass(parsedClasses, newClasses string) bool {
    aClasses := strings.Split(parsedClasses, " ")
    bClasses := strings.Split(newClasses, " ")
                    
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

    return false
}
