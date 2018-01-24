package main

import (
	"fmt"
	"testing"

	grokcgo "github.com/blakesmith/go-grok/grok"
	grokkky "github.com/logrusorgru/grokky"
	grokgo "gitlab.10101111.com/oped/uparse.git/lib/grok"
)

func Benchmark_ky(b *testing.B) {

	g := grokkky.NewBase()
	p, err := g.Compile("%{COMMONAPACHELOG}")
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	for i := 0; i < b.N; i++ {
		values := p.Parse(`127.0.0.1 - - [23/Apr/2014:22:58:32 +0200] "GET /index.php HTTP/1.1" 404 207`)
		if values == nil {
			fmt.Printf("err:%v", err)
		} else {
			//for k, v := range values {
			//fmt.Printf("%+15s: %s\n", k, v)
			//}
		}
	}
	b.ReportAllocs()

}
func Benchmark_Gogo(b *testing.B) {

	g, _ := grokgo.New()
	for i := 0; i < b.N; i++ {
		_, err := g.Parse("%{COMMONAPACHELOG}", `127.0.0.1 - - [23/Apr/2014:22:58:32 +0200] "GET /index.php HTTP/1.1" 404 207`)
		if err != nil {
			fmt.Printf("err:%v", err)
		} else {
			//for k, v := range values {
			//fmt.Printf("%+15s: %s\n", k, v)
			//}
		}
	}
	b.ReportAllocs()

}

func Benchmark_Cgo(b *testing.B) {

	g := grokcgo.New()
	err := g.Compile("%{COMMONAPACHELOG}")
	if err != nil {
		return
	}
	for i := 0; i < b.N; i++ {
		values := g.Match(`127.0.0.1 - - [23/Apr/2014:22:58:32 +0200] "GET /index.php HTTP/1.1" 404 207`)
		if values == nil {
			fmt.Printf("err:%v", err)
		} else {

			_ = values.Captures()
			//for k, v := range values.Captures() {
			//fmt.Printf("%+15s: %s\n", k, v[0])
			//}
		}
	}
	b.ReportAllocs()

}
