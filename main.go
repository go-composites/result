package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	Result "github.com/golang-cop/result/src"
)

func main() {
	r := Result.New()
	spew.Dump(r)
	fmt.Println(r.HasError())
}
