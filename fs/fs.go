package fs

import (
	"io"
	"os"
	"path/filepath"
	"strconv"
)

func OutputStream(port int) io.WriteCloser {
	home, _ := os.UserHomeDir()
	port_str := strconv.Itoa(port)
	dir := filepath.Join(home, ".local/state/blockchain_from_scratch")
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		panic(err)
	}
	file := filepath.Join(dir, "data_"+port_str+".json")
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	return f
}
