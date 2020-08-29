package main

import (
    "testing"
    "golang.org/x/net/html"
    //"io"
    "strings"
    "fmt"
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
                {Namespace: "", Key: "class", Val: "blue"},
            },
        },
    },
}

func TestFindAllChildren(t *testing.T) {
    r := strings.NewReader(`
    <html>
        <head></head>
        <body>
            <div class="red"></div>
            <div class="green"></div>
            <div class="blue"></div>
        </body>
    </html>
    `)
    root, _ := GetRootElement(r)

    cases := []struct{
        got []Element
        elements []Element
    }{
        {
            // gets all children of body
            Element{root.node.LastChild}.FindAllChildren(false, 0),
            expectedOutput[:3],
        },
        {
            // gets 2 children of body
            Element{root.node.LastChild}.FindAllChildren(false, 2),
            expectedOutput[:2],
        },
        {
            // gets 5 children of body, should return 3
            Element{root.node.LastChild}.FindAllChildren(false, 5),
            expectedOutput[:3],
        },
        {
            // gets 1 child of body
            Element{root.node.LastChild}.FindAllChildren(false, 1),
            expectedOutput[:1],
        },
        {
            // gets all children of body, should return all
            Element{root.node.LastChild}.FindAllChildren(false, -1),
            expectedOutput[:3],
        },
    }

    for i, test := range cases {
        t.Run(fmt.Sprintf("case #%d", i), func(t *testing.T) {
            if len(test.got) != len(test.elements) {
                t.Fatalf("len(got)=%d != len(expectedOut)=%d\n", len(test.got), len(test.elements))
            }

            for j, element := range test.elements {
                equal := compareTypeAndData(element, test.got[j]); 
                if equal == false {
                    t.Errorf("%d) either html.Node.Type or html.Node.Data of two elements are not equal\n", j)
                }

                contains := containsClass(test.got[j].node.Attr, element.node.Attr[0])
                if contains == false {
                    t.Errorf("%d) got doesn't contain necessary class names\n", j) 
                }
            }
        }) 
    }

}
