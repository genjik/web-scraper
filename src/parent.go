package webscraper

// Recursively traverses through parent elements of current element until
// it finds the element that satisfies tag and attrs
func (e Element) FindParent(tag string, attrs ...string) Element {
    return e.findElement(getParent, tag, attrs)
}

// Recursively traverses through all parent elements of current element and 
// returns []Element that contains elements that satisfies tag and attrs.
// If limit == -1, then there is no limit.
// If limit == n, it will return only n-number of elements
func (e Element) FindParents(tag string, limit int, attrs ...string) []Element {
    return e.findElements(getParent, tag, limit, attrs)
}
