package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	Result "github.com/go-composites/result/src"
)

func main() {
	r := Result.New()
	spew.Dump(r)
	fmt.Println(r.HasError())
}
