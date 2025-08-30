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

func (ev *Event) Data() string {
	switch {
	case ev.String != nil:
		return *ev.String
	case ev.Number != nil:
		return fmt.Sprint(*ev.Number)
	}
	return ""
}

func (ev *Event) Metadata(evPrev *Event) string {
	str := ev.Date.Format("2006-01-02")
	if ev.Number != nil && evPrev != nil {
		diff := *ev.Number - *evPrev.Number
		str += fmt.Sprintf("; %.2f; %.1f%%", diff, (diff / *evPrev.Number)*100)
	}
	return str
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

		centerLine := "│"
		dataLines := make([]string, len(events))
		metadataLines := make([]string, len(events))
		for i := range dataLines[1:] {
			dataLines[i] = "│"
		}

		// intervals are the number of months between events
		intervals := make([]int, len(events)-1)
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
				centerLine += strings.Repeat("-", max(0, intervals[i]-1))
				centerLine += "│"
			}

			for j := range len(events) {

				// break if this line (and ones after) are done
				if j > len(events)-1-i {
					break
				}

				// is this the line the event should be displayed on?
				inverseIdx := len(events) - 1 - i
				isEvtLine := j == inverseIdx
				isNextEvtLine := j+1 == inverseIdx

				if len(intervals) > i && !isEvtLine {
					dataLines[j] += strings.Repeat(" ", max(0, intervals[i]-1))
					if !isNextEvtLine {
						dataLines[j] += "│"
					}
				}
				metadataLines[i] = dataLines[j]

				if isEvtLine {
					var evPrev *Event
					if i > 0 {
						evPrev = &events[i-1]
					}
					metadataLines[i] += fmt.Sprintf("┌ %s", event.Metadata(evPrev))
					dataLines[j] += fmt.Sprintf("└ %s", event.Data())
				}
			}
		}

		for _, line := range metadataLines {
			fmt.Println(line)
		}
		fmt.Println(centerLine)
		for _, line := range dataLines {
			fmt.Println(line)
		}
	}
}
