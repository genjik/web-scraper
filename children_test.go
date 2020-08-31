package webscraper

import (
    "golang.org/x/net/html"
    "testing"
    "strings"
    "fmt"
)

//FindAllChildren
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
    {
        &html.Node{
            Type: html.ElementNode,
            Data: "div",
            Attr: []html.Attribute{
                {Namespace: "", Key: "class", Val: "green-child1"},
            },
        },
    },
    {
        &html.Node{
            Type: html.ElementNode,
            Data: "div",
            Attr: []html.Attribute{
                {Namespace: "", Key: "class", Val: "green-child2"},
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
                <div class="green-child1"></div>
                <div class="green-child2"></div>
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
            root.FindChildrenByElement("body", 1)[0].
                FindAllChildren(0),
            expectedOutputA[:0],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindAllChildren(1),
            expectedOutputA[:1],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindAllChildren(2),
            expectedOutputA[:2],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindAllChildren(5),
            expectedOutputA[:3],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindAllChildren(-1),
            expectedOutputA[:3],
        },
        {
            root.FindChildrenByElement("head", 1)[0].
                FindAllChildren(-1),
            expectedOutputA[:0],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByClass("red", 1)[0].
                    FindAllChildren(-1),
            expectedOutputA[3:4],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByClass("green", 1)[0].
                    FindAllChildren(-1),
            expectedOutputA[4:6],
        },
    }

    for i, test := range cases {
        t.Run(fmt.Sprintf("case #%d", i), func(t *testing.T) {
            if len(test.got) != len(test.elements) {
                t.Fatalf("len(got)=%d, len(expectedOut)=%d\n", len(test.got), len(test.elements))
            }

            for j, element := range test.elements {
                equal := compareTypeAndData(element, test.got[j]); 
                if equal == false {
                    t.Errorf("%d) Type or Data of two elements are not equal\n", j)
                }

                contains := containsSel(test.got[j].node.Attr, element.node.Attr[0], element.node.Attr[0].Key) 
                if contains == false {
                    t.Errorf("%d) Element doesn't contain classes \n", j) 
                }
            }
        }) 
    }

}

//FindChildrenByClass
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
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByClass("red", 0),
            expectedOutputB[:0],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByClass("red", -1),
            expectedOutputB[0:],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByClass("red", 1),
            expectedOutputB[0:1],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByClass("red", 2),
            expectedOutputB[0:2],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByClass("red", 3),
            expectedOutputB[0:3],
        },
        {
            root.FindChildrenByElement("head", 1)[0].
                FindChildrenByClass("yellow", -1),
            expectedOutputB[:0],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByClass("yellow", -1),
            expectedOutputB[:0],
        },
    }

    for i, test := range cases {
        t.Run(fmt.Sprintf("case #%d", i), func(t *testing.T) {
            if len(test.got) != len(test.elements) {
                t.Fatalf("len(got)=%d, len(expectedOut)=%d\n", len(test.got), len(test.elements))
            }

            for j, element := range test.elements {
                equal := compareTypeAndData(element, test.got[j]); 
                if equal == false {
                    t.Errorf("%d) Type or Data of two elements are not equal\n", j)
                }

                contains := containsSel(test.got[j].node.Attr, element.node.Attr[0], "class") 
                if contains == false {
                    t.Errorf("%d) Element doesn't contain classes\n", j) 
                }
            }
        }) 
    }
}

//FindChildrenByElement
var expectedOutputC []Element = []Element{
    {
        &html.Node{
            Type: html.ElementNode,
            Data: "div",
        },
    },
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
            Data: "h1",
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
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByElement("div", 0),
            expectedOutputC[:0],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByElement("div", -1),
            expectedOutputC[0:2],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByElement("div", 1),
            expectedOutputC[0:1],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByElement("div", 2),
            expectedOutputC[0:2],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByElement("span", -1),
            expectedOutputC[2:3],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByElement("h1", -1),
            expectedOutputC[3:4],
        },
    }

    for i, test := range cases {
        t.Run(fmt.Sprintf("case #%d", i), func(t *testing.T) {
            if len(test.got) != len(test.elements) {
                t.Fatalf("len(got)=%d, len(expectedOut)=%d\n", len(test.got), len(test.elements))
            }

            for j, element := range test.elements {
                equal := compareTypeAndData(element, test.got[j]); 
                if equal == false {
                    t.Errorf("%d) Type or Data of two elements are not equal\n", j)
                }
            }
        }) 
    }
}


//FindChildById
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
            root.FindChildrenByElement("body", 1)[0].
                FindChildById("nothing"),
            expectedOutputD[:0],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildById("green"),
            expectedOutputD[:0],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildById("special"),
            expectedOutputD[:1],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildById("also-special"),
            expectedOutputD[1:2],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildById("special also-special"),
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
                    t.Errorf("Type or Data of two elements are not equal\n")
                }

                contains := containsSel(test.got.node.Attr, element.node.Attr[0], "id") 
                if contains == false {
                    t.Errorf("Element doesn't contain id\n") 
                }
            }
        }) 
    }
}
