package data

import (
	"fmt"
	"os"
)

var (
	Result = make(chan string, 65535)
)

func Save(filename string, format string, Target string) {
	fl, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	data_format := ""
	switch format {
	case "json":
		data_format = "{\"%s\":\"%s\"}\n"
	case "text":
		data_format = "%s:%s\n"

	}
	defer fl.Close()
	for {
		if len(Result) == 0 {
			break
		}
		fl.Write([]byte(fmt.Sprintf(data_format, Target, <-Result)))
	}
	//fl.Write([]byte(""))
}
