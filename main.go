package main

import (
	"AScanPort/config"
	"AScanPort/data"
	"AScanPort/network"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

var (
	Silent     bool
	Outfile    = flag.String("out", "", "result.txt")
	Target     = flag.String("h", "127.0.0.1", "Target:f5.ink|114.67.111.74|114.67.111.74/28|114.67.111.74-80|114.67.111.74-114.67.111.80|114.67.111.*")
	TargetFile = flag.String("hf", "", "Target:ip.txt")
	Thread     = flag.Int("t", 1000, "Maximum threads")
	Timeout    = flag.Int("time", 2, "timeout:3 seconds")
	Format     = flag.String("format", "text", "Result format: text=>ip:port,json=>{\"ip\":\"port\"}")
	MaxCheck   = flag.Int("check", 1, "MaxCheck:Connect check the maximum number")

	//Port_list = flag.Int("p", 0, "Port:80|80,443|1-1024")
)

func init() {
	flag.BoolVar(&Silent, "s", false, "silent mode")
}

func main() {

	flag.Parse()
	if len(os.Args) <= 1 {
		flag.Usage()
	} else {
		network.MaxThread = *Thread
		network.Timeout = *Timeout
		network.MaxCheck = *MaxCheck
		data.Silent = Silent
		config.Init()
		AScanPort()
	}
}
func AScanPort() {
	go scan_logs()
	var start time.Time
	if !data.Silent {
		fmt.Print("AScanPort (Version:1.0.4)\n")
		start = time.Now()
	}
	var Target_range interface{}
	if *TargetFile != "" {
		IPLIST, err := ioutil.ReadFile(*TargetFile)
		if err != nil {
			log.Fatalln("filename doesn't exist")
		}
		Target_range = strings.Split(strings.ReplaceAll(string(IPLIST), "\r", ""), "\n")
	} else {
		Target_range = *Target
	}
	switch Target_range.(type) {
	case string:
		network.Go(Target_range.(string))
	case []string:
		for _, item := range Target_range.([]string) {
			network.Go(item)
		}
	}
	if !data.Silent {
		log.Println("over", time.Since(start))
		log.Println("Open Ports:", network.Port_count)
	}
}
func scan_logs() {
	if *Outfile != "" {
		data.Save(*Outfile, *Format)
	}
}
