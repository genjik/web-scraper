package main

import (
    "testing"
    "golang.org/x/net/html"
    "strings"
    "fmt"
)

// FindAllChildren
var expectedOutputA []Element = []Element{
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
    {
        &html.Node{
            Type: html.ElementNode,
            Data: "div",
            Attr: []html.Attribute{
                {Namespace: "", Key: "class", Val: "red-child"},
            },
        },
    },
}

func TestFindAllChildren(t *testing.T) {
    r := strings.NewReader(`
    <html>
        <head></head>
        <body>
            <div class="red">
                <div class="red-child"></div>
            </div>
            <div class="green">
                <div class="green-child"></div>
            </div>
            <div class="blue">
                <div class="blue-child"></div>
            </div>
        </body>
    </html>
    `)
    root, _ := GetRootElement(r)

    cases := []struct{
        got []Element
        elements []Element
    }{
        {
            // gets nothing 
            Element{root.node.LastChild}.FindAllChildren(0),
            expectedOutputA[:0],
        },
        {
            // gets 1 child of body
            Element{root.node.LastChild}.FindAllChildren(1),
            expectedOutputA[:1],
        },
        {
            // gets 2 children of body
            Element{root.node.LastChild}.FindAllChildren(2),
            expectedOutputA[:2],
        },
        {
            // gets 5 children of body, should return 3
            Element{root.node.LastChild}.FindAllChildren(5),
            expectedOutputA[:3],
        },
        {
            // gets all children of body, should return all
            Element{root.node.LastChild}.FindAllChildren(-1),
            expectedOutputA[:3],
        },
        {
            // gets nothing
            Element{root.node.FirstChild}.FindAllChildren(-1),
            expectedOutputA[:0],
        },
        {
            // gets all children of "red" div class
            // CHANGE IT LATER
            Element{root.node.LastChild.FirstChild.NextSibling}.FindAllChildren(-1),
            expectedOutputA[3:4],
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

                contains := containsSel(test.got[j].node.Attr, element.node.Attr[0], element.node.Attr[0].Key) 
                if contains == false {
                    t.Errorf("%d) got doesn't contain necessary class names\n", j) 
                }
            }
        }) 
    }

}

// FindChildrenByClass
var expectedOutputB []Element = []Element{
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
                {Namespace: "", Key: "class", Val: "red"},
            },
        },
    },
    {
        &html.Node{
            Type: html.ElementNode,
            Data: "div",
            Attr: []html.Attribute{
                {Namespace: "", Key: "class", Val: "red"},
            },
        },
    },
}

func TestFindChildrenByClass(t *testing.T) {
    r := strings.NewReader(`
    <html>
        <head></head>
        <body>
            <div class="red"></div>
            <div class="green"></div>
            <div class="red"></div>
            <div class="blue"></div>
            <div class="red"></div>
        </body>
    </html>
    `)
    root, _ := GetRootElement(r)

    cases := []struct{
        got []Element
        elements []Element
    }{
        {
            // gets nothing 
            Element{root.node.LastChild}.FindChildrenByClass("red", 0),
            expectedOutputB[:0],
        },
        {
            // gets children with red class, no limit 
            Element{root.node.LastChild}.FindChildrenByClass("red", -1),
            expectedOutputB[0:],
        },
        {
            // gets children with red class, limit 1 
            Element{root.node.LastChild}.FindChildrenByClass("red", 1),
            expectedOutputB[0:1],
        },
        {
            // gets children with red class, limit 2 
            Element{root.node.LastChild}.FindChildrenByClass("red", 2),
            expectedOutputB[0:2],
        },
        {
            // gets children with red class, limit 3 
            Element{root.node.LastChild}.FindChildrenByClass("red", 3),
            expectedOutputB[0:3],
        },
        {
            // gets nothing 
            Element{root.node.LastChild}.FindChildrenByClass("yellow", -1),
            expectedOutputB[:0],
        },
        {
            // gets nothing 
            Element{root.node.FirstChild}.FindChildrenByClass("yellow", -1),
            expectedOutputB[:0],
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

                contains := containsSel(test.got[j].node.Attr, element.node.Attr[0], "class") 
                if contains == false {
                    t.Errorf("%d) got doesn't contain necessary class names\n", j) 
                }
            }
        }) 
    }
}

// FindChildrenByElement
var expectedOutputC []Element = []Element{
    {
        &html.Node{
            Type: html.ElementNode,
            Data: "div",
            Attr: []html.Attribute{
                {Namespace: "", Key: "", Val: ""},
            },
        },
    },
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
            Data: "span",
            Attr: []html.Attribute{
                {Namespace: "", Key: "", Val: ""},
            },
        },
    },
    {
        &html.Node{
            Type: html.ElementNode,
            Data: "h1",
            Attr: []html.Attribute{
                {Namespace: "", Key: "", Val: ""},
            },
        },
    },
}

func TestFindChildrenByElement(t *testing.T) {
    r := strings.NewReader(`
    <html>
        <head></head>
        <body>
            <div>123</div>
            <div class="red"></div>
            <span></span>
            <h1></h1>
        </body>
    </html>
    `)
    root, _ := GetRootElement(r)

    cases := []struct{
        got []Element
        elements []Element
    }{
        {
            // gets nothing 
            Element{root.node.LastChild}.FindChildrenByElement("div", 0),
            expectedOutputC[:0],
        },
        {
            // gets div children, no limit 
            Element{root.node.LastChild}.FindChildrenByElement("div", -1),
            expectedOutputC[0:2],
        },
        {
            // gets div children, limit 1 
            Element{root.node.LastChild}.FindChildrenByElement("div", 1),
            expectedOutputC[0:1],
        },
        {
            // gets div children, limit 2 
            Element{root.node.LastChild}.FindChildrenByElement("div", 2),
            expectedOutputC[0:2],
        },
        {
            // gets span children, no limit 
            Element{root.node.LastChild}.FindChildrenByElement("span", -1),
            expectedOutputC[2:3],
        },
        {
            // gets h1 children, no limit 
            Element{root.node.LastChild}.FindChildrenByElement("h1", -1),
            expectedOutputC[3:4],
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
            }
        }) 
    }
}


// FindChildById
var expectedOutputD []Element = []Element{
    {
        &html.Node{
            Type: html.ElementNode,
            Data: "div",
            Attr: []html.Attribute{
                {Namespace: "", Key: "id", Val: "special"},
            },
        },
    },
    {
        &html.Node{
            Type: html.ElementNode,
            Data: "span",
            Attr: []html.Attribute{
                {Namespace: "", Key: "id", Val: "also-special"},
            },
        },
    },
}

func TestFindChildById(t *testing.T) {
    r := strings.NewReader(`
    <html>
        <head></head>
        <body>
            <div id="special" class="red"></div>
            <div class="green"></div>
            <div class="red"></div>
            <span id="also-special"></span>
            <div class="red"></div>
        </body>
    </html>
    `)
    root, _ := GetRootElement(r)

    cases := []struct{
        got Element
        elements []Element
    }{
        { 
            Element{root.node.LastChild}.FindChildById("nothing"),
            expectedOutputD[:0],
        },
        {
            Element{root.node.LastChild}.FindChildById("green"),
            expectedOutputD[:0],
        },
        {
            Element{root.node.LastChild}.FindChildById("special"),
            expectedOutputD[:1],
        },
        {
            Element{root.node.LastChild}.FindChildById("also-special"),
            expectedOutputD[1:2],
        },
        {
            Element{root.node.LastChild}.FindChildById("special also-special"),
            expectedOutputD[:0],
        },
    }

    for i, test := range cases {
        t.Run(fmt.Sprintf("case #%d", i), func(t *testing.T) {
            if test.got == (Element{}) && len(test.elements) > 0 {
                t.Fatalf("got=nil, expectedOut != nil\n")
            }

            for _, element := range test.elements {
                equal := compareTypeAndData(element, test.got); 
                if equal == false {
                    t.Errorf("Either html.Node.Type or html.Node.Data of two elements are not equal\n")
                }

                contains := containsSel(test.got.node.Attr, element.node.Attr[0], "id") 
                if contains == false {
                    t.Errorf("element doesn't contain id\n") 
                }
            }
        }) 
    }
}
