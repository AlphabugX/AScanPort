package network

import (
	"encoding/json"
	"log"
	"net"
	"sync"
	"time"
)

var (
	Scan_Pool_Max = make(chan bool,8000)
	neta = net.Dialer{Timeout: 2*time.Second}
 	Pool sync.WaitGroup
)

func ScanPort(ip string,port string) {
	defer Pool.Done()
	Scan_Pool_Max <- true
	_, err := neta.Dial("tcp", ip+":"+port)
	if err == nil {
		reslut := map[string]interface{}{
			"ip":   ip,
			"port": port,
		}
		data, _ := json.Marshal(reslut)
		log.Println(string(data))
	}
	<- Scan_Pool_Max

}