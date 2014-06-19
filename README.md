HttpError.go
=====

HttpError.go is simple http error generator

##Example
```
package main

import (
	"github.com/nathanfaucett/httperror"
	"fmt"
)

func main() {
	err1 := httperror.New(500, "Some Message")
	fmt.Println(err1.Error())
	
	err2 := httperror.NewCode(500) //message equals status code text from net/http
	fmt.Println(err2.Error())
	
	err3 := httperror.NewMessage("Some Message") //code is 500
	fmt.Println(err3.Error())
}
```