package main

import (
	"fmt"
	"regexp"
)

const (
	IP_REGEX_PATTERN     = `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	DOMAIN_REGEX_PATTERN = `[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+\.?`
)

func IsIP(ip string) (b bool) {
	if m, _ := regexp.MatchString(IP_REGEX_PATTERN, ip); !m {
		//if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
		return false
	}
	return true
}
func IsDomain(domaon string) (b bool) {
	if m, _ := regexp.MatchString(DOMAIN_REGEX_PATTERN, domaon); !m {
		return false
	}
	return true
}

func main() {
	if IsDomain("www.baidu.com") {
		fmt.Printf("ok\n")
	}
	if !IsDomain("wwwbaiducom") {
		fmt.Printf("ok\n")
	}
}
