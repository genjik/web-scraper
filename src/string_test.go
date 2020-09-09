package webscraper

import (
    "testing"
    "fmt"
)

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
            got := compareStr(test.a, test.b)

            if got != test.out {
                t.Errorf("output=%t, expected=%t\n", got, test.out)
            }
        })
    }
}

func TestHasRepetition(t *testing.T) {
    cases := []struct {
        val []string
        out int
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
            if got := hasRepetition(test.val); got != test.out {
                t.Errorf("got=%d, expected=%d\n", got, test.out)
            }
        })
    }
}
