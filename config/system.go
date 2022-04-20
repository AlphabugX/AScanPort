package config

import (
	"runtime"
)

func Init() {
	SystemType := runtime.GOOS
	if SystemType != "windows" {
		//var rlim syscall.Rlimit
		//rlim.Max = 65535
		//rlim.Cur = 65535
		//err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rlim)
		//if err != nil {
		//	fmt.Errorf("SET Rlimit:%s", err)
		//	os.Exit(1)
		//}
		//err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlim)
		//if err != nil {
		//	fmt.Errorf("GET Rlimit:%s", err)
		//	os.Exit(1)
		//}
		//fmt.Printf("ENV RLIMIT_NOFILE : %+v\n", rlim)

	}
}
