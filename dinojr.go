package main

import (
	"./clients"
	"encoding/json"
	twilio "github.com/carlosdp/twiliogo"
	"io/ioutil"
	"log"
	"net/url"
	"time"
)

type (
	UserConfig struct {
		AccountSid string `json:"accountSid"`
		AuthToken  string `json:"authToken"`
		User       string `json:"user"`
		Pw         string `json:"pw"`
	}

	Config struct {
		Platform *clients.Config `json:"platform"`
		User     UserConfig      `json:"dinojr"`
	}
)

var (
	testMessage = twilio.Message{
		Sid:         "testsid",
		DateCreated: time.Now().Format(time.RFC3339Nano),
		DateUpdated: time.Now().Format(time.RFC3339Nano),
		DateSent:    time.Now().Format(time.RFC3339Nano),
		AccountSid:  "AC3TestAccount",
		From:        "+15555555555",
		To:          "+16666666666",
		Body:        "B=6.7 C=90 S=10 L=20 #l",
		NumSegments: "1",
		Status:      "queued",
		Direction:   "outbound-api",
		Price:       "4",
		PriceUnit:   "dollars",
		ApiVersion:  "2008-04-01",
		Uri:         "/2010-04-01/Accounts/AC3TestAccount/Messages/testsid.json",
	}
	testMessages = twilio.MessageList{
		Messages: []twilio.Message{testMessage},
	}
)

func loadSmsMessages(smsClient twilio.Client) *twilio.MessageList {

	if messages, err := twilio.GetMessageList(smsClient); err != nil {
		log.Panic(err)
		return nil
	} else {
		return messages
	}

}

func sendSmsMessage(smsClient twilio.Client, frm, to, msg string) (*twilio.Message, error) {

	if message, err := twilio.NewMessage(smsClient, frm, to, twilio.Body(msg)); err != nil {
		log.Println(err)
		return message, nil
	} else {
		return message, nil
	}

}

func main() {

	var config Config

	jsonConfig, _ := ioutil.ReadFile("./config/script.json")
	_ = json.Unmarshal(jsonConfig, &config)

	//smsClient := twilio.NewClient(config.User.AccountSid, config.User.AuthToken)
	smsClient := new(twilio.MockClient)

	//mocking it!!!
	messagesJson, _ := json.Marshal(testMessages)
	smsClient.On("get", url.Values{}, smsClient.RootUrl()+"/SMS/Messages.json").Return(messagesJson, nil)

	twilioMsgs := loadSmsMessages(smsClient)

	if platform, err := clients.NewClient(
		config.Platform,
		config.User.User,
		config.User.Pw,
	); err != nil {
		//its all over!!
		log.Panic(err)
	} else {
		log.Println("yay logged in as ", config.User.User)
		//load the data
		if twilioMsgs != nil {
			//log.Println("loading ... ", twilioMsgs)
			if err := platform.LoadSmsMessages(twilioMsgs); err != nil {
				log.Println("Error pushing data to platform ", err)
			}
		}
	}

}
