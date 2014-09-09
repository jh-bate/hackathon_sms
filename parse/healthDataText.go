package parse

import (
	"../models"
	"log"
	"strconv"
	"strings"
)

const (
	//Health
	TIME     = "T="
	ACTIVITY = "A="
	BG       = "G="
	CARB     = "C="
	BASAL    = "L="
	BOLUS    = "S="
	NOTE     = "N="

	//Calcs
	LOG_LOW    = "#LG"
	CALC_BOLUS = "#BOL"
	CALC_IOB   = "#IOB"
	CALC_ISB   = "#ISF"

	MMOLL = "mmol/L"
)

/*
 * Text based message data the contains values we can parse
 */
type HealthData struct {
	text, date, device string
}

func NewHealthDataText(text, date, device string) *HealthData {
	return &HealthData{
		text:   text,
		date:   date,
		device: device,
	}
}

func (self *HealthData) ParseHealthData() []interface{} {

	var events []interface{}

	raw := strings.Split(self.text, " ")

outer:
	for i := range raw {
		switch {
		case strings.Index(strings.ToUpper(raw[i]), BG) != -1:
			events = append(events, makeBg(raw[i], self.date, self.device))
			break
		case strings.Index(strings.ToUpper(raw[i]), CARB) != -1:
			events = append(events, makeCarb(raw[i], self.date, self.device))
			break
		case strings.Index(strings.ToUpper(raw[i]), BASAL) != -1:
			events = append(events, makeBasal(raw[i], self.date, self.device))
			break
		case strings.Index(strings.ToUpper(raw[i]), BOLUS) != -1:
			events = append(events, makeBolus(raw[i], self.date, self.device))
			break
		case strings.Index(strings.ToUpper(raw[i]), LOG_LOW) != -1:
			//hard code 'LOW'
			events = append(events, makeBg("G=3.9", self.date, self.device))
			break
		case strings.Index(strings.ToUpper(raw[i]), ACTIVITY) != -1:
			log.Println("Will be an activity ", raw[i])
			break
		case strings.Index(strings.ToUpper(raw[i]), NOTE) != -1:
			events = append(events, makeNote(raw[i], self.date, self.device))
			break
		default:
			events = append(events, makeNote(self.text, self.date, self.device))
			break outer
		}
	}
	return events
}

func makeBg(bgString, date, device string) *models.BgEvent {
	bg := strings.Split(bgString, BG)
	bgVal, _ := strconv.ParseFloat(bg[1], 64)

	return &models.BgEvent{Common: models.Common{Type: "smbg", DeviceId: device, Source: "dinojr", Time: date}, Value: bgVal}
}

func makeNote(noteString, date, device string) *models.NoteEvent {
	return &models.NoteEvent{Common: models.Common{Type: "note", Source: "dinojr", DeviceId: device, Time: date}, Text: noteString, CreatorId: device}
}

func makeCarb(carbString, date, device string) *models.FoodEvent {
	carb := strings.Split(carbString, CARB)
	carbVal, _ := strconv.ParseFloat(carb[1], 64)
	return &models.FoodEvent{Common: models.Common{Type: "food", DeviceId: device, Source: "dinojr", Time: date}, Carbs: carbVal}
}

func makeBolus(bolusString, date, device string) *models.BolusEvent {
	bolus := strings.Split(bolusString, BOLUS)
	bolusVal, _ := strconv.ParseFloat(bolus[1], 64)
	return &models.BolusEvent{Common: models.Common{Type: "bolus", DeviceId: device, Source: "dinojr", Time: date}, SubType: "injected", Value: bolusVal, Insulin: "novolog"}
}

func makeBasal(basalString, date, device string) *models.BasalEvent {
	basal := strings.Split(basalString, BASAL)
	basalVal, _ := strconv.ParseFloat(basal[1], 64)
	return &models.BasalEvent{Common: models.Common{Type: "basal", DeviceId: device, Source: "dinojr", Time: date}, DeliveryType: "injected", Value: basalVal, Insulin: "lantus", Duration: 86400000}
}
