package parse

import (
	"log"
	"testing"
	"time"
)

func TestTranslate_AccountName(t *testing.T) {

	txt := NewAccountDataFromText("ACC_N=Jamie", time.Now().String())
	accEvents := txt.ParseAccountData()
	log.Println("account name ", accEvents)

}

func TestTranslate_AddUploader(t *testing.T) {

	txt := NewAccountDataFromText("ACC_U=+64123454", time.Now().String())
	accEvents := txt.ParseAccountData()
	log.Println("account uploader ", accEvents)

}

func TestTranslate_AddViewer(t *testing.T) {

	txt := NewAccountDataFromText("ACC_V=+64123556", time.Now().String())
	accEvents := txt.ParseAccountData()
	log.Println("account viewer ", accEvents)

}
