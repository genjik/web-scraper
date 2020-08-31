package webscraper

import (
    "golang.org/x/net/html"
    "testing"
    "fmt"
)

func TestContainsClass(t *testing.T) {
    cases := []struct {
        attrs []html.Attribute
        attr html.Attribute
        expectedOut bool
    }{
        // returns true
        {
            []html.Attribute{
                {Namespace: "", Key: "claSs", Val: "red"},
            },
            html.Attribute{Namespace: "", Key: "Class", Val: "red"},
            true,
        },
        {
            []html.Attribute{
                {Namespace: "", Key: "class", Val: "red"},
            },
            html.Attribute{Namespace: "", Key: "class", Val: "red"},
            true,
        },
        {
            []html.Attribute{
                {Namespace: "", Key: "class", Val: "red green"},
            },
            html.Attribute{Namespace: "", Key: "class", Val: "green"},
            true,
        },
        {
            []html.Attribute{
                {Namespace: "", Key: "class", Val: "red green blue"},
            },
            html.Attribute{Namespace: "", Key: "class", Val: "Red green"},
            true,
        },
        {
            []html.Attribute{
                {Namespace: "", Key: "class", Val: "red green blue"},
            },
            html.Attribute{Namespace: "", Key: "class", Val: "red green"},
            true,
        },
        {
            []html.Attribute{
                {Namespace: "", Key: "class", Val: "red blue green"},
            },
            html.Attribute{Namespace: "", Key: "class", Val: "red green"},
            true,
        },
        {
            []html.Attribute{
                {Namespace: "", Key: "class", Val: "red blue green"},
            },
            html.Attribute{Namespace: "", Key: "class", Val: "red green blue"},
            true,
        },
        {
            []html.Attribute{
                {Namespace: "", Key: "class", Val: "red blue green"},
            },
            html.Attribute{Namespace: "", Key: "class", Val: "red blue green"},
            true,
        },
        {
            []html.Attribute{
                {Namespace: "", Key: "id", Val: "red"},
                {Namespace: "", Key: "src", Val: "www.google.com"},
                {Namespace: "", Key: "class", Val: "blue green"},
            },
            html.Attribute{Namespace: "", Key: "class", Val: "blue"},
            true,
        },
        {
            []html.Attribute{
                {Namespace: "", Key: "class", Val: "red green red green"},
            },
            html.Attribute{Namespace: "", Key: "class", Val: "red green"},
            true,
        },
        {
            []html.Attribute{
                {Namespace: "", Key: "class", Val: "red green red"},
            },
            html.Attribute{Namespace: "", Key: "class", Val: "red green"},
            true,
        },
        {
            []html.Attribute{
                {Namespace: "", Key: "class", Val: "red green red"},
            },
            html.Attribute{Namespace: "", Key: "class", Val: "red"},
            true,
        },
        {
            []html.Attribute{
                {Namespace: "", Key: "class", Val: "red red red"},
            },
            html.Attribute{Namespace: "", Key: "class", Val: "red"},
            true,
        },
        {
            []html.Attribute{
                {Namespace: "", Key: "class", Val: "red green blue red red"},
            },
            html.Attribute{Namespace: "", Key: "class", Val: "red"},
            true,
        },
        {
            []html.Attribute{
                {Namespace: "", Key: "id", Val: "red"},
                {Namespace: "", Key: "src", Val: "www.google.com"},
                {Namespace: "", Key: "class", Val: "blue green"},
            },
            html.Attribute{Namespace: "", Key: "id", Val: "red"},
            true,
        },
        {
            []html.Attribute{
                {Namespace: "", Key: "class", Val: "red"},
            },
            html.Attribute{Namespace: "", Key: "class", Val: "Red"},
            true,
        },
        // returns false
        {
            []html.Attribute{
                {Namespace: "", Key: "id", Val: "red"},
            },
            html.Attribute{Namespace: "", Key: "class", Val: "red"},
            false,
        },
        {
            []html.Attribute{
                {Namespace: "", Key: "class", Val: "red"},
            },
            html.Attribute{Namespace: "", Key: "id", Val: "red"},
            false,
        },
        {
            []html.Attribute{
                {Namespace: "", Key: "class", Val: "red red red"},
            },
            html.Attribute{Namespace: "", Key: "class", Val: "purple"},
            false,
        },
        {
            []html.Attribute{
                {Namespace: "", Key: "class", Val: "red blue green"},
            },
            html.Attribute{Namespace: "", Key: "class", Val: "purple"},
            false,
        },
        {
            []html.Attribute{
                {Namespace: "", Key: "class", Val: "blue"},
            },
            html.Attribute{Namespace: "", Key: "class", Val: "red green"},
            false,
        },
        {
            []html.Attribute{
                {Namespace: "", Key: "class", Val: "red"},
            },
            html.Attribute{Namespace: "", Key: "class", Val: "red green"},
            false,
        },
        {
            []html.Attribute{
                {Namespace: "", Key: "class", Val: "red"},
            },
            html.Attribute{Namespace: "", Key: "class", Val: "green"},
            false,
        },
        {
            []html.Attribute{
                {Namespace: "", Key: "class", Val: "red"},
            },
            html.Attribute{Namespace: "", Key: "class", Val: "red red red"},
            false,
        },
        {
            []html.Attribute{
                {Namespace: "", Key: "class", Val: "red green blue"},
            },
            html.Attribute{Namespace: "", Key: "class", Val: "red red"},
            false,
        },
        { //this one is not bug
            []html.Attribute{
                {Namespace: "", Key: "class", Val: "red green blue red"},
            },
            html.Attribute{Namespace: "", Key: "class", Val: "red red"},
            false,
        },
    }

    for i, test := range cases {
        t.Run(fmt.Sprintf("case #%d", i), func(t *testing.T) {
            got := containsSel(test.attrs, test.attr, test.attr.Key); 
            if got != test.expectedOut {
                t.Errorf("got=%t, expected=%t\n", got, test.expectedOut)
            }
        })
    }
}

func TestHasRepetition(t *testing.T) {
    cases := []struct {
        val []string
        expectedOut int
    }{
        {
            []string{"red", "green", "blue"},
            0,
        },
        {
            []string{"red", "green", "blue", "red", "red"},
            2,
        },
        {
            []string{"red", "green", "blue", "red", "red", "green"},
            3,
        },
        {
            []string{"red", "green", "red"},
            1,
        },
        {
            []string{"red", "red", "red"},
            2,
        },
        {
            []string{"red", "green", "red", "green"},
            2,
        },
        {
            []string{"red", "green", "blue", "red"},
            1,
        },
        {
            []string{"red"},
            0,
        },
        {
            []string{},
            0,
        },
    }

    for i, test := range cases {
        t.Run(fmt.Sprintf("case #%d", i), func(t *testing.T) {
            if got := hasRepetition(test.val); got != test.expectedOut {
                t.Errorf("got=%d, expected=%d\n", got, test.expectedOut)
            }
        })
    }
}

func TestCompareTypeAndData(t *testing.T) {
    cases := []struct {
        e Element
        e2 Element
        expectedOut bool
    }{
        {
            Element{
                &html.Node{
                    Type: html.ElementNode,
                    Data: "div",
                },
            },
            Element{
                &html.Node{
                    Type: html.ElementNode,
                    Data: "div",
                },
            },
            true,
        },
        {
            Element{
                &html.Node{
                    Type: html.ElementNode,
                    Data: "div",
                },
            },
            Element{
                &html.Node{
                    Type: html.ElementNode,
                    Data: "DIV",
                },
            },
            true,
        },
        {
            Element{
                &html.Node{
                    Type: html.ElementNode,
                    Data: "div",
                },
            },
            Element{
                &html.Node{
                    Type: html.ElementNode,
                    Data: "span",
                },
            },
            false,
        },
        {
            Element{
                &html.Node{
                    Type: html.ElementNode,
                    Data: "div",
                },
            },
            Element{
                &html.Node{
                    Type: html.TextNode,
                    Data: "div",
                },
            },
            false,
        },
    }

    for i, test := range cases {
        t.Run(fmt.Sprintf("case #%d", i), func(t *testing.T) {
            if got := compareTypeAndData(test.e, test.e2); got != test.expectedOut {
                t.Errorf("got=%t, expected=%t\n", got, test.expectedOut)
            }
        })
    }
}
