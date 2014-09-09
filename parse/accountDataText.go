package parse

import (
	"log"
	"strings"
)

const (
	//Account
	ACCT_ADD_VEIWER   = "ACC_V="
	ACCT_ADD_UPLOADER = "ACC_U="
	ACCT_NAME         = "ACC_N="
)

/*
 * Text based message data the contains values we can parse
 */
type AccountData struct {
	text, date string
}

func NewAccountDataFromText(text, date string) *AccountData {
	return &AccountData{
		text: text,
		date: date,
	}
}

func (self *AccountData) ParseAccountData() []interface{} {

	var account []interface{}

	raw := strings.Split(self.text, " ")

outer:
	for i := range raw {
		switch {
		case strings.Index(strings.ToUpper(raw[i]), ACCT_NAME) != -1:
			log.Println("update account name ", self.text)
			//account = append(account, makeBg(raw[i], self.date))
			break
		case strings.Index(strings.ToUpper(raw[i]), ACCT_ADD_VEIWER) != -1:
			log.Println("add acct viewer ", self.text)
			//account = append(account, makeCarb(raw[i], self.date))
			break
		case strings.Index(strings.ToUpper(raw[i]), ACCT_ADD_UPLOADER) != -1:
			log.Println("add account uploader ", self.text)
			//account = append(account, makeBasal(raw[i], self.date))
			break
		default:
			log.Println("Not valid ", self.text)
			break outer
		}
	}
	return account
}
