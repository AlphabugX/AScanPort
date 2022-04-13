package main

import (
	"AlphaNet/network"
	"flag"
	"log"
	"os"
	"strconv"
	"time"
)

var target = flag.String("t", "f5.ink", "127.0.0.1 or f5.ink")
//var message chan interface{}
func main() {
	flag.Parse()
	if len(os.Args)<=1 {
		flag.Usage()
	}else {
		AScan()
	}

}
func AScan()  {
	start := time.Now()
	log.Println("Start")
	log.Println(time.Since(start))
	for i := 1; i < 65536; i++ {
		network.Pool.Add(1)
		go network.ScanPort(*target,strconv.Itoa(i))
	}
	network.Pool.Wait()
	log.Println("over")
	log.Println(time.Since(start))
}