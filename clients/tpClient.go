package clients

import (
	"../models"
	"bytes"
	"encoding/json"
	"errors"
	twilio "github.com/carlosdp/twiliogo"
	"io/ioutil"
	"log"
	"net/http"
)

type (
	//generic interface
	PlatformClient interface {
		LoadInto(userid string, data *twilio.MessageList) error
		LoadFrom(userid string) (interface{}, error)
	}
	TidepoolClient struct {
		config *TidepoolClientConfig
	}
	//
	TidepoolClientConfig struct {
		jellyFishUrl     string `json:"jellyFishUrl"`
		tideWhispererUrl string `json:"tideWhispererUrl"`
	}
)

func NewPlatformClient() *TidepoolClient {
	//hardcode for now
	return &TidepoolClient{config: &TidepoolClientConfig{jellyFishUrl: "http://localhost:9122/data", tideWhispererUrl: "http://localhost:9127/"}}
}

func (tc *TidepoolClient) loadBlock(block []byte) (int, error) {

	if res, err := http.Post(tc.config.jellyFishUrl, "application/json", bytes.NewBufferString(string(block))); err != nil {
		return 0, err
	} else {
		return res.StatusCode, nil
	}

}

func (tc *TidepoolClient) LoadInto(userid string, data *twilio.MessageList) error {

	for i := range data.Messages {
		message := data.Messages[i]

		block := models.Translate(message.Body, message.DateSent)

		jsonBlock, _ := json.Marshal(block)

		if status, err := tc.loadBlock(jsonBlock); err != nil {
			log.Println("Error loading data ", err)
			return err
		} else {
			if status != http.StatusCreated {
				return errors.New("An issue, we got " + string(status))
			}
		}
	}
	return nil
}

func (tc *TidepoolClient) LoadFrom(userid string) (data interface{}, err error) {

	res, err := http.Get(tc.config.tideWhispererUrl + userid)
	if err != nil {
		return data, err
	}
	data, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return data, err
	}
	return data, nil
}
