package main

import (
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
	Target  = flag.String("t", "f5.ink", "127.0.0.1 or f5.ink")
	Outfile = flag.String("o", "result_"+time.Now().Format("20060102_150405")+".txt", "result.txt")
	Format  = flag.String("d", "text", "text|json")
)

func main() {
	flag.Parse()
	if len(os.Args) <= 1 {
		flag.Usage()
	} else {
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
