package main

import (
	"AlphaNet/config"
	"AlphaNet/data"
	"AlphaNet/network"
	"flag"
	"os"
)

var (
	Target   = flag.String("h", "f5.ink", "Target:f5.ink|114.67.111.74|114.67.111.74/28|114.67.111.74-80|114.67.111.74-114.67.111.80|114.67.111.*")
	Thread   = flag.Int("t", 10000, "Maximum threads")
	Timeout  = flag.Int("time", 2, "timeout:3 seconds")
	Outfile  = flag.String("out", "", "result.txt")
	Format   = flag.String("format", "text", "Result format: text=>ip:port,json=>{\"ip\":\"port\"}")
	MaxCheck = flag.Int("check", 1, "MaxCheck:Connect check the maximum number")
	//Port_list = flag.Int("p", 0, "Port:80|80,443|1-1024")
)

func main() {
	//network.MaxThread = *Thread
	//network.Timeout = *Timeout
	//network.MaxCheck = *MaxCheck
	//AScanPort()
	//os.Exit(1)
	flag.Parse()
	if len(os.Args) <= 1 {
		flag.Usage()
	} else {
		network.MaxThread = *Thread
		network.Timeout = *Timeout
		network.MaxCheck = *MaxCheck
		config.Init()
		AScanPort()
	}
}
func AScanPort() {
	network.Go(*Target)
	if *Outfile != "" {
		data.Save(*Outfile, *Format)
	}
}
