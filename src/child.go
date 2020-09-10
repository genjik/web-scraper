package webscraper

// Traverses through children elements of current element and returns
// first-found child element that satisfies tag and attributes
// If it doesn't find any element, than it returns nil
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


// Traverses through children elements of current element and appends to
// the []Element any child element that satisfies tag and attributes. 
// If recursive == true, than it will look for children elements of
// children elements, and so on. If limit == -1, then there is no limit. 
// if limit == n, it will return only n-number of elements
func (e Element) FindAll(tag string, recursive bool, limit int, attrs ...string) []Element {
    pseudoEl := createPseudoEl(tag, attrs)

    if (e.node == nil) {
        return []Element{}
    }

    temp := e.firstChild()

    if recursive == true {
        return findAllR(temp, pseudoEl, limit)
    }

    return findAll(temp, pseudoEl, limit)
}

func findAll(e Element, pseudoEl Element, limit int) []Element {
    var elements []Element
    temp := e

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

func findAllR(e Element, pseudoEl Element, limit int) []Element {
    var elements []Element
    temp := e

    for temp != (Element{}) {
        if limit == 0 {
            break
        }

        if temp.compareTo(pseudoEl) == true {
            elements = append(elements, temp)
            limit -= 1
        }

        if temp.firstChild() != (Element{}) {
            found := findAllR(temp.firstChild(), pseudoEl, limit)
            if len(found) > 0 {
                elements = append(elements, found...) 
                limit -= len(found)
            }
        }

        temp = temp.nextSibling()
    }

    return elements
}
