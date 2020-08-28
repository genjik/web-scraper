package main // change it

import (
    "testing"
    //"io"
    //"strings"
    "golang.org/x/net/html"
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
        // returns false
        {
            []html.Attribute{
                {Namespace: "", Key: "class", Val: "red"},
            },
            html.Attribute{Namespace: "", Key: "class", Val: "Red"},
            false,
        },
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
                {Namespace: "", Key: "id", Val: "red"},
                {Namespace: "", Key: "src", Val: "www.google.com"},
                {Namespace: "", Key: "class", Val: "blue green"},
            },
            html.Attribute{Namespace: "", Key: "id", Val: "red"},
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
                {Namespace: "", Key: "class", Val: "red green blue red"},
            },
            html.Attribute{Namespace: "", Key: "class", Val: "red red"},
            false,
        },
        {
            []html.Attribute{
                {Namespace: "", Key: "class", Val: "red green blue"},
            },
            html.Attribute{Namespace: "", Key: "class", Val: "red red"},
            false,
        },
        // temp bugs, need to fix them
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
    }

    for i, test := range cases {
        if got := containsClass(test.attrs, test.attr); got != test.expectedOut {
            t.Errorf("%d) got=%t, expected=%t\n", i+1, got, test.expectedOut)
        }
    }
}

func TestHasRepetition(t *testing.T) {
    cases := []struct {
        val []string
        expectedOut int
    }{
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
    }
    for i, test := range cases {
        if got := hasRepetition(test.val); got != test.expectedOut {
            t.Errorf("%d) got=%d, expected=%d\n", i+1, got, test.expectedOut)
        }
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
        if got := test.e.compareTypeAndData(test.e2); got != test.expectedOut {
            t.Errorf("%d) got=%t, expected=%t\n", i+1, got, test.expectedOut)
        }
    }
}

//func TestGetRootElement(t *testing.T) {
//    cases := []struct {
//        html io.Reader
//        expectedErr bool
//        expectedOut Element
//    }{
//        {
//            strings.NewReader("<html><head></head></html>"),
//            false,
//            Element{
//                &html.Node{
//                    Type: html.ElementNode,
//                    Data: "html",
//                },
//            },
//        },
//        {
//            strings.NewReader(""),
//            false,
//            Element{
//                &html.Node{
//                    Type: html.ElementNode,
//                    Data: "html",
//                },
//            },
//        },
//    }
//
//    for i, test := range cases {
//        got, err := GetRootElement(test.html)
//
//        if (err == nil && test.expectedErr == true) || 
//        (err != nil && test.expectedErr == false) {
//            t.Errorf("%d) error = %v, expectedErr = %t", i, err,
//                test.expectedErr)
//        }
//
//        if got.compareTypeAndData(test.expectedOut) != true {
//            t.Errorf("%d) got=%+v, expected=%+v\n", i+1, got, test.expectedOut)
//        }
//    }
//}
