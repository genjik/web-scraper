package main

import (
    //"testing"
    "golang.org/x/net/html"
)

var expectedOutput []Element = []Element{
    //FindAllChildren
    {
        &html.Node{
            Type: html.ElementNode,
            Data: "div",
            Attr: []html.Attribute{
                {Namespace: "", Key: "class", Val: "red"},
            },
        },
    },
    {
        &html.Node{
            Type: html.ElementNode,
            Data: "div",
            Attr: []html.Attribute{
                {Namespace: "", Key: "class", Val: "green"},
            },
        },
    },
    {
        &html.Node{
            Type: html.ElementNode,
            Data: "div",
            Attr: []html.Attribute{
                {Namespace: "", Key: "class", Val: "green"},
            },
        },
    },
}

//func TestFindAllChildren(t *testing.T) {
//    cases := []struct{
//        recursive bool
//        limit int
//        expectedOut []Element
//    }{
//        {
//            false,
//            0,
//            expectedOutput[:3],
//        },
//    }
//
//    for _, got := 
//}
