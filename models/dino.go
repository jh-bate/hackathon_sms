package models

import (
	"strings"
)

type (
	Common struct {
		Type     string `json:"type"`
		DeviceId string `json:"deviceId"`
		Source   string `json:"source"`
		Time     string `json:"time"`
	}

	BgEvent struct {
		Common
		Units string `json:"units"`
		Value string `json:"value"`
	}

	FoodEvent struct {
		Common
		Carbs string `json:"carbs"`
	}

	BasalEvent struct {
		Common
		DeliveryType string `json:"deliveryType"`
		Value        string `json:"value"`
	}
	BolusEvent struct {
		Common
		SubType string `json:"subType"`
		Value   string `json:"value"`
	}
)

const (
	BG    = "BG="
	CARB  = "CHO="
	BASAL = "SA="
	BOLUS = "LA="
)

func Translate(smsString string, date string) []interface{} {

	var events []interface{}

	raw := strings.Split(smsString, " ")

outer:
	for i := range raw {
		switch {
		case strings.Index(raw[i], BG) != -1:
			events = append(events, makeBg(raw[i], date))
			break
		case strings.Index(raw[i], CARB) != -1:
			events = append(events, makeCarb(raw[i], date))
			break
		case strings.Index(raw[i], BASAL) != -1:
			events = append(events, makeBasal(raw[i], date))
			break
		case strings.Index(raw[i], BOLUS) != -1:
			events = append(events, makeBolus(raw[i], date))
			break
		default:
			events = append(events, makeNote(smsString, date))
			break outer
		}
	}
	return events
}

func makeBg(bgString, date string) *BgEvent {
	bg := strings.Split(bgString, BG)
	return &BgEvent{Common: Common{Type: "smbg", Source: "dinojr", Time: date}, Units: "", Value: bg[1]}
}

func makeNote(noteString, date string) *Common {
	return &Common{Type: "note", Source: "dinojr", Time: date}
}

func makeCarb(carbString, date string) *FoodEvent {
	carb := strings.Split(carbString, CARB)
	return &FoodEvent{Common: Common{Type: "food", Source: "dinojr", Time: date}, Carbs: carb[1]}
}

func makeBolus(bolusString, date string) *BolusEvent {
	bolus := strings.Split(bolusString, BOLUS)
	return &BolusEvent{Common: Common{Type: "bolus", Source: "dinojr", Time: date}, SubType: "injected", Value: bolus[1]}
}

func makeBasal(basalString, date string) *BasalEvent {
	basal := strings.Split(basalString, BASAL)
	return &BasalEvent{Common: Common{Type: "basal", Source: "dinojr", Time: date}, DeliveryType: "injected", Value: basal[1]}
}
