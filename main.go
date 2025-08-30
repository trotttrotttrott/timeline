package main

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Timeline struct {
	Events []Event
}

type Event struct {
	Date   time.Time
	Type   string
	String string
	Number float64
}

func main() {

	eventsPath, err := os.ReadFile(os.Args[1])

	if err != nil {
		fmt.Printf("Error reading YAML file: %v\n", err)
		return
	}

	var timeline Timeline

	err = yaml.Unmarshal(eventsPath, &timeline.Events)

	if err != nil {
		fmt.Printf("Error unmarshaling YAML: %v\n", err)
		return
	}

	fmt.Println(timeline)
}
