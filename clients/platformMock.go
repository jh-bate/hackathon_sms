package clients

import (
	twilio "github.com/carlosdp/twiliogo"
	"log"
)

type (
	MockClient struct {
		token string
	}
)

func NewMockClient(usr, pw string) (*MockClient, error) {

	client := &MockClient{}

	if tkn, err := client.login(usr, pw); err != nil {
		log.Println("Error init client: ", err)
		return nil, err
	} else {
		client.token = tkn
		return client, nil
	}
}

func (tc *MockClient) login(usr, pw string) (token string, err error) {
	return "fairy.dust.as.a.token", nil
}

func (tc *MockClient) LoadSmsMessages(data *twilio.MessageList) error {
	log.Println("loding sms messages into the platform")
	return nil
}

func (tc *MockClient) SendSmsMessage(smsBody string) error {
	log.Println("loding sms messages into the platform")
	return nil
}
