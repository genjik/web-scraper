package webscraper

func (e Element) FindNextSibling(tag string, attrs ...string) Element {
    pseudoEl := createPseudoEl(tag, attrs)

    if (e.node == nil) {
        return Element{}
    }

    temp := e.nextSibling()

    for temp != (Element{}) {
        if temp.compareTo(pseudoEl) == true {
            return temp
        }
        temp = temp.nextSibling()
    }

    return Element{}
}

func (e Element) FindNextSiblings(tag string, limit int, attrs ...string) []Element {
    var elements []Element
    pseudoEl := createPseudoEl(tag, attrs)

    if (e.node == nil) {
        return elements
    }

    temp := e.nextSibling()

    for temp != (Element{}) {
        if limit == 0 {
            break
        }

        if temp.compareTo(pseudoEl) == true {
            elements = append(elements, temp)
            limit -= 1
        }

        temp = temp.nextSibling()
    }

    return elements
}

func (e Element) FindPrevSibling(tag string, attrs ...string) Element {
    pseudoEl := createPseudoEl(tag, attrs)

    if (e.node == nil) {
        return Element{}
    }

    temp := e.prevSibling()

    for temp != (Element{}) {
        if temp.compareTo(pseudoEl) == true {
            return temp
        }
        temp = temp.prevSibling()
    }

    return Element{}
}

func (e Element) FindPrevSiblings(tag string, limit int, attrs ...string) []Element {
    var elements []Element
    pseudoEl := createPseudoEl(tag, attrs)

    if (e.node == nil) {
        return elements
    }

    temp := e.prevSibling()

    for temp != (Element{}) {
        if limit == 0 {
            break
        }

        if temp.compareTo(pseudoEl) == true {
            elements = append(elements, temp)
            limit -= 1
        }

        temp = temp.prevSibling()
    }

    return elements
}
