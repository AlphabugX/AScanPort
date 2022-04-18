package network

import (
	"AlphaNet/data"
	"encoding/json"
	"log"
	"net"
	"sync"
	"time"
)

var (
	MaxThread     = 14000
	Timeout       = 3
	MaxCheck      = 3
	Scan_Pool_Max = make(chan bool, MaxThread)
	neta          = net.Dialer{Timeout: time.Duration(Timeout) * time.Second}
	Pool          sync.WaitGroup
	Port_count    = 0
)

func ScanPort(ip string, port string) {
	defer Pool.Done()
	Scan_Pool_Max <- true
	//client, err := neta.Dial("tcp", ip+":"+port)
	for check := 0; check < MaxCheck; check++ {
		_, err := neta.Dial("tcp", ip+":"+port)
		if err == nil {
			Port_count += 1
			data.Result <- port
			reslut := map[string]interface{}{
				"ip":   ip,
				"port": port,
			}
			data, _ := json.Marshal(reslut)
			log.Println(string(data))
			break
		}
	}
	<-Scan_Pool_Max
}
