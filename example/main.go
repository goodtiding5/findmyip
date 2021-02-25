package main

import "github.com/goodtiding5/findmyip"

func main() {
	ips, err := findmyip.FindMyIp()
	if err != nil {
		panic(err.Error())
	}
	print(ips[0])
}

