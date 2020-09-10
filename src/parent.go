package webscraper

// Recursively traverses through parent elements of current element until
// it finds the element that satisfies tag and attrs
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

// Recursively traverses through all parent elements of current element and 
// returns []Element that contains elements that satisfies tag and attrs.
// If limit == -1, then there is no limit.
// If limit == n, it will return only n-number of elements
func (e Element) FindParents(tag string, limit int, attrs ...string) []Element {
    pseudoEl := createPseudoEl(tag, attrs)

    if (e.node == nil) {
        return []Element{}
    }

    var elements []Element
    temp := e.parent()

    for temp != (Element{}) {
        if limit == 0 {
            break
        }

        if temp.compareTo(pseudoEl) == true {
            elements = append(elements, temp)
            limit -= 1
        }

        temp = temp.parent()
    }
    return elements
}
