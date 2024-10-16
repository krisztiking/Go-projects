package main

import (
	"fmt"
	"strings"
)

var exampleHtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
</body>
</html>`

// this file you can only call with parse.go
func forparse() {
	r := strings.NewReader(exampleHtml)
	links, err := Parse(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", links)
}
