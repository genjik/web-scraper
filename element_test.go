package webscraper

import (
    "golang.org/x/net/html"
    "testing"
    "fmt"
)

func TestCreatePseudoEl(t *testing.T) {
    cases := []struct{
        tag string
        attrs []string
        out Element
    }{
        {
            "div",
            []string{},
            Element{
                &html.Node{
                    Data: "div",
                    Type: html.ElementNode,
                    Attr: []html.Attribute{},
                },
            },
        },
        {
            "div",
            []string{"class", "red"},
            Element{
                &html.Node{
                    Data: "div",
                    Type: html.ElementNode,
                    Attr: []html.Attribute{
                        {Key: "class", Val: "red"},
                    },
                },
            },
        },
        {
            "div",
            []string{"class", "red", "id", "special"},
            Element{
                &html.Node{
                    Data: "div",
                    Type: html.ElementNode,
                    Attr: []html.Attribute{
                        {Key: "class", Val: "red"},
                        {Key: "id", Val: "special"},
                    },
                },
            },
        },
        {
            "div",
            []string{"class", "red", "id"},
            Element{
                &html.Node{
                    Data: "div",
                    Type: html.ElementNode,
                    Attr: []html.Attribute{
                        {Key: "class", Val: "red"},
                    },
                },
            },
        },
        {
            "div",
            []string{"class", "red", "class", "green"},
            Element{
                &html.Node{
                    Data: "div",
                    Type: html.ElementNode,
                    Attr: []html.Attribute{
                        {Key: "class", Val: "red"},
                    },
                },
            },
        },
        {
            "div",
            []string{"class", "red", "id", "special", "class", "green"},
            Element{
                &html.Node{
                    Data: "div",
                    Type: html.ElementNode,
                    Attr: []html.Attribute{
                        {Key: "class", Val: "red"},
                        {Key: "id", Val: "special"},
                    },
                },
            },
        },
        {
            "div",
            []string{"class", "red", "id", "special", "src", "www.com"},
            Element{
                &html.Node{
                    Data: "div",
                    Type: html.ElementNode,
                    Attr: []html.Attribute{
                        {Key: "class", Val: "red"},
                        {Key: "id", Val: "special"},
                        {Key: "src", Val: "www.com"},
                    },
                },
            },
        },
        {
            "",
            []string{"class", "red", "id", "special", "class", "green"},
            Element{},
        },
    }

    for i, test := range cases {
        t.Run(fmt.Sprintf("Case #%d\n", i), func(t *testing.T) {
            got := createPseudoEl(test.tag, test.attrs)

            if test.out.compareTo(got) == false {
                t.Errorf("expected=%+v, got=%+v\n", test.out, got) 
            }
        })
    }
}

func TestValidateAttrs(t *testing.T) {
    cases := []struct{
        a []string
        out []html.Attribute
    }{
        {
            []string{},
            []html.Attribute{},
        },
        {
            []string{"class", "red"},
            []html.Attribute{
                {Key: "class", Val: "red"},
            },
        },
        {
            []string{"class", "red", "id", "special"},
            []html.Attribute{
                {Key: "class", Val: "red"},
                {Key: "id", Val: "special"},
            },
        },
        {
            []string{"class", "red", "id", "special", "src", "www.com"},
            []html.Attribute{
                {Key: "class", Val: "red"},
                {Key: "id", Val: "special"},
                {Key: "src", Val: "www.com"},
            },
        },
        {
            []string{"class", "red", "id"},
            []html.Attribute{
                {Key: "class", Val: "red"},
            },
        },
        {
            []string{"class", "red", "class", "green"},
            []html.Attribute{
                {Key: "class", Val: "red"},
            },
        },
        {
            []string{"class", "red", "id", "special", "class", "green"},
            []html.Attribute{
                {Key: "class", Val: "red"},
                {Key: "id", Val: "special"},
            },
        },
    }

    for i, test := range cases {
        t.Run(fmt.Sprintf("Case #%d\n", i), func(t *testing.T) {
            got := validateAttrs(test.a)

            if compareAttrs(test.out, got) == false {
                t.Errorf("expected=%+v, got=%+v\n", test.out, got) 
            }
        })
    }
}

func TestCompareTo(t *testing.T) {
    elements := []Element{
        {
            &html.Node{
                Type: html.ElementNode,
                Data: "div",
                Attr: []html.Attribute{
                    {Key: "class", Val: "red"},
                    {Key: "href",  Val: "link"},
                    {Key: "src", Val: "link2"},
                },
            },
        },
        {
            &html.Node{
                Type: html.DoctypeNode,
                Data: "div",
                Attr: []html.Attribute{
                    {Key: "class", Val: "red"},
                    {Key: "href",  Val: "link"},
                    {Key: "src", Val: "link2"},
                },
            },
        },
        {
            &html.Node{
                Type: html.ElementNode,
                Data: "div",
                Attr: []html.Attribute{
                    {Key: "class", Val: "red green"},
                    {Key: "href",  Val: "link"},
                    {Key: "src", Val: "link2"},
                },
            },
        },
    }

    cases := []struct {
        el Element
        el2 Element
        out bool
    }{
        {
            elements[0],
            createPseudoEl("div", []string{"class", "red"}),
            true,
        },
        {
            elements[0],
            createPseudoEl("div", []string{"class", "red", "src", "link2"}),
            true,
        },
        {
            elements[0],
            createPseudoEl("div", []string{"class", "red", "id"}),
            true,
        },
        {
            elements[0],
            createPseudoEl("div", []string{"class", "red", "class"}),
            true,
        },
        {
            elements[0],
            createPseudoEl("div", []string{"class", "red", "class", "green"}),
            true,
        },
        {
            elements[0],
            createPseudoEl("div", []string{"href", "link", "class", "red", "src", "link2"}),
            true,
        },
        {
            Element{},
            Element{},
            true,
        },
        {
            elements[0],
            createPseudoEl("div", []string{"class", "yellow"}),
            false,
        },
        {
            elements[0],
            createPseudoEl("span", []string{"class", "red"}),
            false,
        },
        {
            elements[0],
            createPseudoEl("div", []string{"class", "red", "href", "somewhere", "src", "also", "a", "b"}),
            false,
        },
        {
            elements[0],
            createPseudoEl("div", []string{"cl", "r", "hr", "s", "src", "olso", "a", "b"}),
            false,
        },
        {
            elements[0],
            createPseudoEl("div", []string{"class", "red", "href", "somewhere", "src", "olso"}),
            false,
        },
        {
            elements[1],
            createPseudoEl("div", []string{"class", "red"}),
            false,
        },
        {
            elements[0],
            createPseudoEl("div", []string{"class", "red", "href", "nolink"}),
            false,
        },
        {
            Element{},
            createPseudoEl("div", []string{"class", "red", "href", "nolink"}),
            false,
        },
    }

    for i, test := range cases {
        t.Run(fmt.Sprintf("Case #%d\n", i), func(t *testing.T) {
            got := test.el.compareTo(test.el2)

            if got != test.out {
                t.Errorf("expected=%t, got=%t\n", test.out, got) 
            }
        })
    }
}

func TestCompareAttrs(t *testing.T) {
    cases := []struct{
        attrs []html.Attribute    
        attrs2 []html.Attribute
        out bool
    }{
        {
            []html.Attribute{
                {Key: "class", Val: "red"},
            },
            []html.Attribute{
                {Key: "class", Val: "red"},
            },
            true,
        },
        {
            []html.Attribute{
                {Key: "class", Val: "red"},
                {Key: "id", Val: "special-id"},
                {Key: "src", Val: "link"},
            },
            []html.Attribute{
                {Key: "class", Val: "red"},
            },
            true,
        },
        {
            []html.Attribute{
                {Key: "class", Val: "red"},
                {Key: "id", Val: "special-id"},
                {Key: "src", Val: "link"},
            },
            []html.Attribute{
                {Key: "class", Val: "red"},
                {Key: "id", Val: "special-id"},
            },
            true,
        },
        {
            []html.Attribute{
                {Key: "claSs", Val: "red"},
            },
            []html.Attribute{
                {Key: "Class", Val: "red"},
            },
            true,
        },
        {
            []html.Attribute{
                {Key: "id", Val: "red"},
            },
            []html.Attribute{
                {Key: "id", Val: "Red"},
            },
            true,
        },
        {
            []html.Attribute{
                {Key: "iD", Val: "red"},
            },
            []html.Attribute{
                {Key: "id", Val: "Red"},
            },
            true,
        },
        {
            []html.Attribute{
                {Key: "class", Val: "red"},
            },
            []html.Attribute{
                {Key: "class", Val: "yellow"},
            },
            false,
        },
        {
            []html.Attribute{
                {Key: "class", Val: "red"},
            },
            []html.Attribute{
                {Key: "class", Val: "yellow"},
                {Key: "id", Val: "yellow"},
            },
            false,
        },
        {
            []html.Attribute{
                {Key: "class", Val: "red"},
            },
            []html.Attribute{
                {Key: "class", Val: "red"},
                {Key: "id", Val: "yellow"},
            },
            false,
        },
        {
            []html.Attribute{
            },
            []html.Attribute{
                {Key: "class", Val: "red"},
                {Key: "id", Val: "yellow"},
            },
            false,
        },
    }

    for i, test := range cases {
        t.Run(fmt.Sprintf("case #%d", i), func(t *testing.T) {
            got := compareAttrs(test.attrs, test.attrs2); 
            if got != test.out {
                t.Errorf("got=%t, expected=%t\n", got, test.out)
            }
        })
    }
}

func TestContainsClass(t *testing.T) {
    cases := []struct{
        a string
        b string
        out bool
    }{
        // returns true
        {
            "red",
            "red",
            true,
        },
        {
            "red green",
            "green",
            true,
        },
        {
            "red green blue",
            "Red green",
            true,
        },
        {
            "red green blue",
            "red green",
            true,
        },
        {
            "red blue green",
            "red green",
            true,
        },
        {
            "red blue green",
            "red green blue",
            true,
        },
        {
            "red blue green",
            "red blue green",
            true,
        },
        {
            "red green red green",
            "red green",
            true,
        },
        {
            "red green red",
            "red green",
            true,
        },
        {
            "red green red",
            "red",
            true,
        },
        {
            "red red red",
            "red",
            true,
        },
        {
            "red green blue red red",
            "red",
            true,
        },
        {
            "red",
            "Red",
            true,
        },
        // returns false
        {
            "red red red",
            "purple",
            false,
        },
        {
            "red blue green",
            "purple",
            false,
        },
        {
            "blue",
            "red green",
            false,
        },
        {
            "red",
            "red green",
            false,
        },
        {
            "red",
            "green",
            false,
        },
        {
            "red",
            "red red red",
            false,
        },
        {
            "red green blue",
            "red red",
            false,
        },
        {
            "red green blue red",
            "red red",
            false,
        },
    }

    for i, test := range cases {
        t.Run(fmt.Sprintf("case #%d", i), func(t *testing.T) {
            got := containsClass(test.a, test.b); 
            if got != test.out {
                t.Errorf("got=%t, expected=%t\n", got, test.out)
            }
        })
    }
}
