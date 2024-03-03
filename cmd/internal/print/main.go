package main

import (
	"encoding/json"
	"fmt"
	trucktel "github.com/wojcikt/trucktel-go"
	"log"
)

func main() {
	tel, err := trucktel.Open()
	if err != nil {
		log.Fatalln(err)
	}
	defer func(tel trucktel.Telemetry) {
		err := tel.Close()
		if err != nil {
			log.Println(err)
		}
	}(tel)

	var values trucktel.Values
	if err = tel.Read(&values); err != nil {
		log.Fatalln(err)
	}

	result, err := json.Marshal(values)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(result))
}
