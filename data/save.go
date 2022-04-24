package data

import (
	"fmt"
	"os"
	"strings"
)

var (
	Result = make(chan string, 1234567)
)

func Save(filename string, format string) {
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
		tmp := strings.Split(<-Result, ":")
		fl.Write([]byte(fmt.Sprintf(data_format, tmp[0], tmp[1])))
	}
}
