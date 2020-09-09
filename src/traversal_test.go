package webscraper

import (
    "golang.org/x/net/html"
    "testing"
    "strings"
    "fmt"
)

func TestParent(t *testing.T) {
    r := strings.NewReader(`
        <html>
            <head></head>
            <body>
                <div class="box">
                    <div class="red"></div>
                    <div class="green"></div>
                </div>
                <div id="special" class="box"></div>
            </body>
        </html>
    `)

    root, _ := GetRootElement(r)

    cases := []struct{
        out Element
        got Element
    }{
        {
            Element{
                &html.Node{
                    Data: "div",
                    Type: html.ElementNode,
                    Attr: []html.Attribute{
                        {Key: "class", Val: "box"},
                    },
                },
            },
            root.FindOne("div", true, "class", "green").parent(),
        },
        {
            Element{
                &html.Node{
                    Data: "body",
                    Type: html.ElementNode,
                },
            },
            root.FindOne("div", true, "class", "box").parent(),
        },
        {
            Element{
                &html.Node{
                    Data: "html",
                    Type: html.ElementNode,
                },
            },
            root.FindOne("body", false).parent(), 
        },
        {
            Element{},
            root.parent(),
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

func TestFirstChild(t *testing.T) {
    r := strings.NewReader(`
        <html>
            <head></head>
            <body>
                <div class="box">
                    <div class="red">
                        <div>
                            <span id="secret">Hello World!</span>
                        </div>
                    </div>
                    <div class="green"></div>
                </div>
                <div id="special" class="box"></div>
            </body>
        </html>
    `)

    root, _ := GetRootElement(r) 

    cases := []struct{
        out Element
        got Element
    }{
        {
            Element{
                &html.Node{
                    Data: "div",
                    Type: html.ElementNode,
                    Attr: []html.Attribute{
                        {Key: "class", Val: "box"},
                    },
                },
            },
            root.firstChild().nextSibling().firstChild(),
        },
        {
            Element{
                &html.Node{
                    Data: "div",
                    Type: html.ElementNode,
                    Attr: []html.Attribute{
                        {Key: "class", Val: "red"},
                    },
                },
            },
            root.firstChild().nextSibling().firstChild().firstChild(),
        },
        {
            Element{
                &html.Node{
                    Data: "span",
                    Type: html.ElementNode,
                    Attr: []html.Attribute{
                        {Key: "id", Val: "secret"},
                    },
                },
            },
            root.firstChild().nextSibling().
                firstChild().
                    firstChild().
                        firstChild().
                            firstChild(),
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

func TestPrevSibling(t *testing.T) {
    r := strings.NewReader(`
        <html>
            <head></head>
            <body>
                <div class="box">
                    <div class="red"></div>
                    <div class="green"></div>
                </div>
                <div id="special" class="box"></div>
            </body>
        </html>
    `)

    root, _ := GetRootElement(r)

    cases := []struct{
        out Element
        got Element
    }{
        {
            Element{
                &html.Node{
                    Data: "div",
                    Type: html.ElementNode,
                    Attr: []html.Attribute{
                        {Key: "class", Val: "box"},
                    },
                },
            },
            root.FindOne("div", true, "class", "box", "id", "special").
                prevSibling(),
        },
        {
            Element{
                &html.Node{
                    Data: "div",
                    Type: html.ElementNode,
                    Attr: []html.Attribute{
                        {Key: "class", Val: "red"},
                    },
                },
            },
            root.FindOne("div", true, "class", "green").
                prevSibling(),
        },
        {
            Element{
                &html.Node{
                    Data: "head",
                    Type: html.ElementNode,
                },
            },
            root.FindOne("body", false).
                prevSibling(),
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

func TestNextSibling(t *testing.T) {
    r := strings.NewReader(`
        <html>
            <head></head>
            <body>
                <div class="box">
                    <div class="red"></div>
                    <div class="green"></div>
                </div>
                <div id="special" class="box"></div>
            </body>
        </html>
    `)

    root, _ := html.Parse(r)

    cases := []struct{
        out Element
        got Element
    }{
        {
            Element{
                &html.Node{
                    Data: "div",
                    Type: html.ElementNode,
                    Attr: []html.Attribute{
                        {Key: "id", Val: "special"},
                        {Key: "class", Val: "box"},
                    },
                },
            },
            Element{root.FirstChild.LastChild}.firstChild().nextSibling(),
        },
        {
            Element{
                &html.Node{
                    Data: "div",
                    Type: html.ElementNode,
                    Attr: []html.Attribute{
                        {Key: "class", Val: "green"},
                    },
                },
            },
            Element{root.FirstChild.LastChild}.
                firstChild().
                    firstChild().
                        nextSibling(),
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
