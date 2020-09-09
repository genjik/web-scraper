package webscraper

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
