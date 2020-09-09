package webscraper

import (
    "golang.org/x/net/html"
    "testing"
    "strings"
    "fmt"
)

func TestFindOne(t *testing.T) {
    r := strings.NewReader(`
        <html>
            <head></head>
            <body>
                <div id="red" class="box">
                    <div class="container">
                        <div id="special" class="box"></div>
                    </div>
                </div>
                <div id="green" class="box">
                    <div>
                        <div class="list-item"></div>
                        <div class="list-item"></div>
                        <div class="list-item"></div>
                        <div class="list-item">
                            <div class="find-me">You did it!</div>
                        </div>
                        <div class="list-item"></div>
                    </div>
                </div>
                <div id="blue" class="box">hello world!</div>
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
        {
            &html.Node{
                Type: html.ElementNode,
                Data: "div",
                Attr: []html.Attribute{
                    {Key: "class", Val: "find-me"},
                },
            },
        },
    }
    
    cases := []struct {
        out Element
        got Element
    }{
        // returns non-nil elements
            // recursive true
        {
            elements[1],
            root.FindOne("div", true, "class", "container"),
        },
        {
            elements[2],
            root.FindOne("div", true, "class", "box", "id", "special"),
        },
        {
            elements[5],
            root.FindOne("div", true, "class", "find-me"),
        },
            // recursive false
        {
            elements[0],
            root.FindOne("body", false).
                FindOne("div", false, "id", "red"),
        },
        {
            elements[0],
            root.FindOne("body", false).
                FindOne("div", false, "class", "box"),
        },
        {
            elements[4],
            root.FindOne("body", false).
                FindOne("div", false, "id", "blue"),
        },
        {
            elements[0],
            root.FindOne("body", false).
                FindOne("div", false),
        },
        {
            elements[0],
            root.FindOne("body", false).
                FindOne("div", false, "class", "box", "id"),
        },
        // returns nil elements
        {
            Element{},
            root.FindOne("body", false).
                FindOne("div", false, "class", "box", "id", "nope"),
        },
        {
            Element{},
            root.FindOne("body", false).
                FindOne("span", false, "class", "box", "id", "special"),
        },
        {
            Element{},
            root.FindOne("body", false).
                FindOne("div", false, "class", "not-a-box", "id", "special"),
        },
        {
            Element{},
            root.FindOne("body", false).
                FindOne("", false, "class", "box", "id", "special"),
        },
        {
            Element{},
            root.FindOne("body", false).
                FindOne("span", false),
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
