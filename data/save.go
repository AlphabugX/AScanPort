package data

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	Result = make(chan string, 99999999)
	Silent = false
)

func Save(filename string, format string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
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
	defer file.Close()
	write := bufio.NewWriter(file)
	for {
		tmp_string := <-Result
		tmp := strings.Split(tmp_string, ":")
		writeline := fmt.Sprintf(data_format, tmp[0], tmp[1])
		write.WriteString(writeline)
		write.Flush()
	}

}
