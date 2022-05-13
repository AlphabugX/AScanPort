package network

import (
	"AScanPort/data"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

var (
	MaxThread       = 10000
	MaxHostThread   = 100
	Timeout         = 1
	MaxCheck        = 2
	ScanPortPoolMax = make(chan bool, MaxThread)
	ScanHostPoolMax = make(chan bool, MaxHostThread)
	neta            = net.Dialer{Timeout: time.Duration(Timeout) * time.Second}
	Pool            sync.WaitGroup
	PoolHost        sync.WaitGroup
	Port_count      = 0
)

func ScanPort(ip string, port string) {
	defer Pool.Done()
	//client, err := neta.Dial("tcp", ip+":"+port)
	for check := 0; check < MaxCheck; check++ {
		_, err := neta.Dial("tcp", ip+":"+port)
		if err == nil {
			Port_count += 1
			data.Result <- ip + ":" + port
			if !data.Silent {
				log.Printf("{\"%s\":\"%s\"}\n", ip, port)
			} else {
				fmt.Println(ip + ":" + port)
			}
			break
		}
	}
	<-ScanPortPoolMax
}
func HOSTScan(IP string) {
	defer PoolHost.Done()
	for i := 1; i < 65536; i++ {
		ScanPortPoolMax <- true
		Pool.Add(1)
		go ScanPort(IP, strconv.Itoa(i))
	}
	<-ScanHostPoolMax
}
func Go(Target string) {
	IPLIST := IPLIST(Target)
	for _, ip := range IPLIST {
		ScanHostPoolMax <- true
		PoolHost.Add(1)
		go HOSTScan(ip)
	}
	PoolHost.Wait()
	Pool.Wait()

}
