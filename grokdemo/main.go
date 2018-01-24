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
	values := g.Match(`127.0.0.1 - - [23/Apr/2014:22:58:32 +0200] "GET /index.php HTTP/1.1" 404 207`)
	if values == nil {
		fmt.Printf("err:%v", err)
	} else {
		for k, v := range values.Captures() {
			fmt.Printf("%+15s: %s\n", k, v[0])
		}
	}

	fmt.Println("\n# Named Capture :")
	values = g.Match(`127.0.0.1 - - [23/Apr/2014:22:58:32 +0200] "GET /index.php HTTP/1.1" 404 207`)
	for k, v := range values.Captures() {
		fmt.Printf("%+15s: %s\n", k, v)
	}

	fmt.Println("\n# Add custom patterns :")
	// We add 3 patterns to our Grok instance, to structure an IRC message
	g = grok.New()
	g.AddPattern("IRCUSER", `\A@(\w+)`)
	g.AddPattern("IRCBODY", `.*`)
	g.AddPattern("IRCMSG", `%{IRCUSER:user} .* : %{IRCBODY:message}`)
	err = g.Compile("%{IRCMSG}")
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	values = g.Match(`@vjeantet said : Hello !`)
	for k, v := range values.Captures() {
		fmt.Printf("%+15s: %s\n", k, v)
	}
}
