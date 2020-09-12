# Web-scraper library for Golang
[![Build status](https://travis-ci.org/genjik/web-scraper.svg?branch=master)](https://travis-ci.org/github/genjik/web-scraper)
web-scraper is a small library for go which is built on top of golang.org/x/net/html. It is used to parse and search for html elements

## Installation
The go version has to be with go modules.
Inside the working directory, where the go.mod file is located type in command line:
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
  
**Retrieving raw text from element**
```go
func (e Element) GetText() string
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

**Example**
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
        // Error handling code
    }

    el := root.FindOne("div", true, "id", "special").GetText()
    fmt.Println(el) // Special Message

    elements := root.FindAll("div", true, -1, "class", "list-item") 
    for _, element := range elements {
        fmt.Println(element.GetText()) // List#1-5
    }
}
```
