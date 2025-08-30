package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type Timeline struct {
	Events []Event
}

type Event struct {
	Date   time.Time
	Type   string
	String *string
	Number *float64
}

func (ev *Event) ToString() string {

	switch {
	case ev.String != nil:
		return *ev.String
	case ev.Number != nil:
		return fmt.Sprint(*ev.Number)
	}

	return ""
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

	sort.Slice(timeline.Events, func(i, j int) bool {
		return timeline.Events[i].Date.Before(timeline.Events[j].Date)
	})

	types := make(map[string][]Event)

	for _, e := range timeline.Events {
		types[e.Type] = append(types[e.Type], e)
	}

	for name, events := range types {

		fmt.Printf("\n%s\n\n", name)

		now := "now"
		events = append(events, Event{Date: time.Now(), String: &now})
		intervals := make([]int, len(events)-1)
		lines := make([]string, len(events)+1)
		for i := range lines[1:] {
			lines[i] = "│"
		}

		// intervals are the number of months between events
		for i := range len(intervals) {
			ev := events[i]
			nextEv := events[i+1]
			yearDiff := nextEv.Date.Year() - ev.Date.Year()
			monthDiff := int(nextEv.Date.Month()) - int(ev.Date.Month())
			months := yearDiff*12 + monthDiff
			intervals[i] = months
		}

		for i, event := range events {

			if i < len(intervals) {
				lines[0] += strings.Repeat("-", intervals[i]-1)
				lines[0] += "│"
			}

			for j := range lines[1:] {

				// break if this line (and ones after) are done
				if j > len(events)-1-i {
					break
				}

				// is this the line the event should be displayed on?
				inverseIdx := len(events) - 1 - i
				isEvtLine := j == inverseIdx
				isNextEvtLine := j+1 == inverseIdx

				line := &lines[j+1]
				if len(intervals) > i && !isEvtLine {
					*line += strings.Repeat(" ", intervals[i]-1)
					if !isNextEvtLine {
						*line += "│"
					}
				}

				if isEvtLine {
					*line += fmt.Sprintf("└ %s", event.ToString())
				}
			}
		}

		for i := range lines {
			fmt.Println(lines[i])
		}
	}
}
