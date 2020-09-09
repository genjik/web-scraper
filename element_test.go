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
    }

    for i, test := range cases {
        t.Run(fmt.Sprintf("Case #%d\n", i), func(t *testing.T) {
            if test.out.compareTo(test.got) == false {
                t.Errorf("expected=%+v, got=%+v\n", test.out, test.got) 
            }
        })
    }
}

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
