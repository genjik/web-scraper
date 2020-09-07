package main

import (
    "testing"
    "fmt"
    "golang.org/x/net/html"
)

func TestValidateAttrs(t *testing.T) {
    cases := []struct{
        a []string
        out []string
    }{
        {
            []string{},
            []string{},
        },
        {
            []string{"class", "red"},
            []string{"class", "red"},
        },
        {
            []string{"class", "red", "id", "special"},
            []string{"class", "red", "id", "special"},
        },
        {
            []string{"class", "red", "id"},
            []string{"class", "red"},
        },
        {
            []string{"class", "red", "class", "green"},
            []string{"class", "red"},
        },
        {
            []string{"class", "red", "id", "special", "class", "green"},
            []string{"class", "red", "id", "special"},
        },
        {
            []string{"class", "red", "id", "special", "src", "www.com"},
            []string{"class", "red", "id", "special", "src", "www.com"},
        },
        {
            []string{"class", "red", "id", "special", "class", "green"},
            []string{"class", "red", "id", "special"},
        },
    }

    for i, test := range cases {
        t.Run(fmt.Sprintf("Case #%d\n", i), func(t *testing.T) {
            got := validateAttrs(test.a)

            if compareStr(test.out, got) == false {
                t.Errorf("expected=%+v, got=%+v\n", test.out, got) 
            }
        })
    }
}

func TestContains(t *testing.T) {
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
        tag string
        validatedAttrs[]string
        out bool
    }{
        {
            elements[0],
            "div",
            validateAttrs([]string{"class", "red"}),
            true,
        },
        {
            elements[0],
            "div",
            validateAttrs([]string{"class", "red", "src", "link2"}),
            true,
        },
        {
            elements[0],
            "div",
            validateAttrs([]string{"class", "red", "id"}),
            true,
        },
        {
            elements[0],
            "div",
            validateAttrs([]string{"class", "red", "class"}),
            true,
        },
        {
            elements[0],
            "div",
            validateAttrs([]string{"class", "red", "class", "green"}),
            true,
        },
        {
            elements[0],
            "div",
            validateAttrs([]string{"class", "yellow"}),
            false,
        },
        {
            elements[0],
            "span",
            validateAttrs([]string{"class", "red"}),
            false,
        },
        {
            elements[0],
            "div",
            validateAttrs([]string{"class", "red", "href", "somewhere", "src", "also", "a", "b"}),
            false,
        },
        {
            elements[0],
            "div",
            validateAttrs([]string{"cl", "r", "hr", "s", "src", "olso", "a", "b"}),
            false,
        },
        {
            elements[0],
            "div",
            validateAttrs([]string{"class", "red", "href", "somewhere", "src", "olso"}),
            false,
        },
        {
            elements[1],
            "div",
            validateAttrs([]string{"class", "red"}),
            false,
        },
        {
            elements[0],
            "div",
            validateAttrs([]string{"class", "red", "href", "nolink"}),
            false,
        },
    }

    for i, test := range cases {
        t.Run(fmt.Sprintf("Case #%d\n", i), func(t *testing.T) {
            got := test.el.contains(test.tag, test.validatedAttrs) 

            if got != test.out {
                t.Errorf("expected=%t, got=%t\n", test.out, got) 
            }
        })
    }
}

func TestContainsClass(t *testing.T) {
    cases := []struct {
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
