package main

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/moovweb/rubex"
	"github.com/tsingakbar/cre2"
)

func Benchmark_regex(b *testing.B) {

	g := regexp.MustCompile("h1ealth(C|c)heck")
	var aa, bb int
	for i := 0; i < b.N; i++ {
		ismatch := g.MatchString(`healthcheck"`)
		if ismatch {
			aa++
		} else {
			bb++
		}
	}
	b.ReportAllocs()
	fmt.Printf("\n\ntrue:%d", aa)
	fmt.Printf("false:%d\n\n", bb)

}

func Benchmark_rubes(b *testing.B) {
	var aa, bb int
	g := rubex.MustCompile("health(C|c)heck")
	for i := 0; i < b.N; i++ {
		ismatch := g.MatchString(`h1ealthcheck"`)
		if ismatch {
			aa++
		} else {
			bb++
		}
	}
	b.ReportAllocs()
	fmt.Printf("\n\ntrue:%d", aa)
	fmt.Printf("false:%d\n\n", bb)

}

func Benchmark_cre2(b *testing.B) {
	var aa, bb int
	g, cc := re2.MustCompile("health(C|c)heck")
	for i := 0; i < b.N; i++ {
		ismatch := g.MatchString(`h1ealthcheck"`)
		if ismatch {
			aa++
		} else {
			bb++
		}
	}
	cc.Close(g)
	b.ReportAllocs()
	fmt.Printf("\n\ntrue:%d", aa)
	fmt.Printf("false:%d\n\n", bb)

}
