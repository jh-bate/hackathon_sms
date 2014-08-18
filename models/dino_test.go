package models

import (
	"encoding/json"
	"log"
	"testing"
	"time"
)

const SRC_DEVICE = "+666517771"

func TestTranslate_BG(t *testing.T) {

	msgEvents := Translate("BG=8.9", time.Now().String(), SRC_DEVICE)

	if len(msgEvents) != 1 {
		t.Fatalf("We expect one event %v", msgEvents)
	}

	jsonEvents, _ := json.Marshal(msgEvents)
	s := string(jsonEvents[:])

	log.Println(s)
}

func TestTranslate_BG_CARB(t *testing.T) {

	msgEvents := Translate("BG=8.9 CHO=90", time.Now().String(), SRC_DEVICE)

	if len(msgEvents) != 2 {
		t.Fatalf("We expect two events %v", msgEvents)
	}

	jsonEvents, _ := json.Marshal(msgEvents)
	s := string(jsonEvents[:])

	log.Println(s)
}

func TestTranslate_BG_CARB_BOLUS(t *testing.T) {

	msgEvents := Translate("BG=8.9 CHO=90 SA=10", time.Now().String(), SRC_DEVICE)

	if len(msgEvents) != 3 {
		t.Fatalf("We expect three events %v", msgEvents)
	}

	jsonEvents, _ := json.Marshal(msgEvents)
	s := string(jsonEvents[:])

	log.Println(s)
}

func TestTranslate_BG_CARB_BOLUS_BASAL(t *testing.T) {

	msgEvents := Translate("BG=8.9 CHO=90 SA=10 LA=20", time.Now().String(), SRC_DEVICE)

	if len(msgEvents) != 4 {
		t.Fatalf("We expect four events %v", msgEvents)
	}

	jsonEvents, _ := json.Marshal(msgEvents)
	s := string(jsonEvents[:])

	log.Println(s)
}

func TestTranslate_Bollocks(t *testing.T) {

	msgEvents := Translate("blah blah", time.Now().String(), SRC_DEVICE)

	if len(msgEvents) != 1 {
		t.Fatalf("We expect one event %v", msgEvents)
	}

	jsonEvents, _ := json.Marshal(msgEvents)
	s := string(jsonEvents[:])

	log.Println(s)
}
