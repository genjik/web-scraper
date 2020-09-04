package main

import (
    "testing"
    "fmt"
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
            output := validateAttrs(test.a)

            if compareStr(test.out, output) == false {
                t.Errorf("expected=%+v, got=%+v\n", test.out, output) 
            }
        })
    }
}

func TestCompareStr(t *testing.T) {
    cases := []struct{
        a []string
        b []string
        out bool
    }{
        {
            []string{"class", "red", "id", "red"},
            []string{"class", "red", "id", "red"},
            true,
        },
        {
            []string{"class", "red", "id", "red"},
            []string{"blue", "red", "id", "red"},
            false,
        },
        {
            []string{"red", "id", "red"},
            []string{"blue", "red", "id", "red"},
            false,
        },
    }

    for i, test := range cases {
        t.Run(fmt.Sprintf("Case #%d\n", i), func(t *testing.T) {
            output := compareStr(test.a, test.b)

            if output != test.out {
                t.Errorf("output=%t, expected=%t\n", output, test.out)
            }
        })
    }
}
