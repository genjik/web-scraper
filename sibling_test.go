package webscraper

import (
    "testing"
    "fmt"
    "golang.org/x/net/html"
    "strings"
)

func TestFindNextSibling(t *testing.T) {
    r := strings.NewReader(`
        <html>
            <head></head>
            <body>
                <div id="red" class="box"></div>
                <div id="green" class="box">
                    <div class="list-item" id="1l"></div>
                    <div class="list-item" id="2l"></div>
                    <div class="list-item" id="3l"></div>
                    <div class="list-item" id="4l"></div>
                    <div class="list-item" id="5l"></div>
                </div>
                <div id="blue" class="box"></div>
            </body>
        </html>
    `)

    root, _ := GetRootElement(r)
    
    elements := []Element{
        {
            &html.Node{
                Type: html.ElementNode,
                Data: "div",
                Attr: []html.Attribute{
                    {Key: "class", Val: "list-item"},
                    {Key: "id", Val: "2l"},
                },
            },
        },
        {
            &html.Node{
                Type: html.ElementNode,
                Data: "div",
                Attr: []html.Attribute{
                    {Key: "class", Val: "list-item"},
                    {Key: "id", Val: "5l"},
                },
            },
        },
        {
            &html.Node{
                Type: html.ElementNode,
                Data: "div",
                Attr: []html.Attribute{
                    {Key: "class", Val: "box"},
                    {Key: "id", Val: "green"},
                },
            },
        },
        {
            &html.Node{
                Type: html.ElementNode,
                Data: "div",
                Attr: []html.Attribute{
                    {Key: "class", Val: "box"},
                    {Key: "id", Val: "blue"},
                },
            },
        },
        {
            &html.Node{
                Type: html.ElementNode,
                Data: "body",
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
            root.FindOne("div", true, "class", "list-item", "id", "1l").
                FindNextSibling("div"),
        },
        {
            elements[0],
            root.FindOne("div", true, "class", "list-item", "id", "1l").
                FindNextSibling("div", "id", "2l"),
        },
        {
            elements[1],
            root.FindOne("div", true, "class", "list-item", "id", "1l").
                FindNextSibling("div", "id", "5l"),
        },
        {
            elements[2],
            root.FindOne("div", true, "class", "box", "id", "red").
                FindNextSibling("div"),
        },
        {
            elements[2],
            root.FindOne("div", true, "class", "box", "id", "red").
                FindNextSibling("div", "id", "green"),
        },
        {
            elements[3],
            root.FindOne("div", true, "class", "box", "id", "red").
                FindNextSibling("div", "id", "blue"),
        },
        {
            elements[4],
            root.FindOne("head", false).
                FindNextSibling("body"), 
        },
        // returns nil elements
        {
            Element{},
            root.FindOne("body", false).
                FindNextSibling("body"), 
        },
        {
            Element{},
            root.FindOne("div", true, "class", "box", "id", "blue").
                FindNextSibling("div", "id", "green"),
        },
    }

    for i, test := range cases {
        t.Run(fmt.Sprintf("Case #%d\n", i), func(t *testing.T) {
            if test.out.compareTo(test.got) == false {
                t.Errorf("expected=%+v, got=%+v\n", test.out, test.got) 
            }
        })
    }
}

func TestFindPrevSibling(t *testing.T) {
    r := strings.NewReader(`
        <html>
            <head></head>
            <body>
                <div id="red" class="box"></div>
                <div id="green" class="box">
                    <div class="list-item" id="1l"></div>
                    <div class="list-item" id="2l"></div>
                    <div class="list-item" id="3l"></div>
                    <div class="list-item" id="4l"></div>
                    <div class="list-item" id="5l"></div>
                </div>
                <div id="blue" class="box"></div>
            </body>
        </html>
    `)

    root, _ := GetRootElement(r)
    
    elements := []Element{
        {
            &html.Node{
                Type: html.ElementNode,
                Data: "div",
                Attr: []html.Attribute{
                    {Key: "class", Val: "list-item"},
                    {Key: "id", Val: "4l"},
                },
            },
        },
        {
            &html.Node{
                Type: html.ElementNode,
                Data: "div",
                Attr: []html.Attribute{
                    {Key: "class", Val: "list-item"},
                    {Key: "id", Val: "1l"},
                },
            },
        },
        {
            &html.Node{
                Type: html.ElementNode,
                Data: "div",
                Attr: []html.Attribute{
                    {Key: "class", Val: "box"},
                    {Key: "id", Val: "green"},
                },
            },
        },
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
                Data: "head",
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
            root.FindOne("div", true, "class", "list-item", "id", "5l").
                FindPrevSibling("div"),
        },
        {
            elements[0],
            root.FindOne("div", true, "class", "list-item", "id", "5l").
                FindPrevSibling("div", "id", "4l"),
        },
        {
            elements[1],
            root.FindOne("div", true, "class", "list-item", "id", "5l").
                FindPrevSibling("div", "id", "1l"),
        },
        {
            elements[2],
            root.FindOne("div", true, "class", "box", "id", "blue").
                FindPrevSibling("div"),
        },
        {
            elements[2],
            root.FindOne("div", true, "class", "box", "id", "blue").
                FindPrevSibling("div", "id", "green"),
        },
        {
            elements[3],
            root.FindOne("div", true, "class", "box", "id", "blue").
                FindPrevSibling("div", "id", "red"),
        },
        {
            elements[4],
            root.FindOne("body", false).
                FindPrevSibling("head"), 
        },
        // returns nil elements
        {
            Element{},
            root.FindOne("head", false).
                FindPrevSibling("head"), 
        },
        {
            Element{},
            root.FindOne("div", true, "class", "box", "id", "red").
                FindPrevSibling("div", "id", "green"),
        },
    }

    for i, test := range cases {
        t.Run(fmt.Sprintf("Case #%d\n", i), func(t *testing.T) {
            if test.out.compareTo(test.got) == false {
                t.Errorf("expected=%+v, got=%+v\n", test.out, test.got) 
            }
        })
    }
}
