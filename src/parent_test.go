package webscraper

import (
    "golang.org/x/net/html"
    "testing"
    "strings"
    "fmt"
)

func TestFindParent(t *testing.T) {
    r := strings.NewReader(`
        <html>
            <head></head>
            <body>
                <div id="red" class="box"></div>
                <div id="green" class="box">
                    <div class="list-item"></div>
                    <div class="list-item"></div>
                    <div class="list-item"></div>
                    <div class="list-item" id="special">
                        <div class="find-me">You did it!</div>
                    </div>
                    <div class="list-item"></div>
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
                    {Key: "id", Val: "special"},
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
                Data: "body",
            },
        },
        {
            &html.Node{
                Type: html.ElementNode,
                Data: "html",
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
            root.FindOne("div", true, "class", "find-me").
                FindParent("div", "class", "list-item"),
        },
        {
            elements[0],
            root.FindOne("div", true, "class", "find-me").
                FindParent("div", "class", "list-item", "id", "special"),
        },
        {
            elements[0],
            root.FindOne("div", true, "class", "find-me").
                FindParent("div"),
        },
        {
            elements[1],
            root.FindOne("div", true, "class", "find-me").
                FindParent("div", "class", "box", "id", "green"),
        },
        {
            elements[1],
            root.FindOne("div", true, "class", "find-me").
                FindParent("div", "id", "green"),
        },
        {
            elements[1],
            root.FindOne("div", true, "class", "list-item").
                FindParent("div", "id", "green"),
        },
        {
            elements[2],
            root.FindOne("div", true, "class", "find-me").
                FindParent("body"),
        },
        {
            elements[3],
            root.FindOne("div", true, "class", "find-me").
                FindParent("html"),
        },
        // returns nil elements
        {
            Element{},
            root.FindOne("body", false).
                FindParent("div", "class", "box", "id", "nope"),
        },
        {
            Element{},
            root.FindOne("body", false).
                FindParent("span", "class", "box", "id", "special"),
        },
        {
            Element{},
            root.FindOne("body", false).
                FindParent("div", "class", "not-a-box", "id", "special"),
        },
        {
            Element{},
            root.FindOne("body", false).
                FindParent("", "class", "box", "id", "special"),
        },
        {
            Element{},
            root.FindOne("body", false).
                FindParent("span"),
        },
        {
            Element{},
            Element{}.FindParent("span"),
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

func TestFindParents(t *testing.T) {
    r := strings.NewReader(`
        <html>
            <head></head>
            <body>
                <div class="box" id="red"></div>

                <div class="box" id="green">
                    <div class="box" id="child-green">
                        <div class="box" id="grand-child-green">
                            <div class="box" id="grand-grand-child-green"></div>
                        </div>
                    </div>
                </div>
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
                    {Key: "id", Val: "grand-child-green"},
                },
            },
        },
        {
            &html.Node{
                Type: html.ElementNode,
                Data: "div",
                Attr: []html.Attribute{
                    {Key: "class", Val: "box"},
                    {Key: "id", Val: "child-green"},
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
                Data: "body",
            },
        },
    }
    
    cases := []struct {
        out []Element
        got []Element
    }{
        // returns non-nil elements
        {
            elements[:3],
            root.FindOne("div", true, "id", "grand-grand-child-green").
                FindParents("div", -1, "class", "box"),
        },
        {
            elements[:1],
            root.FindOne("div", true, "id", "grand-grand-child-green").
                FindParents("div", 1, "class", "box"),
        },
        {
            elements[:1],
            root.FindOne("div", true, "id", "grand-grand-child-green").
                FindParents("div", 1),
        },
        {
            elements[3:4],
            root.FindOne("div", true, "id", "grand-grand-child-green").
                FindParents("body", 1),
        },
        // returns nil elements
        {
            elements[:0],
            root.FindOne("div", true, "id", "grand-grand-child-green").
                FindParents("span", -1),
        },
        {
            elements[:0],
            root.FindOne("div", true, "id", "grand-grand-child-green").
                FindParents("span", 0, "class", "box"),
        },
        {
            elements[:0],
            Element{}.FindParents("span", 0, "class", "box"),
        },
    }

    for i, test := range cases {
        t.Run(fmt.Sprintf("Case #%d\n", i), func(t *testing.T) {
            if len(test.out) != len(test.got) {
                t.Fatalf("len(test.out)=%d, len(test.got)=%d\n", len(test.out), len(test.got))
            }

            for j, el := range test.out {
                t.Run(fmt.Sprintf("el #%d\n", j), func(t *testing.T) {
                    if el.compareTo(test.got[j]) == false {
                        t.Errorf("expected=%+v, got=%+v\n", test.out[j], test.got[j]) 
                    }
                })
            }
        })
    }
}
