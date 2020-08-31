package main // change it 

import (
    "testing"
    "strings"
    "fmt"
    "golang.org/x/net/html"
)

var expectedOutputE []Element = []Element{
    {
        &html.Node{
            Type: html.ElementNode,
            Data: "div",
        },
    },
    {
        &html.Node{
            Type: html.ElementNode,
            Data: "span",
        },
    },
    {
        &html.Node{
            Type: html.ElementNode,
            Data: "html",
        },
    },
    {
        &html.Node{
            Type: html.ElementNode,
            Data: "head",
        },
    },
}

func TestGetParent(t *testing.T) {
    r := strings.NewReader(`
    <html>
        <head>
            <title></title>
        </head>
        <body>
            <div class="red-parent">
                <div class="red-child"></div>
            </div>
            <span class="green-parent">
                <span></span>
            </span>
        </body>
    </html>
    `)
    root, _ := GetRootElement(r)

    cases := []struct{
        got Element
        element Element
    }{
        { 
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByClass("red-parent", 1)[0].
                    FindChildrenByClass("red-child", 1)[0].
                        Parent(),
            expectedOutputE[0],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByClass("green-parent", 1)[0].
                    FindAllChildren(-1)[0].
                        Parent(),
            expectedOutputE[1],
        },
        {
            root.FindChildrenByElement("head", 1)[0].
                Parent(),
            expectedOutputE[2],
        },
        {
            root.FindChildrenByElement("head", 1)[0].
                FindChildrenByElement("title", 1)[0].
                    Parent(),
            expectedOutputE[3],
        },
    }

    for i, test := range cases {
        t.Run(fmt.Sprintf("case #%d", i), func(t *testing.T) {
            if test.got == (Element{}) && test.element != (Element{}) {
                t.Fatalf("got=nil, expectedOut != nil\n")
            }

            equal := compareTypeAndData(test.got, test.element); 
            if equal == false {
                t.Errorf("Either Type or Data of two elements are not equal\n")
            }
        }) 
    }
}
