package main

import (
    "testing"
    "fmt"
    "golang.org/x/net/html"
    "strings"
)

func TestFindOne(t *testing.T) {
    r := strings.NewReader(`
        <html>
            <head></head>
            <body>
                <div id="red" class="box"></div>
                    <div class="container">
                        <div id="special" class="box"></div>
                    </div>
                <div id="green" class="box"></div>
                <div id="blue" class="box"></div>
            </body>
        </html>
    `)

    root, _ := html.Parse(r) 
    
    elements := []Element{
        {
            &html.Node{
                Type: html.ElementNode,
                Data: "div",
                Attr: []html.Attribute{
                    {Key: "class", Val: "box"},
                    {Key: "id", Val: "red"},
                },
            },
        },
        {
            &html.Node{
                Type: html.ElementNode,
                Data: "div",
                Attr: []html.Attribute{
                    {Key: "class", Val: "container"},
                },
            },
        },
        {
            &html.Node{
                Type: html.ElementNode,
                Data: "div",
                Attr: []html.Attribute{
                    {Key: "id", Val: "special"},
                    {Key: "class", Val: "box"},
                },
            },
        },
        {
            &html.Node{
                Type: html.ElementNode,
                Data: "div",
                Attr: []html.Attribute{
                    {Key: "id", Val: "green"},
                    {Key: "class", Val: "box"},
                },
            },
        },
        {
            &html.Node{
                Type: html.ElementNode,
                Data: "div",
                Attr: []html.Attribute{
                    {Key: "id", Val: "blue"},
                    {Key: "class", Val: "box"},
                },
            },
        },
    }
    
    cases := []struct {
        out Element
        got Element
    }{
        // returns non-nil elements
        {
            elements[0],
            Element{root.FirstChild.LastChild}.
                FindOne("div", false, "id", "red"),
        },
        {
            elements[0],
            Element{root.FirstChild.LastChild}.
                FindOne("div", false, "class", "box"),
        },
        {
            elements[4],
            Element{root.FirstChild.LastChild}.
                FindOne("div", false, "id", "blue"),
        },
        {
            elements[0],
            Element{root.FirstChild.LastChild}.
                FindOne("div", false),
        },
        {
            elements[1],
            Element{root.FirstChild.LastChild}.
                FindOne("div", false, "class", "container"),
        },
        {
            elements[2],
            Element{root.FirstChild.LastChild}.
                FindOne("div", false, "class", "box", "id", "special"),
        },
        {
            elements[0],
            Element{root.FirstChild.LastChild}.
                FindOne("div", false, "class", "box", "id", ""),
        },
        {
            elements[0],
            Element{root.FirstChild.LastChild}.
                FindOne("div", false, "class", "box", "id"),
        },
        // returns nil elements
        {
            Element{},
            Element{root.FirstChild.LastChild}.
                FindOne("div", false, "class", "box", "id", "nope"),
        },
        {
            Element{},
            Element{root.FirstChild.LastChild}.
                FindOne("span", false, "class", "box", "id", "special"),
        },
        {
            Element{},
            Element{root.FirstChild.LastChild}.
                FindOne("div", false, "class", "not-a-box", "id", "special"),
        },
        {
            Element{},
            Element{root.FirstChild.LastChild}.
                FindOne("", false, "class", "box", "id", "special"),
        },
    }

    for i, test := range cases {
        t.Run(fmt.Sprintf("Case #%d\n", i), func(t *testing.T) {
            
        })
    }
}
