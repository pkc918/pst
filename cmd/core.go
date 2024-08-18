package cmd

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func isConnection(host string, port int, timeout time.Duration) bool {
	time.Sleep(time.Millisecond * 1)
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	if err != nil {
		fmt.Println(err)
		return false
	}
	_ = conn.Close()
	return true
}

func Scan() {
	var wg sync.WaitGroup
	timeout := time.Millisecond * time.Duration(Timeout)
	for port := StartPort; port <= EndPort; port++ {
		wg.Add(1)
		go func(p int) {
			connected := isConnection(Host, p, timeout)
			if connected {
				fmt.Printf("Connection successful - port: %d \n", p)
			}
			wg.Done()
		}(port)
	}
	wg.Wait()
	fmt.Println("scan done!")
}
