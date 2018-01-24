package main

import (
	"fmt"

	grok "github.com/blakesmith/go-grok/grok"
)

func main() {
	fmt.Println("# Default Capture :")
	g := grok.New()
	err := g.Compile("%{COMMONAPACHELOG}")
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	for i := 0; i < 10000; i++ {
		values := g.Match(`127.0.0.1 - - [23/Apr/2014:22:58:32 +0200] "GET /index.php HTTP/1.1" 404 207`)
		if values == nil {
			fmt.Printf("err:%v", err)
		} else {
			for k, v := range values.Captures() {
				fmt.Printf("%+15s: %s\n", k, v[0])
			}
		}
	}

}
