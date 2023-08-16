package debinterface

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

func readFileByLine(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	result := make([]string, 0)
	buf := bufio.NewReader(f)

	for {
		line, _, err := buf.ReadLine()
		l := strings.TrimSpace(string(line))
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		result = append(result, l)
	}
	return result, nil
}

func split(s string) []string {
	return deleteEmpty(strings.Split(s, " "))
}

func deleteEmpty(slice []string) []string {
	j := 0
	for _, v := range slice {
		if v != "" {
			slice[j] = v
			j++
		}
	}
	return slice[:j]
}

func ipv4Mask(m []byte) net.IP {
	return net.ParseIP(fmt.Sprintf("%d.%d.%d.%d", m[0], m[1], m[2], m[3]))
}
