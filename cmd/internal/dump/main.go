package main

import (
	trucktel "github.com/wojcikt/trucktel-go"
	"github.com/wojcikt/trucktel-go/internal"
	"io"
	"log"
	"os"
)

func main() {
	shm, err := internal.OpenShm(trucktel.DefaultMmfName)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(shm internal.Shm) {
		err := shm.Close()
		if err != nil {
			log.Println(err)
		}
	}(shm)

	fileName := "data.bin"
	dst, err := os.Create(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(dest *os.File) {
		_ = dest.Close()
	}(dst)

	written, err := io.Copy(dst, shm.Reader())
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("written %d bytes to %s\n", written, fileName)
}
