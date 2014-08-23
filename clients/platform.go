package clients

import (
	"../models"
	"bytes"
	"encoding/json"
	"errors"
	twilio "github.com/carlosdp/twiliogo"
	"log"
	"net/http"
)

const (
	TP_SESSION_TOKEN = "x-tidepool-session-token"
)

type (
	Platform interface {
		login(usr, pw string) (string, error)
		LoadSmsMessages(data *twilio.MessageList) error
		SendSmsMessage(smsBody string) error
	}
	Client struct {
		config     *Config
		token      string
		httpClient *http.Client
		user       []map[string]string
	}
	Config struct {
		Auth   string `json:"auth"`
		Upload string `json:"upload"`
	}
)

func NewClient(cfg *Config, usr, pw string) (*Client, error) {

	client := &Client{config: cfg, httpClient: &http.Client{}}

	if tkn, err := client.login(usr, pw); err != nil {
		log.Println("Error init client: ", err)
		return nil, err
	} else {
		//log.Println("here it is ", tkn)
		client.token = tkn
		return client, nil
	}
}

func (tc *Client) login(usr, pw string) (token string, err error) {

	req, err := http.NewRequest("POST", tc.config.Auth+"/login", nil)
	req.SetBasicAuth(usr, pw)
	if resp, err := tc.httpClient.Do(req); err != nil {
		return "", err
	} else {
		if resp.StatusCode == http.StatusOK {
			return resp.Header.Get(TP_SESSION_TOKEN), nil
		}
		return "", errors.New("Issue logging in: " + string(resp.StatusCode))
	}
}

func (tc *Client) LoadSmsMessages(data *twilio.MessageList) error {

	for i := range data.Messages {
		message := data.Messages[i]

		block := models.Translate(message.Body, message.DateSent, message.From)

		jsonBlock, _ := json.Marshal(block)

		log.Println(" block to load ", bytes.NewBufferString(string(jsonBlock)))
		//log.Println(" token ", tc.token)

		req, _ := http.NewRequest("POST", tc.config.Upload, bytes.NewBufferString(string(jsonBlock)))
		req.Header.Add(TP_SESSION_TOKEN, tc.token)
		req.Header.Set("content-type", "application/json")

		if resp, err := tc.httpClient.Do(req); err != nil {
			log.Println("Error loading messages: ", err)
			return err
		} else {
			log.Printf("all good? [%d] [%s] ", resp.StatusCode, resp.Status)
			updatedToken := resp.Header.Get(TP_SESSION_TOKEN)
			if updatedToken != "" && tc.token != updatedToken {
				tc.token = updatedToken
				log.Println("updated the token")
			}
		}
	}

	return nil
}

func (tc *Client) SendSmsMessage(smsBody string) error {
	return nil
}
