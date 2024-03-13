### Map

- Maps are like dictionary in Python
- All keys of same type and values also of same type, key and values don't have to be of the same type
- https://app.diagrams.net/#Hnuthanc%2Fgolang_project%2Fmain%2Fdiagrams%2F05%2Fdiagrams.xml#%7B%22pageId%22%3A%2290dd907b-916a-6cda-bd28-38c707c73094%22%7D
- Check `main.go` for Examples
```go
package main

import "fmt"

func main() {
	// 1st way
	// colors := map[string]string{
	// 	"red": "#ff0000",
	// 	"blue": "#f2312",
	// }

	// 2nd way
	// var colors map[string]string
	// colors = make(map[string]string)
	// colors["red"] = "#ff0000"
	// colors["blue"] = "#ff0120"

	// 3rd way
	colors := make(map[int]string)
	colors[1] = "#ff0000"
	colors[2] = "#ff0120"
	colors[3] = "#ff0120"

	delete(colors, 1)

	fmt.Println(colors)
}
```

### Structs vs Maps

* https://app.diagrams.net/#Hnuthanc%2Fgolang_project%2Fmain%2Fdiagrams%2F05%2Fdiagrams.xml#%7B%22pageId%22%3A%223f52bd12-7fc2-9203-a9ff-7c15ea555526%22%7D