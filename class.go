package main

import (
    "golang.org/x/net/html"
    "strings"
    "sort"
)

// makes sure the len(slice)=even and there is no recurrence among [even] index
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

func (e Element) contains(tag string, attrs []string) bool {
    if e.node.Type != html.ElementNode {
        return false
    }

    if strings.ToLower(e.node.Data) != strings.ToLower(tag) {
        return false
    }

    if (len(attrs) / 2) > len(e.node.Attr) {
        return false
    }

    if len(e.node.Attr) < 1 && len(attrs) > 1 {
        return false
    }

    count := 0
    for i:=0; i < len(attrs); i++ {
        if i % 2 != 0 {
            continue
        }

        for j:=0; j < len(e.node.Attr); j++ {
            // if the attrs[i] == class, then check it differently
            if attrs[i] == "class" {
                if e.node.Attr[j].Key != "class" {
                    continue
                }

                res := containsClass(e.node.Attr[j].Val, attrs[i+1])
                if res == true {
                    count += 1
                    break
                }
                continue
            }

            if attrs[i] == e.node.Attr[j].Key && attrs[i+1] == e.node.Attr[j].Val {
                count += 1 
                break
            }
        }
    }

    if count == (len(attrs) / 2) {
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
