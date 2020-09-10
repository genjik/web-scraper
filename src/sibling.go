package webscraper

// Traverses through sibling elements BEFORE current element,
// and returns element if it satisfies the searching parameters.
// Otherwise, returns nil 
func (e Element) FindPrevSibling(tag string, attrs ...string) Element {
    return findElement(e, prevSibling, tag, attrs)
}

// Traverses through sibling elements AFTER current element,
// and returns element if it satisfies the searching parameters.
// Otherwise, returns nil 
func (e Element) FindNextSibling(tag string, attrs ...string) Element {
    return findElement(e, nextSibling, tag, attrs)
}

// Traverses through sibling elements BEFORE current element,
// and returns []Element that contains elements that satisfies the searching
// parameters. Otherwise, returns nil 
func (e Element) FindPrevSiblings(tag string, limit int, attrs ...string) []Element {
    return findElements(e, prevSibling, tag, limit, attrs)
}

// Traverses through sibling elements AFTER current element,
// and returns []Element that contains elements that satisfies the searching
// parameters. Otherwise, returns nil 
func (e Element) FindNextSiblings(tag string, limit int, attrs ...string) []Element {
    return findElements(e, nextSibling, tag, limit, attrs)
}
