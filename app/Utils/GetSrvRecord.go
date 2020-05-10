package Utils

import (
	"fmt"
	"net"

	//"encoding/json"
	"runtime"
	"strconv"
	"strings"

	"github.com/fatih/structs"
)

func GetRecordSrv(service string) string {
	cName, addrs, err := net.LookupSRV("", "", service)
	if err != nil {
		return ""
	}
	if cName != "" {
		fmt.Println(cName)
	}
	dat1 := structs.Map(addrs[0])
	return "http://" + strings.Trim(dat1["Target"].(string), ".") + ":" + strconv.Itoa(int(dat1["Port"].(uint16)))

}

func GetNumCpu() int {
	num := runtime.NumCPU()
	return num
}
