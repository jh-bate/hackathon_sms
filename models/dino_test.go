package models

import (
	"log"
	"testing"
)

func TestTranslate_BG(t *testing.T) {

	msgEvents := Translate("BG=8.9")

	log.Println(msgEvents[0])

	if len(msgEvents) != 1 {
		t.Fatalf("We expect one event %v", msgEvents)
	}
}

func TestTranslate_BG_CARB(t *testing.T) {

	msgEvents := Translate("BG=8.9 CHO=90")

	log.Println(msgEvents[0])
	log.Println(msgEvents[1])

	if len(msgEvents) != 2 {
		t.Fatalf("We expect two events %v", msgEvents)
	}
}

func TestTranslate_BG_CARB_BOLUS(t *testing.T) {

	msgEvents := Translate("BG=8.9 CHO=90 SA=10")

	log.Println(msgEvents[0])
	log.Println(msgEvents[1])
	log.Println(msgEvents[2])

	if len(msgEvents) != 3 {
		t.Fatalf("We expect three events %v", msgEvents)
	}
}

func TestTranslate_BG_CARB_BOLUS_BASAL(t *testing.T) {

	msgEvents := Translate("BG=8.9 CHO=90 SA=10 LA=20")

	log.Println(msgEvents[0])
	log.Println(msgEvents[1])
	log.Println(msgEvents[2])
	log.Println(msgEvents[3])

	if len(msgEvents) != 4 {
		t.Fatalf("We expect four events %v", msgEvents)
	}
}

func TestTranslate_Bollocks(t *testing.T) {

	msgEvents := Translate("blah blah")

	if len(msgEvents) != 1 {
		t.Fatalf("We expect one event %v", msgEvents)
	}
}
