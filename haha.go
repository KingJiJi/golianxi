package main

import (
	"fmt"
	"net"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func Log(name string) {
	fmt.Printf(" %s\n ", name)
}

func Demo() string {
	return "demo"
}

func TestEpmeyStr() {
	haha, err := strconv.Atoi("")
	if err != nil {
		fmt.Printf("err %v\n", err)
	} else {
		fmt.Printf(" empty string atoi %d \n", haha)
	}
}

type PWD struct {
	U string
	P string
}

var NilPwd PWD = PWD{U: "XMAN", P: "love letter"}

func main() {
	strings.Split("hahahahah hahah hahah", " ")
	fmt.Printf("helloworld haha\n")
	fmt.Printf("hahahah\n")
	Log("haha")

	mydemo := ""
	if mydemo = Demo(); mydemo == "demo" {
		fmt.Printf("hahaah demo ok\n")
	}

	TestEpmeyStr()

	if strings.ContainsAny("123456", "1df") {
		fmt.Printf(" 123456 ContainsAny  1df\n")
	}
	if strings.ContainsAny("123456", "df") {
		fmt.Printf(" 123456 ContainsAny  df\n")
	}

	var (
		a interface{}
		b interface{}
	)

	if a == b {
		fmt.Printf("default interface is same\n")
	}

	c := 5
	d := 6
	e := 5

	f := &c
	g := &e

	a = c
	b = d
	if reflect.DeepEqual(a, b) {
		fmt.Printf(" c5 d6  interface is same\n")
	}

	b = e
	if a == b {
		//if reflect.DeepEqual(a, b) {   //true
		fmt.Printf(" c5 e5  interface is same\n")
	} else {
		fmt.Printf(" c5 e5  interface is not same\n")

	}

	a = f
	b = g
	if a == b { //false
		//if reflect.DeepEqual(a, b) { //true
		fmt.Printf(" a=&c  b=&e   c=e=5  interface is same\n")
	}

	if strings.HasPrefix("heha", "heha") {
		fmt.Printf("hehahaha  has haha prefix \n\n\n")
	}

	var intface interface{}
	data := []interface{}{1, 2, 3, 4, 5, 6, 7}
	intface = data
	fmt.Printf("%d, %s\n", intface, reflect.TypeOf(intface).String())
	ip, port, err := net.SplitHostPort("1.2.3.4:80")
	if err != nil {
		fmt.Printf(" err :%v")
	} else {
		fmt.Println(ip, port)
	}
	ip, port, err = net.SplitHostPort("1.2.3.4.5:80")
	if err != nil {
		fmt.Printf(" err :%v\n")
	} else {
		fmt.Println(ip, port)
	}
	ip, port, err = net.SplitHostPort("hello.go:80")
	if err != nil {
		fmt.Printf(" err :%v\n")
	} else {
		fmt.Println(ip, port)
	}
	ip, port, err = net.SplitHostPort("www.baidu.com:80")
	if err != nil {
		fmt.Printf(" err :%v\n")
	} else {
		fmt.Println(ip, port)
	}
	iip, port, err := net.SplitHostPort("www.baidu.co")
	if err != nil {
		fmt.Printf(" err :%v\n", err)
	} else {
		fmt.Println(iip, port)
	}

	hapwd := PWD{U: "XMAN", P: "love letter"}
	if hapwd == NilPwd {
		fmt.Printf(" struct value is equal\n")
	} else {
		fmt.Printf(" struct value is no equal\n")

	}
	{
		fmt.Printf("[%q]\n", strings.Trim("alter table `Achtung! Achtung` add column `haha` int", "`"))
		fmt.Printf("[%q]\n", strings.Replace("alter table `Achtung! Achtung` add column `haha` int", "`", "", -1))
	}

	asb := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	asc := asb[5:6]
	fmt.Printf("%v\n", len(asc))
	mytime := time.Now()
	fmt.Printf("%s\n", mytime)
	fmt.Println(mytime.Format("2016-03-24 00:00:00"))
	fmt.Println(mytime.Format("20060102150405"))

	WEB_FORMAT := "20060102150405"
	thaha, terr := time.Parse(WEB_FORMAT, "20170809123456")
	if terr != nil {
		fmt.Printf("err:%s\n", terr.Error())
	} else {
		fmt.Printf("20170809123456:%v\n", thaha)
	}
	fmt.Printf("%d\n", mytime.Unix())
	fmt.Printf("%d\n", mytime.UnixNano())
	NGINX_FORMAT := "01/Jul/2006:15:04:05 -0700"
	thaha, terr = time.Parse(NGINX_FORMAT, "13/Jun/2017:16:31:13 +0800")
	if terr != nil {
		fmt.Printf("err:%s\n", terr.Error())
	} else {
		fmt.Printf("13/Jun/2017:16:31:13 +0800 :%v\n", thaha)
	}

	imytime := time.Now().Format("2006-01-02T15:04:05.000Z")
	fmt.Printf("mytime:%s\n", imytime)
}
