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

//func (e Element) FindNextSiblings(selector, limit int) []Element{}

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

//func (e Element) FindPrevSiblings(selector, limit int) []Element {}
