# assetfilepongo
AssetFile for pongo2 template system

Code Cover %100 https://gocover.io/github.com/s4l1h/assetfile/assetfilepongo

Doc https://godoc.org/github.com/s4l1h/assetfile/assetfilepongo

# install
```
go get -u github.com/s4l1h/assetfile/assetfilepongo
```

# example
```
package main

import (
	"fmt"

	"github.com/s4l1h/assetfile/assetfilepongo"
	"github.com/flosch/pongo2"
)

func main() {

	css := assetfilepongo.New()

	// Compile the template first (i. e. creating the AST)
	tpl, err := pongo2.FromString(`
		{{CSS.Add("a.css")}}
		{{CSS.Add("b.css")}}
		{{CSS.Add("foo.css?v=%d",10)}}
		
		{{ hello }}
		{% for cssURL in CSS.List()%}
			<link href="{{ cssURL }}" rel="stylesheet">
		{% endfor %}
		`)
	if err != nil {
		panic(err)
	}

	out, err := tpl.Execute(pongo2.Context{"CSS": css, "hello": "Hello Man"})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output:
	/*

	   Hello Man
	           <link href="a.css" rel="stylesheet">

	           <link href="b.css" rel="stylesheet">

	           <link href="foo.css?v=10" rel="stylesheet">
	*/

	//fmt.Println(css.List()) // [a.css b.css foo.css?v=10]
}


```