package parse

import (
	"encoding/json"
	"log"
	"testing"
	"time"
)

const SRC_DEVICE = "+666517771"

func TestTranslate_BG(t *testing.T) {

	txt := NewHealthDataText("G=8.9", time.Now().String(), SRC_DEVICE)
	msgEvents := txt.ParseHealthData()

	if len(msgEvents) != 1 {
		t.Fatalf("We expect one event %v", msgEvents)
	}

	jsonEvents, _ := json.Marshal(msgEvents)
	s := string(jsonEvents[:])

	log.Println(s)
}

func TestTranslate_BG_CARB(t *testing.T) {

	txt := NewHealthDataText("G=8.9 C=90", time.Now().String(), SRC_DEVICE)
	msgEvents := txt.ParseHealthData()

	if len(msgEvents) != 2 {
		t.Fatalf("We expect two events %v", msgEvents)
	}

	jsonEvents, _ := json.Marshal(msgEvents)
	s := string(jsonEvents[:])

	log.Println(s)
}

func TestTranslate_Low(t *testing.T) {

	txt := NewHealthDataText("#LG", time.Now().String(), SRC_DEVICE)
	msgEvents := txt.ParseHealthData()

	if len(msgEvents) != 1 {
		t.Fatalf("We expect one event %v", msgEvents)
	}

	jsonEvents, _ := json.Marshal(msgEvents)
	s := string(jsonEvents[:])

	log.Println(s)
}

func TestTranslate_BG_CARB_BOLUS(t *testing.T) {

	txt := NewHealthDataText("G=8.9 C=90 S=10", time.Now().String(), SRC_DEVICE)
	msgEvents := txt.ParseHealthData()

	if len(msgEvents) != 3 {
		t.Fatalf("We expect three events %v", msgEvents)
	}

	jsonEvents, _ := json.Marshal(msgEvents)
	s := string(jsonEvents[:])

	log.Println(s)
}

func TestTranslate_BG_CARB_BOLUS_BASAL(t *testing.T) {

	txt := NewHealthDataText("G=8.9 C=90 S=10 L=20", time.Now().String(), SRC_DEVICE)
	msgEvents := txt.ParseHealthData()

	if len(msgEvents) != 4 {
		t.Fatalf("We expect four events %v", msgEvents)
	}

	jsonEvents, _ := json.Marshal(msgEvents)
	s := string(jsonEvents[:])

	log.Println(s)
}

func TestTranslate_Bollocks(t *testing.T) {

	txt := NewHealthDataText("blah blah", time.Now().String(), SRC_DEVICE)
	msgEvents := txt.ParseHealthData()

	if len(msgEvents) != 1 {
		t.Fatalf("We expect one event %v", msgEvents)
	}

	jsonEvents, _ := json.Marshal(msgEvents)
	s := string(jsonEvents[:])

	log.Println(s)
}
