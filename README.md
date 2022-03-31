# OrderedForm
Go package to create orderd query strings

Usage 
```
package main

import "github.com/66niko99/OrderedValues"

func main() {
  values := OrderdValues.Values{}
  values.Add("abc", "123")
  values.Add("def", "456")
  values.Add("encoded", "x x x")

  fmt.println(values.String) // abc=123&def=456&encoded=x%20x%20x
}
```
