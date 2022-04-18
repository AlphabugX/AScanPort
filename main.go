package main

import (
	"AlphaNet/config"
	"AlphaNet/data"
	"AlphaNet/network"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	Target   = flag.String("h", "f5.ink", "127.0.0.1 or f5.ink")
	Thread   = flag.Int("t", 14000, "Maximum threads")
	Timeout  = flag.Int("time", 3, "timeout:3 seconds")
	Outfile  = flag.String("o", "result_"+time.Now().Format("20060102_150405")+".txt", "result.txt")
	Format   = flag.String("d", "text", "Result format: text=>ip:port,json=>{\"ip\":\"port\"}")
	MaxCheck = flag.Int("check", 1, "MaxCheck:Connect check the maximum number")
)

func main() {
	flag.Parse()
	if len(os.Args) <= 1 {
		flag.Usage()
	} else {
		network.MaxThread = *Thread
		network.Timeout = *Timeout
		network.MaxCheck = *MaxCheck
		config.Init()
		AScan()
	}
}
func AScan() {
	fmt.Print("AScanPort (Version:1.0.1)\n")
	start := time.Now()
	for i := 1; i < 65536; i++ {
		network.Pool.Add(1)
		go network.ScanPort(*Target, strconv.Itoa(i))
	}
	network.Pool.Wait()

	log.Println("over", time.Since(start))
	log.Println("Open Ports:", network.Port_count)
	data.Save(*Outfile, *Format, *Target)
}
