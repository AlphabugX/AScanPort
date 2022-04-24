package network

import (
	"AlphaNet/data"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

var (
	MaxThread          = 10000
	Timeout            = 3
	MaxCheck           = 3
	Scan_Port_Pool_Max = make(chan bool, MaxThread)
	neta               = net.Dialer{Timeout: time.Duration(Timeout) * time.Second}
	Pool               sync.WaitGroup
	Port_count         = 0
)

func ScanPort(ip string, port string) {
	defer Pool.Done()
	Scan_Port_Pool_Max <- true
	//client, err := neta.Dial("tcp", ip+":"+port)
	for check := 0; check < MaxCheck; check++ {
		_, err := neta.Dial("tcp", ip+":"+port)
		if err == nil {
			Port_count += 1
			data.Result <- ip + ":" + port
			log.Printf("{\"%s\":\"%s\"}", ip, port)
			break
		}
	}
	<-Scan_Port_Pool_Max
}
func HOSTScan(IP string) {
	defer Pool.Done()
	for i := 1; i < 65536; i++ {
		Pool.Add(1)
		go ScanPort(IP, strconv.Itoa(i))
	}
}
func Go(Target string) {
	fmt.Print("AScanPort (Version:1.0.1)\n")
	IPLIST := IPLIST(Target)
	start := time.Now()
	for _, ip := range IPLIST {
		Pool.Add(1)
		go HOSTScan(ip)
	}
	Pool.Wait()
	log.Println("over", time.Since(start))
	log.Println("Open Ports:", Port_count)
}
