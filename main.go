// ping project main.go
package main

import (
	//"errors"
	"flag"
	"fmt"
	"net"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func portisalive(ip string) {
	fmt.Printf("connecting %s...\n", ip)
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		fmt.Printf("connet %s is failed!\n", ip)
	} else {
		fmt.Printf("connet %s is sucess!\n", ip)
		conn.Close()
	}
	wg.Done()
}

func main() {
	flag.CommandLine.Usage = func() {
		fmt.Println("Usage:\n  portping ip:port,ip1:port,ip2:port...\nexample:\n  portping 127.0.0.1:22,10.0.0.1:23")
	}
	flag.Parse()
	ipstr := flag.Arg(0)

	if ipstr == "" || ipstr == "-h" || ipstr == "--help" {
		fmt.Println("Usage:\n  portping ip:port,ip1:port,ip2:port...\nexample:\n  portping 127.0.0.1:22,10.0.0.1:23")
		return
	}
	ipstr = strings.TrimSuffix(ipstr, `"`)
	ipstr = strings.TrimPrefix(ipstr, `"`)
	ipstr = strings.TrimSuffix(ipstr, `'`)
	ipstr = strings.TrimPrefix(ipstr, `'`)
	ips := strings.Split(ipstr, ",")
	for _, v := range ips {
		wg.Add(1)
		go portisalive(v)
	}
	wg.Wait()
}
