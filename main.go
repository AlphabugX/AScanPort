package main

import (
	"AlphaNet/config"
	"AlphaNet/data"
	"AlphaNet/network"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

var (
	Target     = flag.String("h", "f5.ink", "Target:f5.ink|114.67.111.74|114.67.111.74/28|114.67.111.74-80|114.67.111.74-114.67.111.80|114.67.111.*")
	TargetFile = flag.String("hf", "ip.txt", "Target:ip.txt")
	Thread     = flag.Int("t", 10000, "Maximum threads")
	Timeout    = flag.Int("time", 1, "timeout:3 seconds")
	Outfile    = flag.String("out", "text.txt", "result.txt")
	Format     = flag.String("format", "text", "Result format: text=>ip:port,json=>{\"ip\":\"port\"}")
	MaxCheck   = flag.Int("check", 2, "MaxCheck:Connect check the maximum number")
	Silent     = flag.Bool("s", false, "silent mode")
	//Port_list = flag.Int("p", 0, "Port:80|80,443|1-1024")
)

func main() {

	flag.Parse()
	if len(os.Args) <= 1 {
		flag.Usage()
	} else {
		network.MaxThread = *Thread
		network.Timeout = *Timeout
		network.MaxCheck = *MaxCheck
		data.Silent = *Silent
		config.Init()
		AScanPort()
	}
}
func AScanPort() {
	go Log_save()
	var start time.Time
	if !data.Silent {
		fmt.Print("AScanPort (Version:1.0.1)\n")
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
func Log_save() {
	if *Outfile != "" {
		data.Save(*Outfile, *Format)
	}
}
