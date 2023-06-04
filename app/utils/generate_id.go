package utils

import (
	"bytes"
	"hash/crc32"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func crc32Hash(s string) string {
	table := crc32.MakeTable(crc32.Castagnoli)
	hash := crc32.Checksum([]byte(s), table)
	return fmt.Sprintf("%08x", hash)
}

func getMacAddr() (addr string) {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			if i.Flags&net.FlagUp != 0 && !bytes.Equal(i.HardwareAddr, nil) != 0 {
				// Don't use random as we have a real address
				addr = i.HardwareAddr.String()
				break
			}
		}
	}
	return
}

func GenerateId() string {
	Time := strconv.Itoa(int(time.Now().UnixNano()))
	Pid := os.Getpid()
	StrPid := strconv.Itoa(Pid)
	var MAC = getMacAddr()

	var Mush = StrPid + "-" + MAC
	HashesMush := crc32Hash(Time) + crc32Hash(Mush)
	return HashesMush
}