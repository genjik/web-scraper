package webscraper

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
            root.firstChild().nextSibling().
                FindOne("div", true, "class", "container"),
        },
        {
            elements[2],
            root.firstChild().nextSibling().
                FindOne("div", true, "class", "box", "id", "special"),
        },
            // recursive false
        {
            elements[0],
            root.firstChild().nextSibling().
                FindOne("div", false, "id", "red"),
        },
        {
            elements[0],
            root.firstChild().nextSibling().
                FindOne("div", false, "class", "box"),
        },
        {
            elements[4],
            root.firstChild().nextSibling().
                FindOne("div", false, "id", "blue"),
        },
        {
            elements[0],
            root.firstChild().nextSibling().
                FindOne("div", false),
        },
        {
            elements[0],
            root.firstChild().nextSibling().
                FindOne("div", false, "class", "box", "id"),
        },
        {
            elements[5],
            root.firstChild().nextSibling().
                FindOne("div", false, "class", "find-me"),
        },
        // returns nil elements
        {
            Element{},
            root.firstChild().nextSibling().
                FindOne("div", false, "class", "box", "id", "nope"),
        },
        {
            Element{},
            root.firstChild().nextSibling().
                FindOne("span", false, "class", "box", "id", "special"),
        },
        {
            Element{},
            root.firstChild().nextSibling().
                FindOne("div", false, "class", "not-a-box", "id", "special"),
        },
        {
            Element{},
            root.firstChild().nextSibling().
                FindOne("", false, "class", "box", "id", "special"),
        },
        {
            Element{},
            root.firstChild().nextSibling().
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
