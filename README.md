## Google Translate API in Golang

 - **free**
 - **thread-safe**
 - **three endpoints**
 - **resilient to socket buffer errors**

Install:
```
go get github.com/spywiree/googletranslate
```

Example usage:
```go
package main

import (
    "fmt"
    gt "github.com/spywiree/googletranslate"
)

func main(){
    const text string = `Hello, World!`
    // you can use "auto" for source language
    // so, translator will detect language
    result, _ := gt.Translate(text, "en", "es")
    fmt.Println(result)
    // Output: "Hola, Mundo!"
}
```

Would you like to perform photo-to-photo translations?\
Take a look at my other package: [translateimage](https://github.com/spywiree/translateimage).