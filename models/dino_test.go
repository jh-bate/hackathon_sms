package models

import (
	"testing"
)

func TestTranslate_BG(t *testing.T) {

	msgEvents := Translate("BG=8.9")

	if len(msgEvents) != 1 {
		t.Fatalf("We expect one event %v", msgEvents)
	}
}

func TestTranslate_BG_CARB(t *testing.T) {

	msgEvents := Translate("BG=8.9 CHO=90")

	if len(msgEvents) != 2 {
		t.Fatalf("We expect two events %v", msgEvents)
	}
}

func TestTranslate_BG_CARB_BOLUS(t *testing.T) {

	msgEvents := Translate("BG=8.9 CHO=90 SA=10")

	if len(msgEvents) != 3 {
		t.Fatalf("We expect three events %v", msgEvents)
	}
}

func TestTranslate_BG_CARB_BOLUS_BASAL(t *testing.T) {

	msgEvents := Translate("BG=8.9 CHO=90 SA=10 LA=20")

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
