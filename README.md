# Web-scraper library for Golang
[![Build status](https://travis-ci.org/genjik/web-scraper.svg?branch=master)](https://travis-ci.org/github/genjik/web-scraper)

web-scraper is a small library for parsing and scraping the Html. It is built on top of golang.org/x/net/html

## Installation
The go version has to be with go modules.
Type the following command inside the working directory where the go.mod file is:
```bash
go get github.com/genjik/web-scraper
```

## Documentation
The `Element` type contains a pointer to the `html.Node`. The whole API uses `Element` type to return html elements.
```go
type Element struct {
    node *html.Node
}
```

`GetRootElement` takes Html as any type as long as it satisfies `io.Reader`. The function returns the `Element` type that contains pointer to the `<html>` node
```go
GetRootElement(r io.Reader) (Element, error)
```
  
**Retrieve raw text from element**
```go
func (e Element) GetText() string
```

**Search for child elements**
```go
func (e Element) FindOne(tag string, recursive bool, attrs ...string) Element
func (e Element) FindAll(tag string, recursive bool, limit int, attrs ...string) []Element
```

**Search for parent elements**
```go
func (e Element) FindParent(tag string, attrs ...string) Element
func (e Element) FindParents(tag string, limit int, attrs ...string) []Element
```

**Search for sibling elements**
```go
func (e Element) FindPrevSibling(tag string, attrs ...string) Element
func (e Element) FindNextSibling(tag string, attrs ...string) Element
func (e Element) FindPrevSiblings(tag string, limit int, attrs ...string) []Element
func (e Element) FindNextSiblings(tag string, limit int, attrs ...string) []Element
```

**Get an element**
```go
func (e Element) Parent() Element // Returns parent element
func (e Element) FirstChild() Element // Not supported yet
func (e Element) PrevSibling() Element // Not supported yet
func (e Element) NextSibling() Element // Not supported yet
```

**Parameters:**  
`tag string` The tag name of element. E.g html/head/body/div/span/h1 and so on.  

`attrs ...string` The attributes of element the method will search for. E.g {"class", "className"}. As many arguments as neccesary can be passed to the parameter, or it can be ommited at all  

`recursive bool` "false" tells a method to look only for the elements that are children for the current element. "true" tells the method to look for child elements until it reaches the last element of `html` tree.

`limit int` The number is used to limit the size of final result. -1 means no limit

## Example
```go
package main

import (
    "strings"
    "github.com/genjik/web-scraper"
    "fmt"
)

func main() {
    r := strings.NewReader(`
        <html>
            <head></head>
            <body>
                <div id="red" class="box">
                    <div id="special">Special Message</div> 
                </div>

                <div id="green" class="box">
                    <div>
                        <div class="list-item" id="l1">List#1</div>
                        <div class="list-item" id="l2">List#2</div>
                        <div class="list-item" id="l3">List#3</div>
                        <div class="list-item" id="l4">List#4</div>
                        <div class="list-item" id="l5">List#5</div>
                    </div>
                </div>
            </body>
        </html>
    `)

    root, err := webscraper.GetRootElement(r)
    if err != nil {
        // Error handling
    }

    el := root.FindOne("div", true, "id", "special").GetText()
    fmt.Println(el) // Special Message

    elements := root.FindAll("div", true, -1, "class", "list-item") 
    for _, element := range elements {
        fmt.Println(element.GetText()) // List#1-5
    }
}
```
