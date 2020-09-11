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

func TestFindNextSibling(t *testing.T) {
    r := strings.NewReader(`
        <html>
            <head></head>
            <body>
                <div id="red" class="box"></div>
                <div id="green" class="box">
                    <div class="list-item" id="l1"></div>
                    <div class="list-item" id="l2"></div>
                    <div class="list-item" id="l3"></div>
                    <div class="list-item" id="l4"></div>
                    <div class="list-item" id="l5"></div>
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
                    {Key: "id", Val: "l2"},
                },
            },
        },
        {
            &html.Node{
                Type: html.ElementNode,
                Data: "div",
                Attr: []html.Attribute{
                    {Key: "class", Val: "list-item"},
                    {Key: "id", Val: "l5"},
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
            root.FindOne("div", true, "class", "list-item", "id", "l1").
                FindNextSibling("div"),
        },
        {
            elements[0],
            root.FindOne("div", true, "class", "list-item", "id", "l1").
                FindNextSibling("div", "id", "l2"),
        },
        {
            elements[1],
            root.FindOne("div", true, "class", "list-item", "id", "l1").
                FindNextSibling("div", "id", "l5"),
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

func TestFindNextSiblings(t *testing.T) {
    r := strings.NewReader(`
        <html>
            <head></head>
            <body>
                <div id="red" class="box"></div>
                <div id="green" class="box">
                    <div class="list-item" id="l1"></div>
                    <div class="list-item" id="l2"></div>
                    <div class="list-item" id="l3"></div>
                    <div class="list-item" id="l4"></div>
                    <div class="list-item" id="l5"></div>
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
                    {Key: "id", Val: "l2"},
                },
            },
        },
        {
            &html.Node{
                Type: html.ElementNode,
                Data: "div",
                Attr: []html.Attribute{
                    {Key: "class", Val: "list-item"},
                    {Key: "id", Val: "l3"},
                },
            },
        },
        {
            &html.Node{
                Type: html.ElementNode,
                Data: "div",
                Attr: []html.Attribute{
                    {Key: "class", Val: "list-item"},
                    {Key: "id", Val: "l4"},
                },
            },
        },
        {
            &html.Node{
                Type: html.ElementNode,
                Data: "div",
                Attr: []html.Attribute{
                    {Key: "class", Val: "list-item"},
                    {Key: "id", Val: "l5"},
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
        out []Element
        got []Element
    }{
        // returns non-nil elements
        {
            elements[0:4],
            root.FindOne("div", true, "class", "list-item", "id", "l1").
                FindNextSiblings("div", -1),
        },
        {
            elements[0:2],
            root.FindOne("div", true, "class", "list-item", "id", "l1").
                FindNextSiblings("div", 2),
        },
        {
            elements[0:4],
            root.FindOne("div", true, "class", "list-item", "id", "l1").
                FindNextSiblings("div", -1, "class", "list-item"),
        },
        {
            elements[1:2],
            root.FindOne("div", true, "class", "list-item", "id", "l1").
                FindNextSiblings("div", -1, "id", "l3"),
        },
        {
            elements[4:6],
            root.FindOne("div", true, "class", "box", "id", "red").
                FindNextSiblings("div", -1),
        },
        {
            elements[6:7],
            root.FindOne("head", true).
                FindNextSiblings("body", -1),
        },
        // returns nil elements
        {
            []Element{},
            root.FindOne("body", false).
                FindNextSiblings("body", -1), 
        },
        {
            []Element{},
            root.FindOne("div", true, "class", "box", "id", "blue").
                FindNextSiblings("div", -1, "id", "green"),
        },
        {
            []Element{},
            Element{}.FindNextSiblings("div", -1),
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

func TestFindPrevSibling(t *testing.T) {
    r := strings.NewReader(`
        <html>
            <head></head>
            <body>
                <div id="red" class="box"></div>
                <div id="green" class="box">
                    <div class="list-item" id="l1"></div>
                    <div class="list-item" id="l2"></div>
                    <div class="list-item" id="l3"></div>
                    <div class="list-item" id="l4"></div>
                    <div class="list-item" id="l5"></div>
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
                    {Key: "id", Val: "l4"},
                },
            },
        },
        {
            &html.Node{
                Type: html.ElementNode,
                Data: "div",
                Attr: []html.Attribute{
                    {Key: "class", Val: "list-item"},
                    {Key: "id", Val: "l1"},
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
            root.FindOne("div", true, "class", "list-item", "id", "l5").
                FindPrevSibling("div"),
        },
        {
            elements[0],
            root.FindOne("div", true, "class", "list-item", "id", "l5").
                FindPrevSibling("div", "id", "l4"),
        },
        {
            elements[1],
            root.FindOne("div", true, "class", "list-item", "id", "l5").
                FindPrevSibling("div", "id", "l1"),
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
        {
            Element{},
            Element{}.FindPrevSibling("div", "id", "green"),
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

func TestFindPrevSiblings(t *testing.T) {
    r := strings.NewReader(`
        <html>
            <head></head>
            <body>
                <div id="red" class="box"></div>
                <div id="green" class="box">
                    <div class="list-item" id="l1"></div>
                    <div class="list-item" id="l2"></div>
                    <div class="list-item" id="l3"></div>
                    <div class="list-item" id="l4"></div>
                    <div class="list-item" id="l5"></div>
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
                    {Key: "id", Val: "l4"},
                },
            },
        },
        {
            &html.Node{
                Type: html.ElementNode,
                Data: "div",
                Attr: []html.Attribute{
                    {Key: "class", Val: "list-item"},
                    {Key: "id", Val: "l3"},
                },
            },
        },
        {
            &html.Node{
                Type: html.ElementNode,
                Data: "div",
                Attr: []html.Attribute{
                    {Key: "class", Val: "list-item"},
                    {Key: "id", Val: "l2"},
                },
            },
        },
        {
            &html.Node{
                Type: html.ElementNode,
                Data: "div",
                Attr: []html.Attribute{
                    {Key: "class", Val: "list-item"},
                    {Key: "id", Val: "l1"},
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
        out []Element
        got []Element
    }{
        // returns non-nil elements
        {
            elements[0:4],
            root.FindOne("div", true, "class", "list-item", "id", "l5").
                FindPrevSiblings("div", -1),
        },
        {
            elements[0:2],
            root.FindOne("div", true, "class", "list-item", "id", "l5").
                FindPrevSiblings("div", 2),
        },
        {
            elements[0:4],
            root.FindOne("div", true, "class", "list-item", "id", "l5").
                FindPrevSiblings("div", -1, "class", "list-item"),
        },
        {
            elements[1:2],
            root.FindOne("div", true, "class", "list-item", "id", "l5").
                FindPrevSiblings("div", -1, "id", "l3"),
        },
        {
            elements[4:6],
            root.FindOne("div", true, "class", "box", "id", "blue").
                FindPrevSiblings("div", -1),
        },
        {
            elements[6:7],
            root.FindOne("body", true).
                FindPrevSiblings("head", -1),
        },
        // returns nil elements
        {
            []Element{},
            root.FindOne("head", false).
                FindPrevSiblings("body", -1), 
        },
        {
            []Element{},
            root.FindOne("div", true, "class", "box", "id", "red").
                FindPrevSiblings("div", -1, "id", "green"),
        },
        {
            []Element{},
            Element{}.FindNextSiblings("div", -1),
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
