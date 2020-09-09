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

//func (e Element) FindParents(selector, limit int) []Element {}
