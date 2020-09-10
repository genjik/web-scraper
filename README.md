# Web-scraper library for Golang

## Installation
The go version has to be with go modules.
Inside the working directory, where the go.mod file is located type:
```bash
go get github.com/genjik/web-scraper
```

## Documentation

The `Element` type contains the pointer to the `html.Node`. The whole API uses `Element` type to return html elements.
```go
type Element struct {
    node *html.Node
}
```

This function takes html as any type as long as it implements `io.Reader`, and returns the `Element` type that contains pointer to the <html> node
```go
GetRootElement(r io.Reader) (Element, error)
```

**Searching for children elements**
```go
func (e Element) FindOne(tag string, recursive bool, attrs ...string) Element
func (e Element) FindAll(tag string, recursive bool, limit int, attrs ...string) []Element
```

**Searching for parent elements**
```go
func (e Element) FindParent(tag string, attrs ...string) Element
func (e Element) FindParents(tag string, limit int, attrs ...string) []Element
```

**Searching for sibling elements**
```go
func (e Element) FindPrevSibling(tag string, attrs ...string) Element
func (e Element) FindNextSibling(tag string, attrs ...string) Element
func (e Element) FindPrevSiblings(tag string, limit int, attrs ...string) []Element
func (e Element) FindNextSiblings(tag string, limit int, attrs ...string) []Element
```

**Parameters:**  

`tag string` The tag name of element. E.g html/head/body/div/span/h1 and etc  

`attrs ...string` Contains the attributes of element that method should search for. E.g {"class", "className"}. As many arguments as neccesary can be passed to the parameter, or it can be ommited at all  

`recursive bool` "true" tells the method to look for children elements of children elements and so on. "false" tells to look only for first child element and all of its sibling elements  

`limit int` The number is used to limit the length of final result. -1 means no limit
