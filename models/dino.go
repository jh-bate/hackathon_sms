package models

import (
	"log"
	"strconv"
	"strings"
)

type (
	//common fields for event
	Common struct {
		Type     string `json:"type"`
		DeviceId string `json:"deviceId"`
		Source   string `json:"source"`
		Time     string `json:"time"`
	}
	BgEvent struct {
		Common
		Value float64 `json:"value"`
	}
	FoodEvent struct {
		Common
		Carbs float64 `json:"carbs"`
	}
	BasalEvent struct {
		Common
		DeliveryType string  `json:"deliveryType"`
		Value        float64 `json:"value"`
		Duration     int     `json:"duration"`
		Insulin      string  `json:"insulin"`
	}
	BolusEvent struct {
		Common
		SubType string  `json:"subType"`
		Value   float64 `json:"value"`
		Insulin string  `json:"insulin"`
	}
	NoteEvent struct {
		Common
		CreatorId string `json:"creatorId"`
		Text      string `json:"text"`
	}
)

const (
	TIME     = "T="
	ACTIVITY = "A="
	BG       = "B="
	CARB     = "C="
	BASAL    = "S="
	BOLUS    = "L="
	NOTE     = "N="
	LOW      = "#L"
	MMOLL    = "mmol/L"
)

func Translate(smsMessage string, date, device string) []interface{} {

	var events []interface{}

	raw := strings.Split(smsMessage, " ")

outer:
	for i := range raw {
		switch {
		case strings.Index(strings.ToUpper(raw[i]), BG) != -1:
			events = append(events, makeBg(raw[i], date, device))
			break
		case strings.Index(strings.ToUpper(raw[i]), CARB) != -1:
			events = append(events, makeCarb(raw[i], date, device))
			break
		case strings.Index(strings.ToUpper(raw[i]), BASAL) != -1:
			events = append(events, makeBasal(raw[i], date, device))
			break
		case strings.Index(strings.ToUpper(raw[i]), BOLUS) != -1:
			events = append(events, makeBolus(raw[i], date, device))
			break
		case strings.Index(strings.ToUpper(raw[i]), LOW) != -1:
			//hard code 'LOW'
			events = append(events, makeBg("B=3.9", date, device))
			break
		case strings.Index(strings.ToUpper(raw[i]), ACTIVITY) != -1:
			//hard code 'LOW'
			log.Panicln("Will be an activity ", raw[i])
			break
		case strings.Index(strings.ToUpper(raw[i]), NOTE) != -1:
			//hard code 'LOW'
			log.Panicln("Will be a note ", raw[i])
			break
		default:
			log.Panicln("Will be a note ", smsMessage)
			break outer
		}
	}
	return events
}

func makeBg(bgString, date, device string) *BgEvent {
	bg := strings.Split(bgString, BG)
	bgVal, _ := strconv.ParseFloat(bg[1], 64)

	return &BgEvent{Common: Common{Type: "smbg", DeviceId: device, Source: "dinojr", Time: date}, Value: bgVal}
}

func makeNote(noteString, date, device string) *Common {
	return nil
	//return &NoteEvent{Common: Common{Type: "note", Source: "dinojr", DeviceId: device, Time: date}, Text: noteString, CreatorId, ""}
}

func makeCarb(carbString, date, device string) *FoodEvent {
	carb := strings.Split(carbString, CARB)
	carbVal, _ := strconv.ParseFloat(carb[1], 64)
	return &FoodEvent{Common: Common{Type: "food", DeviceId: device, Source: "dinojr", Time: date}, Carbs: carbVal}
}

func makeBolus(bolusString, date, device string) *BolusEvent {
	bolus := strings.Split(bolusString, BOLUS)
	bolusVal, _ := strconv.ParseFloat(bolus[1], 64)
	return &BolusEvent{Common: Common{Type: "bolus", DeviceId: device, Source: "dinojr", Time: date}, SubType: "injected", Value: bolusVal, Insulin: "novolog"}
}

func makeBasal(basalString, date, device string) *BasalEvent {
	basal := strings.Split(basalString, BASAL)
	basalVal, _ := strconv.ParseFloat(basal[1], 64)
	return &BasalEvent{Common: Common{Type: "basal", DeviceId: device, Source: "dinojr", Time: date}, DeliveryType: "injected", Value: basalVal, Insulin: "lantus", Duration: 86400000}
}
