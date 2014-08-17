package clients

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type (
	//generic interface
	PlatformClient interface {
		LoadInto(userid string, data interface{}) error
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

func (tc *TidepoolClient) LoadInto(userid string, data interface{}) error {

	jsonData, _ := json.Marshal(data)

	if res, err := http.Post(tc.config.jellyFishUrl, "application/json", bytes.NewBufferString(string(jsonData))); err != nil {
		log.Println("Error loading data into platform ", err)
		return err
	} else {
		if res.StatusCode != http.StatusCreated {
			return errors.New("The loading failed and returned status :" + string(res.StatusCode))
		}
		return nil
	}
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
