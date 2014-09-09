package clients

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
	Sms interface {
		NewSmsClient(cfg SmsConfig) (*SmsClient, error)
		GetSms(providerClient interface{}) (*[]interface{}, error)
		SendSms(providerClient interface{}, from, to, msg string) (*interface{}, error)
	}
	SmsClient struct {
		config *SmsConfig
	}
	SmsConfig struct {
		AccountSid string `json:"accountSid"`
		AuthToken  string `json:"authToken"`
	}
)

func NewSmsClient(cfg *Config) (*SmsClient, error) {
	return &SmsClient{config: cfg}
}

func (c *SmsClient) GetSms(t twilio.Client) *twilio.MessageList {

	if messages, err := twilio.GetMessageList(smsClient); err != nil {
		log.Panic(err)
		return nil
	} else {
		return messages
	}
}

func (c *SmsClient) SendSms(smsClient twilio.Client, frm, to, msg string) (*twilio.Message, error) {

	if message, err := twilio.NewMessage(smsClient, frm, to, twilio.Body(msg)); err != nil {
		log.Println(err)
		return message, nil
	} else {
		return message, nil
	}
}
