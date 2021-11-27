package main

import (
	link "04htmlparser"
	"fmt"
	"strings"
)

// var example1 = `
// <html>
// <body>
//   <h1>Hello!</h1>
//   <a href="/other-page">A link to another page
// 	<span> some span</span>
// 	</a>
// 	<a href="/hello-there">another link</a>
// 	</body>
// </html>
// `
var ex4 = `
<html>
<body>
  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
</body>
</html>
`

func main() {
	r := strings.NewReader(ex4)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
