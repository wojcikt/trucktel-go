package main

import (
	"github.com/wojcikt/trucktel-go/internal/shm"
	"io"
	"log"
	"os"
)

func main() {
	mem, err := shm.Open(shm.DefaultFileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(mem shm.SharedMemory) {
		err := mem.Close()
		if err != nil {
			log.Println(err)
		}
	}(mem)

	fileName := "data.bin"
	dst, err := os.Create(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(dest *os.File) {
		_ = dest.Close()
	}(dst)

	written, err := io.Copy(dst, mem.Reader())
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("written %d bytes to %s\n", written, fileName)
}
