package main

import (
	"encoding/csv"
	"fmt"
	"github.com/karantin2020/csvparse"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
	"strings"
)

// Foo struct
type Foo struct {
	a string
	b float64
}

// NewFoo implemenets csvparse.Newer
func NewFoo() csvparse.Unmarshaller {
	return &Foo{}
}

var enc, _ = charset.Lookup(encf)
var rl = []csvparse.Unmarshaller{
	&Foo{"привет", 9.96370968},
	&Foo{"fdas", 9.38186666},
	&Foo{"fds", 1.713333},
	&Foo{"hgd", 2.6974},
}

func main() {
	re := strings.NewReader(data)
	r := transform.NewReader(re, enc.NewDecoder())
	rc := csv.NewReader(r)
	if result, err := csvparse.ReadAll(rc, NewFoo); err != nil {
		fmt.Printf("%#v\n", err)
	} else {
		for _, i := range result {
			fmt.Printf("%+v\n", *i.(*Foo))
		}
		fmt.Printf("fooEqual: %#v\n", fooEqual(result, rl))
	}
}

const encf = "utf8"
const data = `"привет",9.96370968
"fdas",9.38186666
"fds",1.713333
"hgd",2.6974
`

func fooEqual(one []csvparse.Unmarshaller, two []csvparse.Unmarshaller) bool {
	for i, v := range one {
		if *v.(*Foo) != *two[i].(*Foo) {
			return false
		}
	}
	return true
}
