package fs

import (
	"io"
	"os"
	"path/filepath"
	"strconv"
)

func OutputStream(port int) io.WriteCloser {
	file := getDataFilePath(port)
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	return f
}

func ReadStream(port int) io.ReadCloser {
	file := getDataFilePath(port)
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	return f
}

func getDataFilePath(port int) string {
	home, _ := os.UserHomeDir()
	port_str := strconv.Itoa(port)
	dir := filepath.Join(home, ".local/state/blockchain_from_scratch")
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		panic(err)
	}
	return filepath.Join(dir, "data_"+port_str+".json")
}
