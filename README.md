# pretty-go-bytes

Golang port of the pretty-bytes node.js module

Original node.js module can be found here: https://github.com/sindresorhus/pretty-bytes

## Installation

```bash
go get github.com/jldeep/pretty-go-bytes
```


## Usage

```js
package main
import "github.com/jldeep/pretty-go-bytes"

func main() {
  pretty, error := prettybytes.Pretty(1337)
  
  if error != nil {
    panic(error)
  }
  
	fmt.Printf(pretty)
  // => '1.34 kB'
}
```


## License

Copyright (c) 2017 José Luis Zenteno De León

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
