package main

import (
	"fmt"

	"github.com/amirhosein-dev/go-dsa/pkg/trees" // مسیر خودت در go.mod
)

func main() {
	var t trees.Tree[int]
	t.Insert(60)
	t.Insert(78)
	t.Insert(23)
	t.Insert(47)
	t.Insert(24)
	t.Insert(2)
	t.Insert(90)

	for e := range t.PreOrder() {
		fmt.Println("Visiting:", e)
		if e == 47 {
			break
		}
	}
}
