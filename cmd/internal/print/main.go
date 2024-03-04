package main

import (
	"encoding/json"
	"fmt"
	"github.com/wojcikt/trucktel-go/internal/shm"
	"log"
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

	var values shm.Values
	if err = mem.Read(&values); err != nil {
		log.Fatalln(err)
	}

	result, err := json.Marshal(values)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(result))
}
