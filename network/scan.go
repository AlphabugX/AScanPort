package network

import (
	"AScanPort/data"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
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
	client          net.Conn
	Service         = false
)

func ScanPort(ip string, port string) {
	defer Pool.Done()
	//log.Println(ip,port)
	var err error
	//client, err := neta.Dial("tcp", ip+":"+port)
	for check := 0; check < MaxCheck; check++ {
		client, err = neta.Dial("tcp", ip+":"+port)
		if err == nil {
			if Service {
				if Service_return := Service_probes(client); Service_return != nil {
					Port_count += 1
					data.Result <- ip + ":" + port + ":" + Service_return.(string)
					if !data.Silent {
						log.Printf("{\"%s\":\"%s\",\"service\":\"%s\"}\n", ip, port, Service_return)
					} else {
						fmt.Println(ip + ":" + port + ":" + Service_return.(string))
					}
					break
				}
			} else {
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
	if len(IPLIST) == 1 {
		MaxHostThread = 1
	}
	for _, ip := range IPLIST {
		ScanHostPoolMax <- true
		PoolHost.Add(1)
		go HOSTScan(ip)
	}
	PoolHost.Wait()
	Pool.Wait()

}
func Service_probes(client net.Conn) interface{} {
	client.SetReadDeadline(time.Now().Add(time.Duration(Timeout+2) * time.Second))
	defer client.Close()
	_, err := client.Write([]byte("HEAD / HTTP/1.1\n\n"))
	if err != nil {
		return nil
	}
	buf := make([]byte, 1024)
	n, err := client.Read(buf[:])
	if err != nil {
		return nil
	}
	//log.Println(client.RemoteAddr().String())
	service_data := string(buf[:n])
	service_data = strings.ToLower(service_data)
	switch {
	case strings.Contains(service_data, "http"):
		return "http"
	case strings.Contains(service_data, "ssh"):
		return "ssh"
	case strings.Contains(service_data[15:], "mariadb") || strings.Contains(service_data[15:], "mysql") || strings.Contains(service_data[15:], "native_password"):
		return "mysql"
	case strings.Contains(service_data, "ftp"):
		return "ftp"
	default:
		return "unknown"
	}

}
