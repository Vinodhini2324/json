package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	const jsonStream = `
	{"Name": "Ni", "Text": "Hai"}
	{"Name": "vino", "Text": "How are you?"}
	{"Name": "Ni", "Text": "Gud, you?"}
	{"Name": "vino", "Text": "Fine,where are you going?"}
	{"Name": "Ni", "Text": "Going to my home!"}
`
	type Message struct {
		Name, Text string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}
