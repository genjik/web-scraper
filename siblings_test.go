package webscraper

import (
    "golang.org/x/net/html"
    "testing"
    "strings"
    "fmt"
)

//FindAllSiblings
var expectedOutputF []Element = []Element{
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
            Data: "head",
        },
    },
    {
        &html.Node{
            Type: html.ElementNode,
            Data: "body",
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

func TestFindAllSiblings(t *testing.T) {
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
                FindChildrenByClass("red", 1)[0].
                    FindAllSiblings(0),
            expectedOutputF[:0],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByClass("red", 1)[0].
                    FindAllSiblings(1),
            expectedOutputF[1:2],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByClass("red", 1)[0].
                    FindAllSiblings(2),
            expectedOutputF[1:3],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByClass("red", 1)[0].
                    FindAllSiblings(5),
            expectedOutputF[1:3],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByClass("red", 1)[0].
                    FindAllSiblings(-1),
            expectedOutputF[1:3],
        },
        {
            root.FindChildrenByElement("head", 1)[0].
                FindAllSiblings(-1),
            expectedOutputF[4:5],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByClass("green", 1)[0].
                    FindChildrenByClass("green-child1", 1)[0].
                        FindAllSiblings(-1),
            expectedOutputF[6:7],
        },
        {
            root.FindAllSiblings(-1),
            expectedOutputF[:0],
        },
        //This one produces huge bug, fix it later
        //{
        //    root.Parent().
        //        FindAllSiblings(-1),
        //    expectedOutputF[:0],
        //},
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

                if element.node.Attr != nil {
                    contains := containsSel(test.got[j].node.Attr, element.node.Attr[0], element.node.Attr[0].Key) 
                    if contains == false {
                        t.Errorf("%d) Element doesn't contain classes \n", j) 
                    }
                }
            }
        }) 
    }
    
}

//FindSiblingsByClass
var expectedOutputG []Element = []Element{
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

func TestFindSiblingsByClass(t *testing.T) {
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
                FindChildrenByClass("red", 1)[0].
                    FindSiblingsByClass("green", 0),
            expectedOutputG[:0],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByClass("red", 1)[0].
                    FindSiblingsByClass("green", -1),
            expectedOutputG[1:2],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByClass("red", 1)[0].
                    FindSiblingsByClass("green", 1),
            expectedOutputG[1:2],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByClass("red", 1)[0].
                    FindSiblingsByClass("green", 2),
            expectedOutputG[1:2],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByClass("red", 1)[0].
                    FindSiblingsByClass("red", -1),
            expectedOutputG[3:5],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByClass("red", 1)[0].
                    FindSiblingsByClass("yellow", -1),
            expectedOutputG[:0],
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


//FindSiblingsByElement
var expectedOutputH []Element = []Element{
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
            Data: "span",
            Attr: []html.Attribute{
                {Namespace: "", Key: "class", Val: "blue"},
            },
        },
    },
    {
        &html.Node{
            Type: html.ElementNode,
            Data: "h1",
            Attr: []html.Attribute{
                {Namespace: "", Key: "class", Val: "yellow"},
            },
        },
    },
    {
        &html.Node{
            Type: html.ElementNode,
            Data: "h2",
            Attr: []html.Attribute{
                {Namespace: "", Key: "class", Val: "pink"},
            },
        },
    },
    {
        &html.Node{
            Type: html.ElementNode,
            Data: "h2",
            Attr: []html.Attribute{
                {Namespace: "", Key: "class", Val: "pink"},
            },
        },
    },
}

func TestFindSiblingsByElement(t *testing.T) {
    r := strings.NewReader(`
    <html>
        <head></head>
        <body>
            <div class="red">123</div>
            <div class="green"></div>
            <span class="blue"></span>
            <h1 class="yellow"></h1>
            <h2 class="pink"></h2>
            <h2 class="pink"></h2>
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
                FindChildrenByClass("red", 1)[0].
                    FindSiblingsByElement("div", -1),
            expectedOutputH[1:2],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByClass("red", 1)[0].
                    FindSiblingsByElement("span", -1),
            expectedOutputH[2:3],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByClass("red", 1)[0].
                    FindSiblingsByElement("h1", -1),
            expectedOutputH[3:4],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByClass("red", 1)[0].
                    FindSiblingsByElement("h2", -1),
            expectedOutputH[4:6],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByClass("red", 1)[0].
                    FindSiblingsByElement("div", 0),
            expectedOutputH[:0],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByClass("red", 1)[0].
                    FindSiblingsByElement("h1", 0),
            expectedOutputH[:0],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByClass("red", 1)[0].
                    FindSiblingsByElement("div", 1),
            expectedOutputH[1:2],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByClass("green", 1)[0].
                    FindSiblingsByElement("div", 1),
            expectedOutputH[0:1],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildrenByClass("red", 1)[0].
                    FindSiblingsByElement("div", 2),
            expectedOutputH[1:2],
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


//FindSiblingById
var expectedOutputI []Element = []Element{
    {
        &html.Node{
            Type: html.ElementNode,
            Data: "div",
            Attr: []html.Attribute{
                {Namespace: "", Key: "id", Val: "red"},
            },
        },
    },
    {
        &html.Node{
            Type: html.ElementNode,
            Data: "div",
            Attr: []html.Attribute{
                {Namespace: "", Key: "id", Val: "green"},
            },
        },
    },
    {
        &html.Node{
            Type: html.ElementNode,
            Data: "div",
            Attr: []html.Attribute{
                {Namespace: "", Key: "id", Val: "red"},
            },
        },
    },
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

func TestFindSiblingById(t *testing.T) {
    r := strings.NewReader(`
    <html>
        <head></head>
        <body>
            <div id="red"></div>
            <div id="green"></div>
            <div id="red"></div>
            <div id="special"></div>
            <span id="also-special"></span>
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
                FindChildById("red").
                    FindSiblingById("nothing"),
            expectedOutputI[:0],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildById("red").
                    FindSiblingById("green"),
            expectedOutputI[1:2],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildById("red").
                    FindSiblingById("red"),
            expectedOutputI[2:3],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildById("red").
                    FindSiblingById("special"),
            expectedOutputI[3:4],
        },
        {
            root.FindChildrenByElement("body", 1)[0].
                FindChildById("red").
                    FindSiblingById("also-special"),
            expectedOutputI[4:5],
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
