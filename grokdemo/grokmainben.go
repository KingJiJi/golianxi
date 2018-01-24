package main

import (
	"fmt"

	"gitlab.10101111.com/oped/uparse.git/lib/grok"
)

func main() {
	fmt.Println("# Default Capture :")
	g, _ := grok.New()
	for i := 0; i < 10000; i++ {
		values, err := g.Parse("%{COMMONAPACHELOG}", `127.0.0.1 - - [23/Apr/2014:22:58:32 +0200] "GET /index.php HTTP/1.1" 404 207`)
		if err != nil {
			fmt.Printf("err:%v", err)
		} else {
			for k, v := range values {
				fmt.Printf("%+15s: %s\n", k, v)
			}
		}
	}

}
