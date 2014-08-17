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

func Translate(smsString string) []interface{} {

	var events []interface{}

	raw := strings.Split(smsString, " ")

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

func makeBg(bgString string) *BgEvent {
	bg := strings.Split(bgString, BG)
	return &BgEvent{Common: Common{Type: "smbg", Source: "dinojr"}, Units: "", Value: bg[1]}
}

func makeNote(noteString string) *Common {
	return &Common{Type: "note", Source: "dinojr"}
}

func makeCarb(carbString string) *FoodEvent {
	carb := strings.Split(carbString, CARB)
	return &FoodEvent{Common: Common{Type: "food", Source: "dinojr"}, Carbs: carb[1]}
}

func makeBolus(bolusString string) *BolusEvent {
	bolus := strings.Split(bolusString, BOLUS)
	return &BolusEvent{Common: Common{Type: "bolus", Source: "dinojr"}, SubType: "injected", Value: bolus[1]}
}

func makeBasal(basalString string) *BasalEvent {
	basal := strings.Split(basalString, BASAL)
	return &BasalEvent{Common: Common{Type: "basal", Source: "dinojr"}, DeliveryType: "injected", Value: basal[1]}
}
