package models

import (
	"log"
	"strings"
)

type (
	Event struct {
		Type     string `json:"type"`
		DeviceId string `json:"deviceId"`
		Source   string `json:"source"`
		TZOffset string `json:"timezoneOffset"`
		Units    string `json:"units"`
		Value    string `json:"value"`
	}
)

const (
	BG    = "BG="
	CARB  = "CHO="
	BASAL = "SA="
	BOLUS = "LA="
)

func Translate(smsString string) []*Event {

	var events []*Event

	raw := strings.Split(smsString, " ")

	log.Printf("split as %v ", raw[0])

outer:
	for i := range raw {
		switch {
		case strings.Index(raw[i], BG) != -1:
			events = append(events, makeBg(raw[i]))
			break
		case strings.Index(raw[i], CARB) != -1:
			events = append(events, makeCarb(raw[i]))
			break
		case strings.Index(raw[i], BASAL) != -1:
			events = append(events, makeBasal(raw[i]))
			break
		case strings.Index(raw[i], BOLUS) != -1:
			events = append(events, makeBolus(raw[i]))
			break
		default:
			events = append(events, makeNote(smsString))
			break outer
		}
	}
	return events
}

func makeBg(bgString string) *Event {
	return &Event{Type: "", Source: "", Units: "", Value: ""}
}

func makeNote(noteString string) *Event {
	return &Event{Type: "", Source: "", Units: "", Value: ""}
}

func makeCarb(carbString string) *Event {
	return &Event{Type: "", Source: "", Units: "", Value: ""}
}

func makeBolus(bolusString string) *Event {
	return &Event{Type: "", Source: "", Units: "", Value: ""}
}

func makeBasal(basalString string) *Event {
	return &Event{Type: "", Source: "", Units: "", Value: ""}
}
