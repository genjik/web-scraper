package webscraper

func (e Element) FindOne(tag string, recursive bool, attrs ...string) Element {
    pseudoEl := createPseudoEl(tag, attrs)

    if (e.node == nil) {
        return Element{}
    }

    temp := e.firstChild()

    if recursive == true {
        return findOneR(temp, pseudoEl)
    }

    return findOne(temp, pseudoEl)
}

func findOne(e Element, pseudoEl Element) Element {
    temp := e

    for temp != (Element{}) {
        if temp.compareTo(pseudoEl) == true {
            return temp
        }

        temp = temp.nextSibling()
    }
    return Element{}
}

func findOneR(e Element, pseudoEl Element) Element {
    temp := e

    for temp != (Element{}) {
        if temp.compareTo(pseudoEl) == true {
            return temp
        }

        if temp.firstChild() != (Element{}) {
            found := findOneR(temp.firstChild(), pseudoEl)
            if found != (Element{}) {
                return found
            }
        }

        temp = temp.nextSibling()
    }
    return Element{}
}

//func (e Element) FindAll(sel selector, limit int, recursive bool) []Element {}
