package main

import (
	"fmt"
)

type finder interface {
	find(s string) bool
}

type maker struct {
	brand        string
	manufacturer string
}

func (m *maker) find(s string) bool {

	return m.brand == s
}

type car struct {
	maker
	model string
	year  string
}

func (c *car) find(s string) bool {

	if c.year != s {
		return c.maker.find(s)
	}

	return true
}

func search(f []finder, k string) {
	fmt.Printf("Searching For: keyword %s\n", k)

	for i := range f {

		if f[i].find(k) {
			fmt.Printf("Found %+v\n", f[i])
			return
		}

	}

	fmt.Println("Not Found " + k)

}

func main() {

	items := []finder{
		&car{
			maker: maker{
				brand:        "Hyundai",
				manufacturer: "Hyundai Motors",
			},
			model: "sonata",
			year:  "2015",
		},
		&car{
			maker: maker{
				brand:        "Genesis",
				manufacturer: "Hyundai Motors",
			},
			model: "G80",
			year:  "2018",
		},
		&car{
			maker: maker{
				brand:        "Chevrolet",
				manufacturer: "GM",
			},
			model: "VoIt",
			year:  "2017",
		},
		&car{
			maker: maker{
				brand:        "Tesla",
				manufacturer: "Tesla Motors",
			},
			model: "Tesla 3",
			year:  "2018",
		},
	}

	// for i := range items {
	// 	fmt.Println(items[i])
	// }

	search(items, "2018")

}
