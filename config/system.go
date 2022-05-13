package config

func Init() {
	//if runtime.GOOS != "windows" {
	//	var rlim syscall.Rlimit
	//	rlim.Max = 999999
	//	rlim.Cur = 999999
	//	err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rlim)
	//	if err != nil {
	//		fmt.Errorf("SET Rlimit:%s", err)
	//		os.Exit(1)
	//	}
	//	err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlim)
	//	if err != nil {
	//		fmt.Errorf("GET Rlimit:%s", err)
	//		os.Exit(1)
	//	}
	//	//fmt.Printf("ENV RLIMIT_NOFILE : %+v\n", rlim)
	//
	//}
}
