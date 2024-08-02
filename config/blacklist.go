package config

import "strings"

func CheckIP(addr string) bool{
	arr := []string{"182.18.2.1"}
	addrs := strings.Split(addr, ":")

	for i:=0; i<len(arr); i++{
		if arr[i]==addrs[0] {
			return true
		}
	}
	return false
}