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
	Scan_Pool_Max = make(chan bool, 14000)
	neta          = net.Dialer{Timeout: 3 * time.Second}
	Pool          sync.WaitGroup
	Port_count    = 0
)

func ScanPort(ip string, port string) {
	defer Pool.Done()
	Scan_Pool_Max <- true
	//client, err := neta.Dial("tcp", ip+":"+port)
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
	}
	<-Scan_Pool_Max
}
