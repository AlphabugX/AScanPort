package network

import (
	"bytes"
	"log"
	"math"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	// IP x.x.x.x
	IPLIST_TYPE_1 = regexp.MustCompile(`^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.)([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$`)

	// IP Range x.x.x.x-255
	IPLIST_TYPE_2 = regexp.MustCompile(`^(([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.)([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])-([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$`)

	// IP Range x.x.x.x-x.x.x.x
	IPLIST_TYPE_3 = regexp.MustCompile(`^(([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.)([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])-(([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.)(([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.)(([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.)([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$`)

	// IP Range x.x.x.x/x
	IPLIST_TYPE_4 = regexp.MustCompile(`^(([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.)([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/([1-9]|[1-2][0-9]|3[0-2])$`)

	// IP Range x.x.x.*
	IPLIST_TYPE_5 = regexp.MustCompile(`^(([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.)\*$`)
)

func ParseList(IPLIST string) interface{} {

	switch {
	case IPLIST_TYPE_1.MatchString(IPLIST):
		//log.Println("IPLIST_TYPE_1")
		return IPLIST
	case IPLIST_TYPE_2.MatchString(IPLIST):
		//log.Println("IPLIST_TYPE_2")
		var IPLIST_RANGE []string
		IPLIST_RANGE_Regexp := regexp.MustCompile(`([\d]+\.[\d]+\.[\d]+\.)([\d]+)-([\d]+)`)
		IPLIST_RANGEs := IPLIST_RANGE_Regexp.FindStringSubmatch(IPLIST)
		if len(IPLIST_RANGEs) < 4 {
			return nil
		}
		a, _ := strconv.Atoi(IPLIST_RANGEs[2])
		b, _ := strconv.Atoi(IPLIST_RANGEs[3])
		if a > b {
			return nil
		} else if b > 255 {
			b = 254
		}
		if len(IPLIST_RANGEs) == 4 {
			for i := a; i <= b; i++ {
				IPLIST_RANGE = append(IPLIST_RANGE, IPLIST_RANGEs[1]+strconv.Itoa(i))
			}
		}
		return IPLIST_RANGE
	case IPLIST_TYPE_3.MatchString(IPLIST):
		//log.Println("IPLIST_TYPE_3")
		var IPLIST_RANGE []string
		IPLIST_A := IP2Int(strings.Split(IPLIST, "-")[0])
		IPLIST_B := IP2Int(strings.Split(IPLIST, "-")[1])
		for i := IPLIST_A; i < IPLIST_B+1; i++ {
			IPLIST_RANGE = append(IPLIST_RANGE, Int2IP(i))
		}
		return IPLIST_RANGE
	case IPLIST_TYPE_4.MatchString(IPLIST):
		//log.Println("IPLIST_TYPE_4")
		var IPLIST_RANGE []string
		IPLIST_RANGE_Regexp := regexp.MustCompile(`([\d]+\.[\d]+\.[\d]+\.)([\d]+)/([\d]+)`)
		IPLIST_RANGEs := IPLIST_RANGE_Regexp.FindStringSubmatch(IPLIST)
		if len(IPLIST_RANGEs) < 4 {
			return nil
		}
		n, _ := strconv.Atoi(IPLIST_RANGEs[3])
		IP := IPLIST_RANGEs[1] + IPLIST_RANGEs[2]
		CIDR_n := 32 - n
		Subnet_length := int(math.Pow(2, float64(CIDR_n)))
		CIDR_count := (IP2Int(IP) - IP2Int(IPLIST_RANGEs[1]+"0")) / Subnet_length
		log.Println("CIDR_count", CIDR_count)

		IPLIST_A2int := int64(IP2Int(IPLIST_RANGEs[1] + "0"))
		IPLIST_A2bin := strconv.FormatInt(IPLIST_A2int, 2)[:len(strconv.FormatInt(IPLIST_A2int, 2))-CIDR_n] + strings.Repeat("0", CIDR_n)
		IPLIST_A, _ := strconv.ParseInt(IPLIST_A2bin, 2, 64)
		IPLIST_A = IPLIST_A + int64(CIDR_count*Subnet_length)
		IPLIST_B := int(IPLIST_A) + Subnet_length - 1
		for i := int(IPLIST_A + 1); i < IPLIST_B; i++ {
			IPLIST_RANGE = append(IPLIST_RANGE, Int2IP(i))
		}
		return IPLIST_RANGE
	case IPLIST_TYPE_5.MatchString(IPLIST):
		//log.Println("IPLIST_TYPE_5")
		var IPLIST_RANGE []string
		IPLIST_RANGE_Regexp := regexp.MustCompile(`([\d]+\.[\d]+\.[\d]+\.)\*`)
		IPLIST_RANGEs := IPLIST_RANGE_Regexp.FindStringSubmatch(IPLIST)
		if len(IPLIST_RANGEs) < 2 {
			return nil
		}
		for i := 0; i < 256; i++ {
			IPLIST_RANGE = append(IPLIST_RANGE, IPLIST_RANGEs[1]+strconv.Itoa(i))
		}
		return IPLIST_RANGE
	default:
		addr, err := net.ResolveIPAddr("ip", IPLIST)
		if err != nil {
			log.Println("Resolvtion error", err.Error())
			os.Exit(1)
		}
		return addr.String()
	}
	return nil
}
func IP2Int(ip string) int {
	ipSegs := strings.Split(ip, ".")
	var ipInt int = 0
	var pos uint = 24
	for _, ipSeg := range ipSegs {
		tempInt, _ := strconv.Atoi(ipSeg)
		tempInt = tempInt << pos
		ipInt = ipInt | tempInt
		pos -= 8
	}
	return ipInt
}
func Int2IP(ipInt int) string {
	ipSegs := make([]string, 4)
	var ipSegs_len = len(ipSegs)
	buffer := bytes.NewBufferString("")
	for i := 0; i < ipSegs_len; i++ {
		tempInt := ipInt & 0xFF
		ipSegs[ipSegs_len-i-1] = strconv.Itoa(tempInt)
		ipInt = ipInt >> 8
	}
	for i := 0; i < ipSegs_len; i++ {
		buffer.WriteString(ipSegs[i])
		if i < ipSegs_len-1 {
			buffer.WriteString(".")
		}
	}
	return buffer.String()
}
func IPLIST(ip string) []string {
	IP_check := ParseList(ip)
	switch IP_check.(type) {
	case string:
		return []string{IP_check.(string)}
	case []string:
		return IP_check.([]string)
	}
	return nil
}
