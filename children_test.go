package main

import (
    "testing"
    "golang.org/x/net/html"
    //"io"
    "strings"
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
    }

    for i, test := range cases {
        if len(test.got) != len(test.elements) {
            t.Errorf("case#%d, len(got)=%d != len(expectedOut)=%d\n", i+1,
                len(test.got), len(test.elements))
        }

        for j, element := range test.elements {
            if equal := compareTypeAndData(element, test.got[j]); equal == false {
                t.Errorf("case#%d, %d) either html.Node.Type or html.Node.Data of two elements are not equal\n", i+1, j)
            }

            if contains := containsClass(test.got[j].node.Attr, element.node.Attr[0]);
            contains == false {
                t.Errorf("case#%d, %d) got doesn't contain necessary class names\n",
                i+1, j) 
            }
        }
    
    }

}
