# assetfile
AssetFile for Template System

Code Cover %100 https://gocover.io/github.com/akmyazilim/assetfile

Doc https://godoc.org/github.com/akmyazilim/assetfile

# install
```
go get -u github.com/akmyazilim/assetfile
```

# example
```
package main

import (
	"fmt"

	"github.com/akmyazilim/assetfile"
)

func main() {
	css := assetfile.New()
	css.Add("a.css")
	css.Add("b.css")
	css.Add("a.css")
	css.Add("c.css")
	css.Add("b.css")
	css.Add("foo_%s.css", "module")
	css.Add("foo_%s_%d.css", "module", 1)

	fmt.Println(css.List()) // [a.css b.css c.css foo_module.css foo_module_1.css]
}

```